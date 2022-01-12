package cmd

import (
	"fmt"
	"github.com/denizyoldas/spoty/services"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(volumeCmd)
}

var volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "set volume",
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("volume: %v", args[0])
		cobra.CheckErr(services.Volume(args[0]))
	},
}
