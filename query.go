package psutilsql

import (
	"strings"

	"github.com/noborus/trdsql"
)

func QueryImport(query string, out trdsql.Format) error {
	tables := trdsql.TableNames(query)
	var readers []Reader
	for _, table := range tables {
		reader := psutilReader(table)
		if reader != nil {
			readers = append(readers, reader)
		}
	}
	importer, err := NewMultiImporter(readers...)
	if err != nil {
		return err
	}
	writer := trdsql.NewWriter(trdsql.OutFormat(out))
	trd := trdsql.NewTRDSQL(importer, trdsql.NewExporter(writer))
	err = trd.Exec(query)
	return err
}

func psutilReader(tableName string) Reader {
	var err error
	var reader Reader
	switch strings.ToLower(tableName) {
	case "cputime":
		reader, err = CPUTimeReader(false)
	case "cpuinfo":
		reader, err = CPUInfoReader()
	case "cpupercent":
		reader, err = CPUPercentReader(false)
	case "diskpartition":
		reader, err = DiskPartitionReader(false)
	case "diskusage":
		reader, err = DiskUsageReader("/")
	case "docker":
		reader, err = DockerReader()
	case "hostinfo":
		reader, err = HostInfoReader()
	case "hostuser":
		reader, err = HostUserReader()
	case "hosttemperature":
		reader, err = HostTemperatureReader()
	case "loadavg":
		reader, err = LoadAvgReader()
	case "loadmisc":
		reader, err = LoadMiscReader()
	case "virtualmemory":
		reader, err = VirtualMemoryReader()
	case "swapmemory":
		reader, err = SwapMemoryReader()
	case "net":
		reader, err = NetReader()
	case "process":
		reader, err = NewProcessReader(false)
	case "processex":
		reader, err = NewProcessReader(true)
	default:
		reader = nil
	}
	if err != nil {
		return nil
	}
	return reader
}
