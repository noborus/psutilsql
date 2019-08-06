package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/noborus/trdsql"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "psutilsql",
	Short: "SQL for running processes and system utilization",
	Long: `SQL for running processes and system utilization.

SQL can be executed on the information acquired using gopsutil library.
Default SQL is provided, so you can omit SQL if you select a command.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var OutFormat string
var Header bool
var Delimiter string
var Query string

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&OutFormat, "OutFormat", "o", "", "output format")
	rootCmd.PersistentFlags().BoolVarP(&Header, "Header", "", false, "output header (CSV only)")
	rootCmd.PersistentFlags().StringVarP(&Delimiter, "Delimiter", "d", "", "output header (CSV only)")
	rootCmd.PersistentFlags().StringVarP(&Query, "Query", "q", "", "query")
}

func outFormat() trdsql.Format {
	switch strings.ToUpper(OutFormat) {
	case "CSV":
		return trdsql.CSV
	case "LTSV":
		return trdsql.LTSV
	case "JSON":
		return trdsql.JSON
	case "TBLN":
		return trdsql.TBLN
	case "RAW":
		return trdsql.RAW
	case "MD":
		return trdsql.MD
	case "AT":
		return trdsql.AT
	case "VF":
		return trdsql.VF
	}
	return trdsql.AT
}
