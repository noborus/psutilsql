package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/mem"
)

func MEMReader(memory bool) (*trdsql.SliceReader, error) {
	var v interface{}
	var err error
	if memory {
		v, err = mem.VirtualMemory()
	} else {
		v, err = mem.SwapMemory()
	}
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader("mem", v), nil

}

func MEMQuery(memory bool, query string, out trdsql.Format) error {
	reader, err := MEMReader(memory)
	if err != nil {
		return err
	}

	defaultQuery := "SELECT * FROM mem"
	if query == "" {
		query = defaultQuery
	}
	return readerQuery(reader, query, out)
}
