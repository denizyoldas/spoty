package cmd

import (
	"fmt"
	"github.com/denizyoldas/spoty/services"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "you are version information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(services.VERSION)
	},
}
