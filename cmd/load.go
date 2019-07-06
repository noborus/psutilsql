package cmd

import (
	"github.com/noborus/psutilsql"
	"github.com/shirou/gopsutil/load"
	"github.com/spf13/cobra"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "load information",
	Long: `load information.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		defaultQuery := "SELECT * FROM load"
		var err error
		var v interface{}
		var misc bool
		if misc, err = cmd.PersistentFlags().GetBool("misc"); err != nil {
			return err
		}
		if misc {
			v, err = load.Misc()
		} else {
			v, err = load.Avg()
		}
		if err != nil {
			return err
		}
		query := Query
		if query == "" {
			query = defaultQuery
		}
		return psutilsql.SliceQuery(v, "load", query, outFormat())
	},
}

func init() {
	loadCmd.PersistentFlags().BoolP("misc", "m", false, " miscellaneous host-wide statistics")
	rootCmd.AddCommand(loadCmd)
}
