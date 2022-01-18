package cmd

import (
	"fmt"
	"github.com/denizyoldas/spoty/services"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(prev)
}

var prev = &cobra.Command{
	Use:   "prev",
	Short: "Previous song",
	Long:  `This command will be executed previus to each song`,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	cobra.CheckErr(services.PrevSong())
	fmt.Println(services.NEXT_SONG)
}
