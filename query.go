package psutilsql

import (
	"strings"

	"github.com/noborus/trdsql"
)

func QueryImport(query string, out trdsql.Format) error {
	tables := TableNames(query)
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

// Copy from trdsql...
func TableNames(query string) []string {
	var tables []string
	var tableFlag, frontFlag bool
	word := sqlFields(query)
	for i, w := range word {
		frontFlag = false
		switch {
		case strings.ToUpper(w) == "FROM" || strings.ToUpper(w) == "JOIN":
			tableFlag = true
			frontFlag = true
		case isSQLKeyWords(w):
			tableFlag = false
		case w == ",":
			frontFlag = true
		default:
			frontFlag = false
		}
		if n := i + 1; n < len(word) && tableFlag && frontFlag {
			if t := word[n]; len(t) > 0 {
				if t[len(t)-1] == ')' {
					t = t[:len(t)-1]
				}
				if !isSQLKeyWords(t) {
					tables = append(tables, t)
				}
			}
		}
	}
	return tables
}

func sqlFields(query string) []string {
	parsed := []string{}
	buf := ""
	var singleQuoted, doubleQuoted, backQuote bool
	for _, r := range query {
		switch r {
		case ' ', '\t', '\r', '\n', ',', ';', '=':
			if !singleQuoted && !doubleQuoted && !backQuote {
				if buf != "" {
					parsed = append(parsed, buf)
					buf = ""
				}
				if r == ',' {
					parsed = append(parsed, ",")
				}
			} else {
				buf += string(r)
			}
			continue
		case '\'':
			if !doubleQuoted && !backQuote {
				singleQuoted = !singleQuoted
			}
		case '"':
			if !singleQuoted && !backQuote {
				doubleQuoted = !doubleQuoted
			}
		case '`':
			if !singleQuoted && !doubleQuoted {
				backQuote = !backQuote
			}
		}
		buf += string(r)
	}
	parsed = append(parsed, buf)
	return parsed
}

func isSQLKeyWords(str string) bool {
	switch strings.ToUpper(str) {
	case "WHERE", "GROUP", "HAVING", "WINDOW", "UNION", "ORDER", "LIMIT", "OFFSET", "FETCH",
		"FOR", "LEFT", "RIGHT", "CROSS", "INNER", "FULL", "LETERAL", "(SELECT":
		return true
	}
	return false
}
