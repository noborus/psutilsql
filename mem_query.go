package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/mem"
)

// VirtualMemoryReader returns mem.VirtualMemory result as trdsql.SliceReader.
func VirtualMemoryReader() (*trdsql.SliceReader, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader(psVirtualMemory, v), nil
}

// SwapMemoryReader returns mem.SwapMemory result as trdsql.SliceReader.
func SwapMemoryReader() (*trdsql.SliceReader, error) {
	v, err := mem.SwapMemory()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader(psSwapMemory, v), nil
}

// MEMQuery executes SQL on mem.VirtualMemory or mem.SwapMemory.
func MEMQuery(memory bool, query string, w trdsql.Writer) error {
	if memory {
		defaultQuery := "SELECT * FROM " + psVirtualMemory
		if query == "" {
			query = defaultQuery
		}
		reader, err := VirtualMemoryReader()
		if err != nil {
			return err
		}
		return readerExec(reader, query, w)
	}

	defaultQuery := "SELECT * FROM " + psSwapMemory
	if query == "" {
		query = defaultQuery
	}
	reader, err := SwapMemoryReader()
	if err != nil {
		return err
	}
	return readerExec(reader, query, w)
}
