package cmd

import (
	"fmt"

	"github.com/ktr0731/go-semver"
	"github.com/spf13/cobra"
)

var currentVersion = semver.MustParse("0.0.1")
// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long: "Show nandokuka's version and exit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("nandokuka v%s\n",currentVersion)
	},
}


func init() {
	rootCmd.AddCommand(versionCmd)
}
