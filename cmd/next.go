package cmd

import (
	"fmt"
	"github.com/denizyoldas/spoty/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(nextCmd)
}

var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Next song",
	Long:  `This command will be executed next to each song`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("client %v\ntoken %v", viper.GetString("CLIENT_ID"), viper.GetString("SECRET"))
		cobra.CheckErr(services.NextSong())
	},
}
