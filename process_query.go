package psutilsql

import (
	"io"

	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/process"
)

func GetProccessReader(ex bool, query string) ([]string, []trdsql.Reader, error) {
	defaultQuery := "SELECT * FROM process ORDER BY pid"
	if query == "" {
		query = defaultQuery
	}
	reader, err := NewProcessReader(ex, query)
	if err != nil {
		return nil, nil, err
	}
	tableNames := []string{"process"}
	readers := []trdsql.Reader{}
	readers = append(readers, reader)
	return tableNames, readers, nil
}

func ProcessQuery(ex bool, query string, out trdsql.Format) error {
	if query == "" {
		query = "SELECT * FROM process ORDER BY pid"
	}

	tableNames, readers, err := GetProccessReader(ex, query)
	if err != nil {
		return err
	}

	importer, err := NewImporter(tableNames, readers)
	if err != nil {
		return err
	}
	writer := trdsql.NewWriter(trdsql.OutFormat(out))
	trd := trdsql.NewTRDSQL(importer, trdsql.NewExporter(writer))
	err = trd.Exec(query)
	return err
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
