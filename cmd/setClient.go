package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(stClientCmd)
}

var stClientCmd = &cobra.Command{
	Use:   "set-client",
	Short: "set client id",
	Long:  `set client id`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("Please enter a client id")
			return
		}

		viper.Set("ACCESS_TOKEN", args[0])
		viper.Set("verbose", true)
		fmt.Printf("Client ID: %v", viper.GetString("ACCESS_TOKEN"))
	},
}
