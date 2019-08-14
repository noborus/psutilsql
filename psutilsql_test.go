package psutilsql

import (
	"io/ioutil"

	"github.com/noborus/trdsql"
)

func nullWriter() trdsql.Writer {
	return trdsql.NewWriter(trdsql.OutStream(ioutil.Discard))
}
