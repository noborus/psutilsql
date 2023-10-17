package psutilsql

import "errors"

var (
	ErrNotSupport  = errors.New("not support")
	ErrNoSuchTable = errors.New("no such table")
)
