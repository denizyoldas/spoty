package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const repo = "denizyoldas/spoty"

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update CLI to the latest version.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Println("🎁 update")
		//isUpdated, err := IsUpdated(cmd)
		//if err != nil {
		//	return err
		//}
		//if isUpdated {
		//	return errors.New(internal.ErrAlreadyUpToDate)
		//}
		//
		//current, err := parseVersion(cmd)
		//
		//latest, err := selfupdate.UpdateSelf(current, repo)
		//
		//cmd.Printf("🎁 v%s\n", latest.Version.String())
	},
}

//func IsUpdated(cmd *cobra.Command) (bool, error) {
//	current, err := parseVersion(cmd)
//	if err != nil {
//		return true, err
//	}
//
//	//latest, found, err := selfupdate.DetectLatest(repo)
//	if err != nil {
//		return true, err
//	}
//
//	//isUpdated := !found || current.Equals(latest.Version)
//	//return isUpdated, nil
//	return false, nil
//}
//
//func parseVersion(cmd *cobra.Command) (semver.Version, error) {
//	version := strings.TrimPrefix(cmd.Root().Version, "v")
//	return semver.Parse(version)
//}
