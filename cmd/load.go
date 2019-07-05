package cmd

import (
	"github.com/shirou/gopsutil/load"
	"github.com/spf13/cobra"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
		return sliceQuery(v, "load", query)
	},
}

func init() {
	loadCmd.PersistentFlags().BoolP("misc", "m", false, " miscellaneous host-wide statistics")
	rootCmd.AddCommand(loadCmd)
}
