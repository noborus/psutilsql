package psutilsql

import (
	"io"

	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/process"
)

type ProcessImporter struct {
	readers []*ProcessReader
}

func NewProcessImporter(ex bool) (*ProcessImporter, error) {
	pr, err := NewProcessReader(ex)
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
	data      [][]interface{}
}

func NewProcessReader(ex bool) (*ProcessReader, error) {
	pr := &ProcessReader{}
	pr.tableName = "process"
	if ex {
		ProcessColumn[MEMORYINFOEX].enable = true
		ProcessColumn[MEMORYINFO].enable = false
	} else {
		ProcessColumn[MEMORYINFOEX].enable = false
		ProcessColumn[MEMORYINFO].enable = true
	}
	for _, col := range ProcessColumn {
		if !col.enable {
			continue
		}
		pr.names = append(pr.names, col.names...)
		pr.types = append(pr.types, col.types...)
	}

	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}
	pr.data = make([][]interface{}, len(processes))
	for i, p := range processes {
		pr.data[i] = []interface{}{}
		for _, col := range ProcessColumn {
			if !col.enable {
				continue
			}
			pr.data[i] = append(pr.data[i], col.getFunc(p)...)
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
