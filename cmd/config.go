package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "configuration command",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("client_id", viper.GetString("client-id"))
		viper.Set("secret", viper.GetString("secret-key"))
	},
}

func init() {
	configCmd.PersistentFlags().String("client-id", "", "set client-id")
	configCmd.PersistentFlags().String("secret-key", "", "set secret-key")
	cobra.CheckErr(viper.BindPFlag("client-id", configCmd.PersistentFlags().Lookup("client-id")))
	cobra.CheckErr(viper.BindPFlag("secret-key", configCmd.PersistentFlags().Lookup("secret-key")))
	rootCmd.AddCommand(configCmd)
}
