// +build windows

package cmd

import (
	"github.com/noborus/psutilsql"
	"github.com/shirou/gopsutil/winservices"

	"github.com/spf13/cobra"
)

// winservicesCmd represents the winservices command
var winservicesCmd = &cobra.Command{
	Use:   "winservices",
	Short: "winservices information",
	Long: `winservices information.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		defaultQuery := "SELECT * FROM winservices"

		v, err := winservices.Info()
		if err != nil {
			return err
		}
		query := Query
		if query == "" {
			query = defaultQuery
		}
		return psutilsql.SliceQuery(v, "winservices", query, outFormat())
	},
}

func init() {
	rootCmd.AddCommand(winservicesCmd)
}
