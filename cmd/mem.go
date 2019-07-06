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
