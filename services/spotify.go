package services

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var c = &http.Client{}

type Payload map[string]string

const API = "https://api.spotify.com/v1/"

func Login() error {
	//var body = make(map[string]string)
	//body["grant_type"] = "client_credentials"
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	headerVal := b64.StdEncoding.EncodeToString([]byte(viper.GetString("CLIENT_ID") + ":" + viper.GetString("SECRET")))
	req.Header.Set("Authorization", "Basic "+headerVal)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := c.Do(req)
	if err != nil {
		return err
	}

	//defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Printf("%v", bodyString)

	// res, err := http.Post("https://accounts.spotify.com/api/token", "application/json", nil)
	// if err != nil {
	//	return err
	// }
	return nil
}

func NextSong() error {
	// postBody, _ := json.Marshal(data)
	// responseBody := bytes.NewBuffer(postBody)
	res, err := http.Post(API+"me/player/next", "application/json", nil)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	sb := string(body)
	log.Printf(sb)
	return nil
}
