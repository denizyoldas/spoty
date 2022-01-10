package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(getClientCmd)
}

var getClientCmd = &cobra.Command{
	Use:   "get-client",
	Short: "set client id",
	Long:  `set client id`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Client ID: %v", viper.GetString("ACCESS_TOKEN"))
	},
}
