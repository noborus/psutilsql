package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/load"
)

func LoadAvgReader() (*trdsql.SliceReader, error) {
	v, err := load.Avg()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader("LoadAvg", v), nil
}

func LoadMiscReader() (*trdsql.SliceReader, error) {
	v, err := load.Misc()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader("LoadMisc", v), nil
}

func LoadQuery(misc bool, query string, out trdsql.Format) error {
	if misc {
		defaultQuery := "SELECT * FROM LoadMisc"
		if query == "" {
			query = defaultQuery
		}
		reader, err := LoadMiscReader()
		if err != nil {
			return err
		}
		return readerQuery(reader, query, out)
	}
	defaultQuery := "SELECT * FROM LoadAvg"
	if query == "" {
		query = defaultQuery
	}
	reader, err := LoadAvgReader()
	if err != nil {
		return err
	}
	return readerQuery(reader, query, out)
}
