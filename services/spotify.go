package services

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var c = &http.Client{}

type Payload map[string]interface{}

const API = "https://api.spotify.com/v1/"

type SpotifyClient struct {
	http    *http.Client
	baseURL string

	autoRetry      bool
	acceptLanguage string
}

func (c *SpotifyClient) Check(ctx context.Context) {
}

var ch = make(chan *SpotifyClient)

func Token(code, state string) error {
	// fmt.Printf("code : %v", code)
	data := url.Values{}
	data.Set("code", code)
	data.Set("redirect_uri", "http://localhost:3000/callback")
	data.Set("grant_type", "authorization_code")
	req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	headerVal := b64.StdEncoding.EncodeToString([]byte(viper.GetString("CLIENT_ID") + ":" + viper.GetString("SECRET")))
	req.Header.Set("Authorization", "Basic "+headerVal)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := c.Do(req)
	if err != nil {
		return err
	}

	//defer res.Body.Close()

	bodyBytes, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var jsonRes Payload
	cobra.CheckErr(json.Unmarshal(bodyBytes, &jsonRes))

	fmt.Printf("%v", jsonRes)

	viper.Set("access_token", jsonRes["access_token"])
	viper.Set("expires_in", jsonRes["expires_in"])
	viper.Set("refresh_token", jsonRes["refresh_token"])
	viper.Set("expires_date", time.Now().Local().
		Add(time.Hour*time.Duration(viper.GetInt("expires_in")/60/60)))

	cobra.CheckErr(viper.WriteConfig())
	return nil
}

func RefreshToken() {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", viper.GetString("refresh_token"))
	req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	headerVal := b64.StdEncoding.EncodeToString([]byte(viper.GetString("CLIENT_ID") + ":" + viper.GetString("SECRET")))
	req.Header.Set("Authorization", "Basic "+headerVal)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := c.Do(req)
	cobra.CheckErr(err)

	bodyBytes, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var jsonRes Payload
	cobra.CheckErr(json.Unmarshal(bodyBytes, &jsonRes))

	viper.Set("access_token", jsonRes["access_token"])
	viper.Set("expires_in", jsonRes["expires_in"])
	viper.Set("expires_date", time.Now().Local().
		Add(time.Hour*time.Duration(viper.GetInt("expires_in")/60/60)))

	cobra.CheckErr(viper.WriteConfig())
}

func Login() error {
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go func() {
		err := http.ListenAndServe(":3000", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	shortUrl := "https://accounts.spotify.com/authorize?" + "response_type=code&client_id=" +
		viper.GetString("client_id") +
		"&redirect_uri=http://localhost:3000/callback&scope=user-read-private%20user-modify-playback-state%20user-read-playback-state&state=" +
		StringCode(16)

	fmt.Println("Please log in to Spotify by visiting the following page in your browser: \n", shortUrl)

	client := <-ch

	client.Check(context.Background())
	return nil
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if query.Get("error") != "" {
		log.Fatal("error: ", query.Get("error"))
	}

	if query.Get("code") != "" && query.Get("state") != "" {
		err := Token(query.Get("code"), query.Get("state"))
		if err != nil {
			log.Fatal(err)
		}
	}

	ch <- nil
}

func NextSong() error {
	CheckExpDate()
	req, _ := http.NewRequest("POST", "https://api.spotify.com/v1/me/player/next", nil)
	req.Header.Set("Authorization", "Bearer "+viper.GetString("access_token"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	_, err := c.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func PrevSong() error {
	CheckExpDate()
	req, _ := http.NewRequest("POST", "https://api.spotify.com/v1/me/player/previous", nil)
	req.Header.Set("Authorization", "Bearer "+viper.GetString("access_token"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	_, err := c.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func Pause() error {
	CheckExpDate()
	req, _ := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/pause", nil)
	req.Header.Set("Authorization", "Bearer "+viper.GetString("access_token"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	_, err := c.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func Play() error {
	CheckExpDate()
	req, _ := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/play", nil)
	req.Header.Set("Authorization", "Bearer "+viper.GetString("access_token"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	_, err := c.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func Volume(val string) error {
	CheckExpDate()
	req, _ := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/volume", nil)
	query := req.URL.Query()
	query.Set("volume_percent", val)
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Authorization", "Bearer "+viper.GetString("access_token"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	_, err := c.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func GetAvailableDevices() error {
	CheckExpDate()
	req, _ := http.NewRequest("GET", API+"me/player/devices", nil)
	query := req.URL.Query()
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Authorization", "Bearer "+viper.GetString("access_token"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	_, err := c.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func CheckExpDate() {
	now := time.Now()
	date := viper.GetTime("expires_date")
	if date.Sub(now) <= 0 {
		RefreshToken()
	}
}
