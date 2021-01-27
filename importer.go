package psutilsql

import (
	"context"

	"github.com/noborus/trdsql"
)

// MultiImporter is a structure for importing multiple readers.
type MultiImporter struct {
	readers []Reader
}

// Reader is an interface that can be passed to MultiImporter.
type Reader interface {
	TableName() (string, error)
	Names() ([]string, error)
	Types() ([]string, error)
	PreReadRow() [][]interface{}
	ReadRow([]interface{}) ([]interface{}, error)
}

// NewMultiImporter takes multiple readers as arguments and returns a MultiImporter.
func NewMultiImporter(readers ...Reader) (*MultiImporter, error) {
	r := make([]Reader, len(readers))
	copy(r, readers)
	return &MultiImporter{
		readers: readers,
	}, nil
}

// ImportContext executes import.
func (i *MultiImporter) ImportContext(ctx context.Context, db *trdsql.DB, query string) (string, error) {
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
		err = db.CreateTableContext(ctx, tableName, names, types, true)
		if err != nil {
			return query, err
		}
		err = db.ImportContext(ctx, tableName, names, r)
		if err != nil {
			return query, err
		}
	}
	return query, nil
}

// Import executes import.
func (i *MultiImporter) Import(db *trdsql.DB, query string) (string, error) {
	ctx := context.Background()
	return i.ImportContext(ctx, db, query)
}
