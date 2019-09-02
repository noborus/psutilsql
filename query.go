package psutilsql

import (
	"strings"

	"github.com/noborus/trdsql"
)

const (
	psCPUTime         = "cputime"
	psCPUInfo         = "cpuinfo"
	psCPUPercent      = "cpupercent"
	psDiskPartition   = "diskpartition"
	psDiskUsage       = "diskusage"
	psDocker          = "docker"
	psHostInfo        = "hostinfo"
	psHostUser        = "hostuser"
	psHostTemperature = "hosttemperature"
	psLoadAvg         = "loadavg"
	psLoadMisc        = "loadmisc"
	psVirtualMemory   = "virtualmemory"
	psSwapMemory      = "swapmemory"
	psNet             = "net"
	psProcess         = "process"
	psProcessEx       = "processex"
	psWinservices     = "winservices"
)

func psutilReader(tableName string) Reader {
	var err error
	var reader Reader
	switch strings.ToLower(tableName) {
	case psCPUTime:
		reader, err = CPUTimeReader(false)
	case psCPUInfo:
		reader, err = CPUInfoReader()
	case psCPUPercent:
		reader, err = CPUPercentReader(false)
	case psDiskPartition:
		reader, err = DiskPartitionReader(true)
	case psDiskUsage:
		reader, err = DiskUsageReader("/")
	case psDocker:
		reader, err = DockerReader()
	case psHostInfo:
		reader, err = HostInfoReader()
	case psHostUser:
		reader, err = HostUsersReader()
	case psHostTemperature:
		reader, err = HostTemperatureReader()
	case psLoadAvg:
		reader, err = LoadAvgReader()
	case psLoadMisc:
		reader, err = LoadMiscReader()
	case psVirtualMemory:
		reader, err = VirtualMemoryReader()
	case psSwapMemory:
		reader, err = SwapMemoryReader()
	case psNet:
		reader, err = NetReader()
	case psProcess:
		reader, err = NewProcessReader(false)
	case psProcessEx:
		reader, err = NewProcessReader(true)
	case "pstable":
		reader, err = TableReader()
	default:
		reader = nil
	}
	if err != nil {
		return nil
	}
	return reader
}

func readerExec(reader Reader, query string, writer trdsql.Writer) error {
	importer, err := NewMultiImporter(reader)
	if err != nil {
		return err
	}
	trd := trdsql.NewTRDSQL(importer, trdsql.NewExporter(writer))
	err = trd.Exec(query)
	return err
}

// QueryExec actually executes the passed query and writes it to the writer.
func QueryExec(query string, writer trdsql.Writer) error {
	parsedQuery := trdsql.SQLFields(query)
	tables, _ := trdsql.TableNames(parsedQuery)
	if tables == nil {
		return nil
	}
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
	trd := trdsql.NewTRDSQL(importer, trdsql.NewExporter(writer))
	err = trd.Exec(query)
	return err
}
