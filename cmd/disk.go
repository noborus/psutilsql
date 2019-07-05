package cmd

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/spf13/cobra"
)

// diskCmd represents the disk command
var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		defaultQuery := "SELECT * FROM disk"
		var err error
		var all bool
		if all, err = cmd.PersistentFlags().GetBool("all"); err != nil {
			return err
		}
		var usage string
 		if usage, err = cmd.PersistentFlags().GetString("usage"); err != nil {
			return err
		}
		var v interface{}
		if usage != "" {
			v, err = disk.Usage(usage)
		} else {
			v, err = disk.Partitions(all)
		}
		if err != nil {
			return err
		}
		query := Query
		if Query == "" {
			query = defaultQuery
		}
		return sliceQuery(v, "disk", query)
	},
}

func init() {
	diskCmd.PersistentFlags().BoolP("all", "a", false, "all disk")
	diskCmd.PersistentFlags().StringP("usage", "u", "", "file system usage")

	rootCmd.AddCommand(diskCmd)
}
