package cmd

import (
	"github.com/noborus/psutilsql"
	"github.com/shirou/gopsutil/host"
	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "host information",
	Long: `host information.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		defaultQuery := "SELECT * FROM host"

		var err error
		var v interface{}
		var tempera, users bool
		if tempera, err = cmd.PersistentFlags().GetBool("temperatures"); err != nil {
			return err
		}
		if users, err = cmd.PersistentFlags().GetBool("users"); err != nil {
			return err
		}

		if tempera {
			v, err = host.SensorsTemperatures()
		} else if users {
			v, err = host.Users()
		} else {
			v, err = host.Info()
		}
		if err != nil {
			return err
		}

		query := Query
		if query == "" {
			query = defaultQuery
		}
		return psutilsql.SliceQuery(v, "host", query, outFormat())
	},
}

func init() {
	hostCmd.PersistentFlags().BoolP("temperatures", "t", false, "SensorsTemperatures")
	hostCmd.PersistentFlags().BoolP("users", "u", false, "users")

	rootCmd.AddCommand(hostCmd)
}
