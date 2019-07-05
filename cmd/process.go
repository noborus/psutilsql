package cmd

import (
	"github.com/noborus/psutilsql"
	"github.com/noborus/trdsql"

	"github.com/spf13/cobra"
)

func processQuery(ex bool, query string) error {
	importer, err := psutilsql.NewProcessImporter(ex, query)
	if err != nil {
		return err
	}
	writer := trdsql.NewWriter(trdsql.OutFormat(outFormat()))
	trd := trdsql.NewTRDSQL(importer, trdsql.NewExporter(writer))
	err = trd.Exec(query)
	return err
}

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		defaultQuery := "SELECT * FROM process ORDER BY pid"

		var err error
		var ex bool
		if ex, err = cmd.PersistentFlags().GetBool("ex"); err != nil {
			return err
		}

		query := Query
		if query == "" {
			query = defaultQuery
		}
		return processQuery(ex, query)
	},
}

func init() {
	processCmd.PersistentFlags().BoolP("ex", "", false, "memory info ex")
	rootCmd.AddCommand(processCmd)
}
