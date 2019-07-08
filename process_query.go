package psutilsql

import (
	"github.com/noborus/trdsql"
)

func ProcessQuery(ex bool, query string, out trdsql.Format) error {
	defaultQuery := "SELECT * FROM process ORDER BY pid"
	if query == "" {
		query = defaultQuery
	}

	importer, err := NewProcessImporter(ex, query)
	if err != nil {
		return err
	}
	writer := trdsql.NewWriter(trdsql.OutFormat(out))
	trd := trdsql.NewTRDSQL(importer, trdsql.NewExporter(writer))
	err = trd.Exec(query)
	return err
}

