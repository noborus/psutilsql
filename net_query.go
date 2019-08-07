package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/net"
)

func NetReader() (*trdsql.SliceReader, error) {
	conns, err := net.Connections("all")
	if err != nil {
		return nil, err
	}
	type wrapConnection struct {
		Fd        uint32
		Family    uint32
		Type      uint32
		LaddrIP   string
		LaddrPort uint32
		RaddrIP   string
		RaddrPort uint32
		status    string
		Uids      []int32
		Pid       int32
	}
	data := make([]wrapConnection, len(conns))
	for i, conn := range conns {
		c := wrapConnection{}
		c.Fd = conn.Fd
		c.Family = conn.Family
		c.Type = conn.Type
		c.LaddrIP = conn.Laddr.IP
		c.LaddrPort = conn.Laddr.Port
		c.RaddrIP = conn.Raddr.IP
		c.RaddrPort = conn.Raddr.Port
		c.status = conn.Status
		c.Uids = conn.Uids
		c.Pid = conn.Pid
		data[i] = c
	}
	return trdsql.NewSliceReader("net", data), nil
}

func NetQuery(query string, out trdsql.Format) error {
	reader, err := NetReader()
	if err != nil {
		return err
	}
	defaultQuery := "SELECT * FROM net"
	if query == "" {
		query = defaultQuery
	}
	return readerQuery(reader, query, out)
}
