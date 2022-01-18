package cmd

import (
	"github.com/denizyoldas/spoty/services"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(devicesCmd)
}

var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(services.GetAvailableDevices())
	},
}
