package psutilsql

import (
	"fmt"
	"runtime"

	"github.com/noborus/trdsql"
)

// TableReader return table name as trdsql.SliceReader.
func TableReader() (*trdsql.SliceReader, error) {
	type tableNames struct {
		name string
	}
	tables := []tableNames{
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
	if runtime.GOOS != "windows" {
		tables = append(tables, tableNames{name: psWinservices})
	}
	return trdsql.NewSliceReader("pstable", tables), nil
}

func definitionQuery(tableName string, w trdsql.Writer) error {
	reader := psutilReader(tableName)
	if reader == nil {
		return fmt.Errorf("no such table")
	}
	tn, err := reader.TableName()
	if err != nil {
		return err
	}
	query := "SELECT * FROM " + tn + " LIMIT 0"
	return readerExec(reader, query, w)
}

// PSTableQuery executes SQL on tables.
func PSTableQuery(tableName string, query string, w trdsql.Writer) error {
	if len(tableName) > 0 {
		return definitionQuery(tableName, w)
	}
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
