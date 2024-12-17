//go:build windows
// +build windows

package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/v4/winservices"
)

// WinservicesReader returns winservices.ListServices as trdsql.SliceReader.
func WinservicesReader() (*trdsql.SliceReader, error) {
	v, err := winservices.ListServices()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader("winservices", v), nil
}

// WinservicesQuery executes SQL on winservices.ListServices.
func WinservicesQuery(query string, w trdsql.Writer) error {
	reader, err := WinservicesReader()
	if err != nil {
		return err
	}

	defaultQuery := "SELECT * FROM " + psWinservices
	if query == "" {
		query = defaultQuery
	}
	return readerExec(reader, query, w)
}
