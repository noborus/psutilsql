// +build windows

package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/winservices"
)

func WinservicesQuery(query string, out trdsql.Format) error {
	defaultQuery := "SELECT * FROM winservices"

	v, err := winservices.Info()
	if err != nil {
		return err
	}
	if query == "" {
		query = defaultQuery
	}
	return SliceQuery(v, "winservices", query, out)

}
