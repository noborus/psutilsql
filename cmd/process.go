package cmd

import (
	"github.com/noborus/psutilsql"
	"github.com/spf13/cobra"
)

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "process information",
	Long: `process information.
+-----+------+-----+-----+--------+-------+------+-----+-----+------+-------+--------+------+---------+
| pid | name | CPU | MEM | STATUS | START | USER | RSS | VMS | Data | Stack | locked | Swap | COMMAND |
+-----+------+-----+-----+--------+-------+------+-----+-----+------+-------+--------+------+---------+

Option ex gets the result of platform dependend memory information.
+-----+------+-----+-----+--------+-------+------+-----+-----+--------+------+-----+------+-------+---------+
| pid | name | CPU | MEM | STATUS | START | USER | RSS | VMS | Shared | Text | Lib | Data | Dirty | COMMAND |
+-----+------+-----+-----+--------+-------+------+-----+-----+--------+------+-----+------+-------+---------+
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var ex bool
		if ex, err = cmd.PersistentFlags().GetBool("ex"); err != nil {
			return err
		}
		return psutilsql.ProcessQuery(ex, Query, outFormat())
	},
}

func init() {
	processCmd.PersistentFlags().BoolP("ex", "", false, "memory info ex")
	rootCmd.AddCommand(processCmd)
}
