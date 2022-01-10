package cmd

import (
	"github.com/denizyoldas/spoty/services"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(refreshCmd)
}

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "get new token",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		services.RefreshToken()
	},
}
