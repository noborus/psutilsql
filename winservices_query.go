// +build windows

package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/winservices"
)

func WinservicesReader() (*trdsql.SliceReader, error) {
	v, err := winservices.Info()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader("winservices", v), nil
}

func WinservicesQuery(query string, out trdsql.Format) error {
	reader, err := WinservicesReader()
	if err != nil {
		return err
	}

	defaultQuery := "SELECT * FROM winservices"
	if query == "" {
		query = defaultQuery
	}
	return readerQuery(reader, query, out)
}
