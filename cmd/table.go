package cmd

import (
	"github.com/noborus/psutilsql"

	"github.com/spf13/cobra"
)

// tableCmd represents the table command
var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "table list",
	RunE: func(cmd *cobra.Command, args []string) error {
		var table string
		if len(args) > 0 {
			table = args[0]
		}
		return psutilsql.PSTableQuery(table, Query, outFormat())
	},
}

func init() {
	rootCmd.AddCommand(tableCmd)
}
