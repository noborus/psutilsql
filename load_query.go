package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/load"
)

// LoadAvgReader returns load.Avg result as trdsql.SliceReader.
func LoadAvgReader() (*trdsql.SliceReader, error) {
	v, err := load.Avg()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader(psLoadAvg, v), nil
}

// LoadMiscReader returns load.Misc result as trdsql.SliceReader.
func LoadMiscReader() (*trdsql.SliceReader, error) {
	v, err := load.Misc()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader(psLoadMisc, v), nil
}

// LoadQuery executes SQL on Load.Avg or Load.Misc.
func LoadQuery(misc bool, query string, w trdsql.Writer) error {
	if misc {
		defaultQuery := "SELECT * FROM " + psLoadMisc
		if query == "" {
			query = defaultQuery
		}
		reader, err := LoadMiscReader()
		if err != nil {
			return err
		}
		return readerExec(reader, query, w)
	}
	defaultQuery := "SELECT * FROM " + psLoadAvg
	if query == "" {
		query = defaultQuery
	}
	reader, err := LoadAvgReader()
	if err != nil {
		return err
	}
	return readerExec(reader, query, w)
}
