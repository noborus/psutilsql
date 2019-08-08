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
	return trdsql.NewSliceReader(psDiskPartition, v), nil
}

func DiskUsageReader(usage string) (*trdsql.SliceReader, error) {
	v, err := disk.Usage(usage)
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader(psDiskUsage, v), nil
}

func DiskPartitionQuery(all bool, query string, w trdsql.Writer) error {
	reader, err := DiskPartitionReader(all)
	if err != nil {
		return err
	}
	defaultQuery := "SELECT * FROM " + psDiskPartition
	if query == "" {
		query = defaultQuery
	}
	return readerExec(reader, query, w)
}

func DiskUsageQuery(usage string, query string, w trdsql.Writer) error {
	reader, err := DiskUsageReader(usage)
	if err != nil {
		return err
	}

	defaultQuery := "SELECT * FROM " + psDiskUsage
	if query == "" {
		query = defaultQuery
	}
	return readerExec(reader, query, w)
}
