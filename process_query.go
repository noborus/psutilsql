package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/process"
)

func ProcessQuery(ex bool, query string, out trdsql.Format) error {
	if query == "" {
		query = "SELECT * FROM process ORDER BY pid"
	}
	v, err := process.Processes()
	if err != nil {
		return err
	}
	return SliceQuery(v, "process", query, out)
}
