package psutilsql

import (
	"github.com/noborus/trdsql"
)

type MultiImporter struct {
	readers []Reader
}

type Reader interface {
	TableName() (string, error)
	Names() ([]string, error)
	Types() ([]string, error)
	PreReadRow() [][]interface{}
	ReadRow([]interface{}) ([]interface{}, error)
}

func NewMultiImporter(readers ...Reader) (*MultiImporter, error) {
	r := make([]Reader, len(readers))
	copy(r, readers)
	return &MultiImporter{
		readers: readers,
	}, nil
}

func (i *MultiImporter) Import(db *trdsql.DB, query string) (string, error) {
	for _, r := range i.readers {
		names, err := r.Names()
		if err != nil {
			return query, err
		}
		types, err := r.Types()
		if err != nil {
			return query, err
		}
		tableName, err := r.TableName()
		if err != nil {
			return query, err
		}
		err = db.CreateTable(tableName, names, types, true)
		if err != nil {
			return query, err
		}
		err = db.Import(tableName, names, r)
		if err != nil {
			return query, err
		}
	}
	return query, nil
}
