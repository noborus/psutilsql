package psutilsql

import (
	"github.com/noborus/trdsql"
)

func TableReader() (*trdsql.SliceReader, error) {
	type tableName struct {
		name string
	}
	tables := []tableName{
		{name: psCPUTime},
		{name: psCPUInfo},
		{name: psCPUPercent},
		{name: psDiskPartition},
		{name: psDiskUsage},
		{name: psDocker},
		{name: psHostInfo},
		{name: psHostUser},
		{name: psHostTemperature},
		{name: psLoadAvg},
		{name: psLoadMisc},
		{name: psVirtualMemory},
		{name: psSwapMemory},
		{name: psNet},
		{name: psProcess},
		{name: psProcessEx},
	}
	return trdsql.NewSliceReader("pstable", tables), nil
}

func PSTableQuery(query string, w trdsql.Writer) error {
	reader, err := TableReader()
	if err != nil {
		return err
	}
	defaultQuery := "SELECT * FROM pstable ORDER BY name"
	if query == "" {
		query = defaultQuery
	}
	return readerExec(reader, query, w)
}
