package cmd

import (
	"github.com/denizyoldas/spoty/services"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nextCmd)
}

var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Next song",
	Long:  `This command will be executed next to each song`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(services.NextSong())
	},
}
