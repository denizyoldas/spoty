package cmd

import (
	"fmt"
	"github.com/denizyoldas/spoty/services"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pauseCmd)
}

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "song is stopped",
	Long:  `This command will be pause a song`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(services.Pause())
		fmt.Println("song paused!")
	},
}
