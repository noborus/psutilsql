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

		v, err := disk.Partitions(false)
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
	rootCmd.AddCommand(diskCmd)
}
