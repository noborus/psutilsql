package psutilsql

import (
	"github.com/noborus/trdsql"
)

type Importer struct {
	tableNames []string
	readers    []trdsql.Reader
}

func NewImporter(tableNames []string, readers []trdsql.Reader) (*Importer, error) {
	return &Importer{
		tableNames: tableNames,
		readers:    readers,
	}, nil
}

func (i *Importer) Import(db *trdsql.DB, query string) (string, error) {
	for n, pr := range i.readers {
		names, err := pr.Names()
		if err != nil {
			return query, err
		}
		types, err := pr.Types()
		if err != nil {
			return query, err
		}
		err = db.CreateTable(i.tableNames[n], names, types, true)
		if err != nil {
			return query, err
		}
		err = db.Import(i.tableNames[n], names, pr)
		if err != nil {
			return query, err
		}
	}
	return query, nil
}
