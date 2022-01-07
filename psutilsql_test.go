package psutilsql

import (
	"io"

	"github.com/noborus/trdsql"
)

func nullWriter() trdsql.Writer {
	return trdsql.NewWriter(trdsql.OutStream(io.Discard))
}
