package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of psutilsql",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("psutilsql version %s rev:%s\n", Version, Revision)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
