package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// You can manually update this, or use build flags to inject it
var appVersion = "v0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ask",
	Long:  `All software has versions. This is ask's.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ask CLI %s\n", appVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
