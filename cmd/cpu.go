package cmd

import (
	"github.com/noborus/psutilsql"
	"github.com/spf13/cobra"
)

// cpuCmd represents the cpu command
var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "CPU information",
	Long: `CPU information

Option times (default) gets the result of cpu.Times().
The column names are CPU, User, System, Idle, Nice, Iowait, Irq, Softirq, Steal, Guest, GuestNice, Stolen.

Option info gets the result of cpu.Info().
The column names are CPU, VendorID, Family, Model, Stepping, PhysicalID, CoreID, Cores, ModelName, Mhz, CacheSize, Flags, Microcode.

Option percent gets the result of cpu.Percent().

Option total gets the result of the total on one row.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var per, info, total bool
		if total, err = cmd.PersistentFlags().GetBool("total"); err != nil {
			return err
		}
		if per, err = cmd.PersistentFlags().GetBool("percent"); err != nil {
			return err
		}
		if per {
			return psutilsql.CPUPercentQuery(total, Query, outFormat())
		}
		if info, err = cmd.PersistentFlags().GetBool("info"); err != nil {
			return err
		}
		if info {
			return psutilsql.CPUInfoQuery(Query, outFormat())
		}
		return psutilsql.CPUTimeQuery(total, Query, outFormat())
	},
}

func init() {
	rootCmd.AddCommand(cpuCmd)
	cpuCmd.PersistentFlags().BoolP("info", "i", false, "CPU info")
	cpuCmd.PersistentFlags().BoolP("percent", "p", false, " per CPU")
	cpuCmd.PersistentFlags().BoolP("total", "t", false, "total CPU info")
}
