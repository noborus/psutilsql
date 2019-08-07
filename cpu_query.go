package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/cpu"
)

func CPUTimeReader(total bool) (*trdsql.SliceReader, error) {
	v, err := cpu.Times(!total)
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader("CPUTime", v), nil
}

func CPUInfoReader() (*trdsql.SliceReader, error) {
	v, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader("CPUInfo", v), nil
}

func CPUPercentReader(total bool) (*trdsql.SliceReader, error) {
	v, err := cpu.Percent(0, !total)
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader("CPUPercent", v), nil
}

func CPUTimeQuery(total bool, query string, out trdsql.Format) error {
	reader, err := CPUTimeReader(!total)
	if err != nil {
		return err
	}
	defaultQuery := "SELECT * FROM CPUTime ORDER BY cpu"
	if query == "" {
		query = defaultQuery
	}
	return readerQuery(reader, query, out)
}

func CPUInfoQuery(query string, out trdsql.Format) error {
	reader, err := CPUInfoReader()
	if err != nil {
		return err
	}
	defaultQuery := "SELECT * FROM CPUInfo ORDER BY cpu"
	if query == "" {
		query = defaultQuery
	}
	return readerQuery(reader, query, out)
}

func CPUPercentQuery(total bool, query string, out trdsql.Format) error {
	reader, err := CPUPercentReader(total)
	if err != nil {
		return err
	}
	defaultQuery := "SELECT * FROM CPUPercent"
	if query == "" {
		query = defaultQuery
	}
	return readerQuery(reader, query, out)
}
