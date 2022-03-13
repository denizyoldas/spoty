package cmd

import (
	"fmt"
	"github.com/denizyoldas/spoty/services"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(playCmd)
}

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "song is started",
	Long:  `This command will be start a song`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(services.Play())
		fmt.Println(services.START_SONG)
	},
}
