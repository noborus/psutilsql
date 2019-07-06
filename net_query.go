package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/net"
)

func NetQuery(query string, out trdsql.Format) error {
	defaultQuery := "SELECT * FROM net"
	var v interface{}

	var err error
	v, err = net.Connections("all")
	if err != nil {
		return err
	}
	if query == "" {
		query = defaultQuery
	}
	return SliceQuery(v, "net", query, out)

}
