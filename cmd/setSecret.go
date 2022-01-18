package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(setSecretCmd)
}

var setSecretCmd = &cobra.Command{
	Use:   "set-secret",
	Short: "set secret key",
	Long:  `set secret key`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("Please enter a secret key")
			return
		}

		viper.Set("ACCESS_TOKEN", args[0])
		fmt.Printf("Secret: %v", viper.GetString("SECRET"))
		err := viper.WriteConfig()

		if err != nil {
			fmt.Printf("%v", err)
		}
	},
}
