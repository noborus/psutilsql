package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/mem"
)

func MEMQuery(memory bool, query string, out trdsql.Format) error {
	defaultQuery := "SELECT * FROM mem"
	var v interface{}
	var err error
	if memory {
		v, err = mem.VirtualMemory()
	} else {
		v, err = mem.SwapMemory()
	}
	if err != nil {
		return err
	}
	if query == "" {
		query = defaultQuery
	}
	return SliceQuery(v, "mem", query, out)

}
