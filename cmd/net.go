package cmd

import (
	"github.com/shirou/gopsutil/net"

	"github.com/spf13/cobra"
)

// netCmd represents the net command
var netCmd = &cobra.Command{
	Use:   "net",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		defaultQuery := "SELECT * FROM net"
		var err error
		var v interface{}

		v, err = net.Connections("all")
		if err != nil {
			return err
		}
		query := Query
		if query == "" {
			query = defaultQuery
		}
		return sliceQuery(v, "net", query)
	},
}

func init() {
	netCmd.PersistentFlags().BoolP("conntrack", "t", false, "connection tracking")
	rootCmd.AddCommand(netCmd)
}
