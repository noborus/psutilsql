package cmd

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/spf13/cobra"
)

// cpuCmd represents the cpu command
var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		defaultQuery := "SELECT * FROM cpu ORDER BY CPU"

		var err error
		var per, info, total bool
		if per, err = cmd.PersistentFlags().GetBool("percent"); err != nil {
			return err
		}
		if info, err = cmd.PersistentFlags().GetBool("info"); err != nil {
			return err
		}
		if total, err = cmd.PersistentFlags().GetBool("total"); err != nil {
			return err
		}

		var v interface{}
		if per {
			defaultQuery = "SELECT * FROM cpu"
			v, err = cpu.Percent(0, !total)
		} else if info {
			v, err = cpu.Info()
		} else {
			v, err = cpu.Times(!total)
		}
		if err != nil {
			return err
		}
		query := Query
		if query == "" {
			query = defaultQuery
		}
		return sliceQuery(v, "cpu", query)
	},
}

func init() {
	rootCmd.AddCommand(cpuCmd)
	cpuCmd.PersistentFlags().BoolP("info", "i", false, "CPU info")
	cpuCmd.PersistentFlags().BoolP("percent", "p", false, " per CPU")
	cpuCmd.PersistentFlags().BoolP("total", "t", false, "total CPU info")
}
