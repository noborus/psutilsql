package psutilsql

import (
	"io"

	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/process"
)

type ProcessImporter struct {
	readers []*ProcessReader
}

func NewProcessImporter(ex bool, query string) (*ProcessImporter, error) {
	pr, err := NewProcessReader(ex, query)
	if err != nil {
		return nil, err
	}
	prs := []*ProcessReader{}
	prs = append(prs, pr)
	return &ProcessImporter{
		readers: prs,
	}, nil
}

func (i *ProcessImporter) Import(db *trdsql.DB, query string) (string, error) {
	for _, pr := range i.readers {
		names, err := pr.Names()
		if err != nil {
			return query, err
		}
		types, err := pr.Types()
		if err != nil {
			return query, err
		}
		err = db.CreateTable(pr.tableName, names, types, true)
		if err != nil {
			return query, err
		}
		err = db.Import(pr.tableName, names, pr)
		if err != nil {
			return query, err
		}
	}
	return query, nil
}

type ProcessReader struct {
	tableName string
	names     []string
	types     []string
	funcs     []func(p *process.Process) []interface{}
	data      [][]interface{}
}

func NewProcessReader(ex bool, query string) (*ProcessReader, error) {
	pr := &ProcessReader{}
	pr.tableName = "process"
	columns := []ColumnNum{PID, NAME, CPU, MEM, STATUS, START, USER, MEMORYINFO, COMMAND}
	if ex {
		columns = []ColumnNum{PID, NAME, CPU, MEM, STATUS, START, USER, MEMORYINFOEX, COMMAND}
	}
	for _, cn := range columns {
		col := ProcessColumn[cn]
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
			pr.data[i] = append(pr.data[i], getFunc(p)...)
		}
	}

	return pr, nil
}

func (p *ProcessReader) Names() ([]string, error) {
	return p.names, nil
}

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
