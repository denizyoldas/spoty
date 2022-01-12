package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var spoty = []byte(`
access_token: tokenexpmle
client_id: gettheclientid
device_id: nonerequired
expires_date: tokenexpDate
expires_in: 3600
refresh_token: refreshtoken
secret: scrtkey
volume: 100
`)
var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config-path", "", "config file (default is $HOME/.spoty.yml)")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	rootCmd.PersistentFlags().IntP("volume", "v", 0, "set volume")
	err := viper.BindPFlag("volume", rootCmd.PersistentFlags().Lookup("volume"))
	if err != nil {
		return
	}
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".spoty")
		viper.SetConfigType("yaml")
		viper.SetConfigType("yml")
	}

	if err := viper.ReadInConfig(); err != nil {
		cobra.CheckErr(createFile())
	}
}

var rootCmd = &cobra.Command{
	Use:   "spoty",
	Short: "Spoty is a spotify fast cli",
	Long:  `Spoty best spotify cli tools`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createFile() error {
	home, err := homedir.Dir()
	cobra.CheckErr(err)
	file, err := os.Create(home + "\\.spoty.yml")
	if err != nil {
		return err
	}

	file.Write(spoty)
	file.Close()

	return nil
}
