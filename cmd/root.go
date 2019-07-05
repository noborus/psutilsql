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
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
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
var Query string

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&OutFormat, "OutFormat", "o", "", "output format")
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

func sliceQuery(slice interface{}, tableName string, query string) error {
	// trdsql.EnableDebug()
	importer := trdsql.NewSliceImporter(tableName, slice)
	writer := trdsql.NewWriter(trdsql.OutFormat(outFormat()))
	trd := trdsql.NewTRDSQL(importer, trdsql.NewExporter(writer))
	err := trd.Exec(query)
	return err
}
