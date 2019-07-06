package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/cpu"
)

func CPUQuery(percent bool, info bool, total bool, query string, out trdsql.Format) error {
	defaultQuery := "SELECT * FROM cpu ORDER BY cpu"
	var v interface{}
	var err error
	if percent {
		defaultQuery = "SELECT * FROM cpu"
		v, err = cpu.Percent(0, !total)
	} else if info {
		v, err = cpu.Info()
	} else {
		v, err = cpu.Times(!total)
	}
	if err != nil {
		return err
	}
	if query == "" {
		query = defaultQuery
	}
	return SliceQuery(v, "cpu", query, out)
}
