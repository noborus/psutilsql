// +build windows

package cmd

import (
	"github.com/noborus/psutilsql"

	"github.com/spf13/cobra"
)

// winservicesCmd represents the winservices command
var winservicesCmd = &cobra.Command{
	Use:   "winservices",
	Short: "winservices information",
	Long: `winservices information.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return psutilsql.Winservices(Query, outFormat())
	},
}

func init() {
	rootCmd.AddCommand(winservicesCmd)
}
