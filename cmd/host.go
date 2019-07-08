package cmd

import (
	"github.com/noborus/psutilsql"
	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "host information",
	Long: `host information.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var tempera, users bool
		if tempera, err = cmd.PersistentFlags().GetBool("temperatures"); err != nil {
			return err
		}
		if users, err = cmd.PersistentFlags().GetBool("users"); err != nil {
			return err
		}

		return psutilsql.HostQuery(tempera, users, Query, outFormat())
	},
}

func init() {
	hostCmd.PersistentFlags().BoolP("temperatures", "t", false, "SensorsTemperatures")
	hostCmd.PersistentFlags().BoolP("users", "u", false, "users")

	rootCmd.AddCommand(hostCmd)
}
