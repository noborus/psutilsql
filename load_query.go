package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/load"
)

func LoadReader(misc bool) (*trdsql.SliceReader, error) {
	var v interface{}
	var err error
	if misc {
		v, err = load.Misc()
	} else {
		v, err = load.Avg()
	}
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader("load", v), nil
}

func LoadQuery(misc bool, query string, out trdsql.Format) error {
	reader, err := LoadReader(misc)
	if err != nil {
		return err
	}

	defaultQuery := "SELECT * FROM load"
	if query == "" {
		query = defaultQuery
	}
	return readerQuery(reader, query, out)
}
