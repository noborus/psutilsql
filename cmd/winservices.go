// +build windows

package cmd

import (
	"github.com/shirou/gopsutil/winservices"

	"github.com/spf13/cobra"
)

// winservicesCmd represents the winservices command
var winservicesCmd = &cobra.Command{
	Use:   "winservices",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		v, err := winservices.Info()
		if err != nil {
			return err
		}
		query := "SELECT * FROM process"
		if Query != "" {
			query = Query
		}
		return sliceQuery(v, "process", query)
	},
}

func init() {
	rootCmd.AddCommand(winservicesCmd)
}
