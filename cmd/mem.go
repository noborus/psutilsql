package cmd

import (
	"github.com/noborus/psutilsql"
	"github.com/spf13/cobra"
)

// memCmd represents the mem command
var memCmd = &cobra.Command{
	Use:   "mem",
	Short: "memory information",
	Long: `memory information.

VirtualMemory information
+-------+-----------+------+-------------+------+--------+----------+-------+---------+---------+--------+-----------+-------+--------------+--------+------+--------------+------------+------------+------------+-------------+-------------+-----------+----------+----------+---------+-----------+----------+--------+--------------+-------------+--------------+----------------+---------------+--------------+
| Total | Available | Used | UsedPercent | Free | Active | Inactive | Wired | Laundry | Buffers | Cached | Writeback | Dirty | WritebackTmp | Shared | Slab | SReclaimable | SUnreclaim | PageTables | SwapCached | CommitLimit | CommittedAS | HighTotal | HighFree | LowTotal | LowFree | SwapTotal | SwapFree | Mapped | VMallocTotal | VMallocUsed | VMallocChunk | HugePagesTotal | HugePagesFree | HugePageSize |
+-------+-----------+------+-------------+------+--------+----------+-------+---------+---------+--------+-----------+-------+--------------+--------+------+--------------+------------+------------+------------+-------------+-------------+-----------+----------+----------+---------+-----------+----------+--------+--------------+-------------+--------------+----------------+---------------+--------------+

Option swap gets the result of swap information
+-------+------+------+-------------+-----+------+------+-------+---------+
| Total | Used | Free | UsedPercent | Sin | Sout | PgIn | PgOut | PgFault |
+-------+------+------+-------------+-----+------+------+-------+---------+
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var swap bool
		if swap, err = cmd.PersistentFlags().GetBool("swap"); err != nil {
			return err
		}
		return psutilsql.MEMQuery(!swap, Query, outFormat())
	},
}

func init() {
	memCmd.PersistentFlags().BoolP("swap", "s", false, "swap memory")
	rootCmd.AddCommand(memCmd)
}
