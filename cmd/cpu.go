package cmd

import (
	"github.com/noborus/psutilsql"

	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/cpu"
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
		var per, info, time, total bool

		if total, err = cmd.PersistentFlags().GetBool("total"); err != nil {
			return err
		}
		if per, err = cmd.PersistentFlags().GetBool("percent"); err != nil {
			return err
		}
		if info, err = cmd.PersistentFlags().GetBool("info"); err != nil {
			return err
		}
		if time, err = cmd.PersistentFlags().GetBool("time"); err != nil {
			return err
		}

		defaultQuery := "SELECT * FROM CPUTime ORDER BY cpu"
		var r *trdsql.SliceReader
		if per {
			defaultQuery = "SELECT * FROM CPUPercent"
			c, err := cpu.Percent(0, !total)
			if err != nil {
				return err
			}
			r = trdsql.NewSliceReader("CPUPercent", c)
		}
		if info {
			defaultQuery = "SELECT * FROM CPUInfo ORDER BY cpu"
			c, err := cpu.Info()
			if err != nil {
				return err
			}
			r = trdsql.NewSliceReader("CPUInfo", c)
		}
		if time {
			defaultQuery = "SELECT * FROM CPUTime ORDER BY cpu"
			c, err := cpu.Times(!total)
			if err != nil {
				return err
			}
			r = trdsql.NewSliceReader("CPUTime", c)
		}
		query := Query
		if query == "" {
			query = defaultQuery
		}
		importer, err := psutilsql.NewMultiImporter(r)
		if err != nil {
			return err
		}
		writer := trdsql.NewWriter(trdsql.OutFormat(outFormat()))
		trd := trdsql.NewTRDSQL(importer, trdsql.NewExporter(writer))
		err = trd.Exec(query)
		return err
	},
}

func init() {
	rootCmd.AddCommand(cpuCmd)
	cpuCmd.PersistentFlags().BoolP("info", "i", false, "CPU info")
	cpuCmd.PersistentFlags().BoolP("percent", "p", false, " per CPU")
	cpuCmd.PersistentFlags().BoolP("total", "t", false, "total CPU info")
	cpuCmd.PersistentFlags().BoolP("time", "", true, "CPU Time")
}
