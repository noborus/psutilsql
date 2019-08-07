package cmd

import (
	"github.com/noborus/psutilsql"

	"github.com/spf13/cobra"
)

// diskCmd represents the disk command
var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "DISK information",
	Long: `DISK information
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var all bool
		if all, err = cmd.PersistentFlags().GetBool("all"); err != nil {
			return err
		}
		var usage string
		if usage, err = cmd.PersistentFlags().GetString("usage"); err != nil {
			return err
		}
		if usage != "" {
			return psutilsql.DiskUsageQuery(usage, Query, outFormat())
		}
		return psutilsql.DiskPartitionQuery(all, Query, outFormat())
	},
}

func init() {
	diskCmd.PersistentFlags().BoolP("all", "a", false, "all disk")
	diskCmd.PersistentFlags().StringP("usage", "u", "", "file system usage")

	rootCmd.AddCommand(diskCmd)
}
