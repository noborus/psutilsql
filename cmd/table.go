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
		return psutilsql.PSTableQuery(Query, outFormat())
	},
}

func init() {
	rootCmd.AddCommand(tableCmd)
}
