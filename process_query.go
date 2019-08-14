package psutilsql

import (
	"fmt"
	"io"
	"runtime"

	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/process"
)

// NewProcessReader returns process.Processes result as ProcessReader.
func NewProcessReader(ex bool) (*ProcessReader, error) {
	pr := &ProcessReader{}
	pr.tableName = psProcess
	columns := []pColumnNum{PID, NAME, CPU, MEM, STATUS, START, USER, MEMORYINFO, COMMAND}
	if ex {
		pr.tableName = psProcessEx
		if runtime.GOOS == "linux" {
			columns = []pColumnNum{PID, NAME, CPU, MEM, STATUS, START, USER, MEMORYINFOEX, COMMAND}
		} else {
			return nil, fmt.Errorf("not supported")
		}
	}
	for _, cn := range columns {
		col := processColumn[cn]
		pr.names = append(pr.names, col.names...)
		pr.types = append(pr.types, col.types...)
		pr.funcs = append(pr.funcs, col.getFunc)
	}

	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}
	pr.data = make([][]interface{}, len(processes))
	for i, p := range processes {
		pr.data[i] = []interface{}{}
		for _, getFunc := range pr.funcs {
			v := getFunc(p)
			pr.data[i] = append(pr.data[i], v...)
		}
	}

	return pr, nil
}

// The ProcessReader structure represents a process
// and satisfies the trdsql.Reader interface.
type ProcessReader struct {
	tableName string
	names     []string
	types     []string
	funcs     []func(p *process.Process) []interface{}
	data      [][]interface{}
}

// TableName returns TableName.
func (p *ProcessReader) TableName() (string, error) {
	return p.tableName, nil
}

// Names returns column names.
func (p *ProcessReader) Names() ([]string, error) {
	return p.names, nil
}

// Types returns column types.
func (p *ProcessReader) Types() ([]string, error) {
	return p.types, nil
}

// PreReadRow is returns entity of the data.
func (p *ProcessReader) PreReadRow() [][]interface{} {
	return p.data
}

// ReadRow only returns EOF.
func (p *ProcessReader) ReadRow(row []interface{}) ([]interface{}, error) {
	return nil, io.EOF
}

// ProcessQuery executes SQL on process.Processes.
func ProcessQuery(ex bool, query string, w trdsql.Writer) error {
	reader, err := NewProcessReader(ex)
	if err != nil {
		return err
	}
	defaultQuery := "SELECT * FROM " + psProcess + " ORDER BY pid"
	if ex {
		defaultQuery = "SELECT * FROM " + psProcessEx + " ORDER BY pid"
	}
	if query == "" {
		query = defaultQuery
	}

	return readerExec(reader, query, w)
}
