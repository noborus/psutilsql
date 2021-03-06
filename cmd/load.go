package cmd

import (
	"github.com/noborus/psutilsql"
	"github.com/spf13/cobra"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "load information",
	Long: `load information.

CPU load average
+-------+-------+--------+
| Load1 | Load5 | Load15 |
+-------+-------+--------+

Option misc returnes miscellaneous host-wide statistics
+------------+--------------+--------------+------+
| ProcsTotal | ProcsRunning | ProcsBlocked | Ctxt |
+------------+--------------+--------------+------+
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var misc bool
		if misc, err = cmd.PersistentFlags().GetBool("misc"); err != nil {
			return err
		}
		return psutilsql.LoadQuery(misc, Query, outFormat())
	},
}

func init() {
	loadCmd.PersistentFlags().BoolP("misc", "m", false, "miscellaneous host-wide statistics")
	rootCmd.AddCommand(loadCmd)
}
