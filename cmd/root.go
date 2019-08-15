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
		if len(Query) == 0 && len(args) == 0 {
			return fmt.Errorf("require query")
		}
		return psutilsql.QueryExec(Query, outFormat())
	},
}
var (
	// Version represents the version
	Version string
	// Revision set "git rev-parse --short HEAD"
	Revision string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string, revision string) {
	Version = version
	Revision = revision
	cmd, _, err := rootCmd.Find(os.Args[1:])
	if err != nil || cmd == nil {
		// Not found
		args := append([]string{"query"}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
	if err := rootCmd.Execute(); err != nil {
		rootCmd.SetOutput(os.Stderr)
		rootCmd.Println(err)
		os.Exit(1)
	}
}

// OutFormat is an output format specification.
var OutFormat string

// Header is an output header specification(CSV and RAW only).
var Header bool

// Delimiter is a delimiter specification (CSV ans RAW only).
var Delimiter string

// Query is SQL specification.
var Query string

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&OutFormat, "OutFormat", "o", "at", "output format=at|csv|ltsv|json|tbln|raw|md|vf")
	rootCmd.PersistentFlags().StringVarP(&Delimiter, "Delimiter", "d", ",", "output delimiter (CSV only)")
	rootCmd.PersistentFlags().BoolVarP(&Header, "Header", "O", false, "output header (CSV only)")
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
