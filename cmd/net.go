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
		v, err := net.Pids()
		if err != nil {
			return err
		}
		query := "SELECT * FROM load"
		if Query != "" {
			query = Query
		}
		return sliceQuery(v, "load", query)
	},
}

func init() {
	rootCmd.AddCommand(netCmd)
}
