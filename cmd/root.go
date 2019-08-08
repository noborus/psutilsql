package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/noborus/psutilsql"

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
	RunE: func(c *cobra.Command, args []string) error {
		return psutilsql.QueryExec(Query, outFormat())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	if err != nil || cmd == nil {
		// Not found
		args := append([]string{"query"}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
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
	rootCmd.PersistentFlags().BoolVarP(&Header, "Header", "O", false, "output header (CSV only)")
	rootCmd.PersistentFlags().StringVarP(&Delimiter, "Delimiter", "d", ",", "output header (CSV only)")
	rootCmd.PersistentFlags().StringVarP(&Query, "Query", "q", "", "query")
}

func outFormat() trdsql.Writer {
	var format trdsql.Format
	switch strings.ToUpper(OutFormat) {
	case "CSV":
		format = trdsql.CSV
	case "LTSV":
		format = trdsql.LTSV
	case "JSON":
		format = trdsql.JSON
	case "TBLN":
		format = trdsql.TBLN
	case "RAW":
		format = trdsql.RAW
	case "MD":
		format = trdsql.MD
	case "AT":
		format = trdsql.AT
	case "VF":
		format = trdsql.VF
	default:
		format = trdsql.AT
	}
	w := trdsql.NewWriter(
		trdsql.OutFormat(format),
		trdsql.OutDelimiter(Delimiter),
		trdsql.OutHeader(Header),
	)
	return w
}
