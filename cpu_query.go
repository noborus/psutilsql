package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/cpu"
)

func CPUTimeQuery(total bool, query string, out trdsql.Format) error {
	defaultQuery := "SELECT * FROM CPUTime ORDER BY cpu"
	if query == "" {
		query = defaultQuery
	}
	v, err := cpu.Times(!total)
	if err != nil {
		return err
	}
	return SliceQuery(v, "CPUTime", query, out)
}

func CPUInfoQuery(query string, out trdsql.Format) error {
	defaultQuery := "SELECT * FROM CPUInfo ORDER BY cpu"
	if query == "" {
		query = defaultQuery
	}
	v, err := cpu.Info()
	if err != nil {
		return err
	}
	return SliceQuery(v, "CPUInfo", query, out)
}

func CPUPercentQuery(total bool, query string, out trdsql.Format) error {
	defaultQuery := "SELECT * FROM CPUPercent"
	if query == "" {
		query = defaultQuery
	}
	v, err := cpu.Percent(0, !total)
	if err != nil {
		return err
	}
	return SliceQuery(v, "CPUPercent", query, out)
}
