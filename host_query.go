package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/host"
)

func HostQuery(tempera bool, users bool, query string, out trdsql.Format) error {
	defaultQuery := "SELECT * FROM host"
	if query == "" {
		query = defaultQuery
	}

	var err error
	var v interface{}
	if tempera {
		v, err = host.SensorsTemperatures()
	} else if users {
		v, err = host.Users()
	} else {
		v, err = host.Info()
	}
	if err != nil {
		return err
	}
	return SliceQuery(v, "host", query, out)
}
