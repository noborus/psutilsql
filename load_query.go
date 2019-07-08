package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/load"
)

func LoadQuery(misc bool, query string, out trdsql.Format) error {
	defaultQuery := "SELECT * FROM load"
	if query == "" {
		query = defaultQuery
	}

	var v interface{}
	var err error
	if misc {
		v, err = load.Misc()
	} else {
		v, err = load.Avg()
	}
	if err != nil {
		return err
	}
	return SliceQuery(v, "load", query, out)
}