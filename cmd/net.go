package cmd

import (
	"github.com/noborus/psutilsql"

	"github.com/spf13/cobra"
)

// netCmd represents the net command
var netCmd = &cobra.Command{
	Use:   "net",
	Short: "net information",
	Long: `net information.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return psutilsql.NetQuery(Query, outFormat())
	},
}

func init() {
	//	netCmd.PersistentFlags().BoolP("conntrack", "t", false, "connection tracking")
	rootCmd.AddCommand(netCmd)
}
