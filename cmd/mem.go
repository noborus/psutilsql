package cmd

import (
	"github.com/shirou/gopsutil/mem"

	"github.com/spf13/cobra"
)

// memCmd represents the mem command
var memCmd = &cobra.Command{
	Use:   "mem",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		defaultQuery := "SELECT * FROM mem"

		v, err := mem.VirtualMemory()
		if err != nil {
			return err
		}
		query := Query
		if query == "" {
			query = defaultQuery
		}
		return sliceQuery(v, "mem", query)
	},
}

func init() {
	rootCmd.AddCommand(memCmd)
}
