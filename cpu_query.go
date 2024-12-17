package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/v4/cpu"
)

// CPUTimeReader returns cpu.Times result as trdsql.SliceReader.
func CPUTimeReader(total bool) (*trdsql.SliceReader, error) {
	v, err := cpu.Times(!total)
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader(psCPUTime, v), nil
}

// CPUInfoReader returns cpu.Info result as trdsql.SliceReader.
func CPUInfoReader() (*trdsql.SliceReader, error) {
	v, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader(psCPUInfo, v), nil
}

// CPUPercentReader returns cpu.Percent result as trdsql.SliceReader.
func CPUPercentReader(total bool) (*trdsql.SliceReader, error) {
	v, err := cpu.Percent(0, !total)
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader(psCPUPercent, v), nil
}

// CPUTimeQuery executes SQL on cpu.Time.
func CPUTimeQuery(total bool, query string, w trdsql.Writer) error {
	reader, err := CPUTimeReader(total)
	if err != nil {
		return err
	}
	defaultQuery := "SELECT * FROM " + psCPUTime + " ORDER BY cpu"
	if query == "" {
		query = defaultQuery
	}
	return readerExec(reader, query, w)
}

// CPUInfoQuery executes SQL on cpu.Info.
func CPUInfoQuery(query string, w trdsql.Writer) error {
	reader, err := CPUInfoReader()
	if err != nil {
		return err
	}
	defaultQuery := "SELECT * FROM " + psCPUInfo + " ORDER BY cpu"
	if query == "" {
		query = defaultQuery
	}
	return readerExec(reader, query, w)
}

// CPUPercentQuery executes SQL on cpu.Percent.
func CPUPercentQuery(total bool, query string, w trdsql.Writer) error {
	reader, err := CPUPercentReader(total)
	if err != nil {
		return err
	}
	defaultQuery := "SELECT * FROM " + psCPUPercent
	if query == "" {
		query = defaultQuery
	}
	return readerExec(reader, query, w)
}
