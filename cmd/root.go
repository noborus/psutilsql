package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/noborus/psutilsql"

	"github.com/noborus/trdsql"
	"github.com/spf13/cobra"
)

var (
	ErrNoQuery = errors.New("require query")
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
			return ErrNoQuery
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
	rootCmd.PersistentFlags().StringVarP(&OutFormat, "OutFormat", "o", "at", "output format=at|csv|ltsv|json|jsonl|tbln|raw|md|vf|yaml")
	rootCmd.PersistentFlags().StringVarP(&Delimiter, "Delimiter", "d", ",", "output delimiter (CSV only)")
	rootCmd.PersistentFlags().BoolVarP(&Header, "Header", "O", false, "output header (CSV only)")
	rootCmd.PersistentFlags().StringVarP(&Query, "Query", "q", "", "query")
}

func outFormat() trdsql.Writer {
	format := trdsql.OutputFormat(strings.ToLower(OutFormat))
	w := trdsql.NewWriter(
		trdsql.OutFormat(format),
		trdsql.OutDelimiter(Delimiter),
		trdsql.OutHeader(Header),
	)
	return w
}
