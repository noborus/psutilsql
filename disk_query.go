package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/disk"
)

func DiskPartitionReader(all bool) (*trdsql.SliceReader, error) {
	v, err := disk.Partitions(all)
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader("DiskPartition", v), nil
}

func DiskUsageReader(usage string) (*trdsql.SliceReader, error) {
	v, err := disk.Usage(usage)
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader("DiskUsage", v), nil
}

func DiskPartitionQuery(all bool, query string, out trdsql.Format) error {
	reader, err := DiskPartitionReader(all)
	if err != nil {
		return err
	}
	defaultQuery := "SELECT * FROM DiskPartition"
	if query == "" {
		query = defaultQuery
	}
	return readerQuery(reader, query, out)
}

func DiskUsageQuery(usage string, query string, out trdsql.Format) error {
	reader, err := DiskUsageReader(usage)
	if err != nil {
		return err
	}

	defaultQuery := "SELECT * FROM DiskUsage"
	if query == "" {
		query = defaultQuery
	}
	return readerQuery(reader, query, out)
}
