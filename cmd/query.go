package cmd

import (
	"github.com/noborus/psutilsql"
	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "SQL query command",
	RunE: func(cmd *cobra.Command, args []string) error {
		return psutilsql.QueryExec(args[0], outFormat())
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
