package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/mem"
)

func VirtualMemoryReader() (*trdsql.SliceReader, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader(psVirtualMemory, v), nil
}

func SwapMemoryReader() (*trdsql.SliceReader, error) {
	v, err := mem.SwapMemory()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader(psSwapMemory, v), nil
}

func MEMQuery(memory bool, query string, out trdsql.Format) error {
	if memory {
		defaultQuery := "SELECT * FROM " + psVirtualMemory
		if query == "" {
			query = defaultQuery
		}
		reader, err := VirtualMemoryReader()
		if err != nil {
			return err
		}
		return readerQuery(reader, query, out)
	}

	defaultQuery := "SELECT * FROM " + psSwapMemory
	if query == "" {
		query = defaultQuery
	}
	reader, err := SwapMemoryReader()
	if err != nil {
		return err
	}
	return readerQuery(reader, query, out)
}
