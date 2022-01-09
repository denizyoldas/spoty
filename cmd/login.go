package cmd

import (
	"fmt"
	"github.com/denizyoldas/spoty/services"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login spotify",
	Long:  `This command will be login to spotify`,
	Run: func(cmd *cobra.Command, args []string) {
		err := services.Login()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
