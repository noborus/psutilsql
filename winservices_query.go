// +build windows

package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/winservices"
)

func WinservicesReader() (*trdsql.SliceReader, error) {
	v, err := winservices.ListServices()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader("winservices", v), nil
}

func WinservicesQuery(query string, w trdsql.Writer) error {
	reader, err := WinservicesReader()
	if err != nil {
		return err
	}

	defaultQuery := "SELECT * FROM winservices"
	if query == "" {
		query = defaultQuery
	}
	return readerExec(reader, query, w)
}
