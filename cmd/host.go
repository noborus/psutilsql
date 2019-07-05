package cmd

import (
	"github.com/shirou/gopsutil/host"
	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		defaultQuery := "SELECT * FROM host"

		var err error
		var v interface{}
		var tempera,users bool
		if tempera, err = cmd.PersistentFlags().GetBool("temperatures"); err != nil {
			return err
		}
		if users, err = cmd.PersistentFlags().GetBool("users"); err != nil {
			return err
		}
		
		if tempera {
			v, err = host.SensorsTemperatures()
		} else if users{
			v, err = host.Users()
		} else{
			v, err = host.Info()
		}
		if err != nil {
			return err
		}

		query := Query
		if query == "" {
			query = defaultQuery
		}
		return sliceQuery(v, "host", query)
	},
}

func init() {
	hostCmd.PersistentFlags().BoolP("temperatures", "t", false, "SensorsTemperatures")
	hostCmd.PersistentFlags().BoolP("users", "u", false, "users")

	rootCmd.AddCommand(hostCmd)
}
