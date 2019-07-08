package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/disk"
)

func DiskPartitionQuery(all bool, query string, out trdsql.Format) error {
	defaultQuery := "SELECT * FROM DiskPartition"
	if query == "" {
		query = defaultQuery
	}
	v, err := disk.Partitions(all)
	if err != nil {
		return err
	}
	return SliceQuery(v, "DiskPartition", query, out)
}

func DiskUsage(usage string, query string, out trdsql.Format) error {
	defaultQuery := "SELECT * FROM DiskUsage"
	if query == "" {
		query = defaultQuery
	}
	v, err := disk.Usage(usage)
	if err != nil {
		return err
	}
	return SliceQuery(v, "DiskUsage", query, out)
}
