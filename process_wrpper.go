package psutilsql

import (
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/process"
)

type ColumnNum int

const (
	PID ColumnNum = iota
	NAME
	CPU
	MEM
	STATUS
	START
	USER
	CWD
	EXE
	TERMINAL
	IONICE
	NICE
	NUMFDS
	NUMTHREADS
	PPID
	TGID
	UIDS
	GIDS
	MEMORYINFOEX
	MEMORYINFO
	COMMAND
)

type processColumn struct {
	column  ColumnNum
	enable  bool
	names   []string
	types   []string
	getFunc func(p *process.Process) []interface{}
}

var ProcessColumn = []processColumn{
	{
		column:  PID,
		enable:  true,
		names:   []string{"pid"},
		types:   []string{"int"},
		getFunc: getPid,
	},
	{
		column:  NAME,
		enable:  true,
		names:   []string{"name"},
		types:   []string{"text"},
		getFunc: getName,
	},
	{
		column:  CPU,
		enable:  true,
		names:   []string{"CPU"},
		types:   []string{"float"},
		getFunc: cpuPercent,
	},
	{
		column:  MEM,
		enable:  true,
		names:   []string{"MEM"},
		types:   []string{"float"},
		getFunc: memPercent,
	},
	{
		column:  STATUS,
		enable:  true,
		names:   []string{"STATUS"},
		types:   []string{"text"},
		getFunc: status,
	},
	{
		column:  START,
		enable:  true,
		names:   []string{"START"},
		types:   []string{"timestamp"},
		getFunc: createTime,
	},
	{
		column:  USER,
		enable:  true,
		names:   []string{"USER"},
		types:   []string{"text"},
		getFunc: getUser,
	},
	{
		column:  CWD,
		enable:  false,
		names:   []string{"Cwd"},
		types:   []string{"text"},
		getFunc: cwd,
	},
	{
		column:  EXE,
		enable:  false,
		names:   []string{"Exe"},
		types:   []string{"text"},
		getFunc: exe,
	},
	{
		column:  TERMINAL,
		enable:  false,
		names:   []string{"Terminal"},
		types:   []string{"text"},
		getFunc: terminal,
	},
	{
		column:  IONICE,
		enable:  false,
		names:   []string{"IONice"},
		types:   []string{"int"},
		getFunc: ioNice,
	},
	{
		column:  NICE,
		enable:  false,
		names:   []string{"Nice"},
		types:   []string{"int"},
		getFunc: nice,
	},
	{
		column:  NUMFDS,
		enable:  false,
		names:   []string{"NumFDs"},
		types:   []string{"int"},
		getFunc: numFDs,
	},
	{
		column:  NUMTHREADS,
		enable:  false,
		names:   []string{"NumThreads"},
		types:   []string{"int"},
		getFunc: numThreads,
	},
	{
		column:  PPID,
		enable:  false,
		names:   []string{"pPid"},
		types:   []string{"int"},
		getFunc: ppid,
	},
	{
		column:  TGID,
		enable:  false,
		names:   []string{"Tgid"},
		types:   []string{"int"},
		getFunc: tgid,
	},
	{
		column:  UIDS,
		enable:  false,
		names:   []string{"Uids"},
		types:   []string{"text"},
		getFunc: uids,
	},
	{
		column:  GIDS,
		enable:  false,
		names:   []string{"Gids"},
		types:   []string{"text"},
		getFunc: gids,
	},
	{
		column:  MEMORYINFOEX,
		enable:  false,
		names:   []string{"RSS", "VMS", "Shared", "Text", "Lib", "Data", "Dirty"},
		types:   []string{"int", "int", "int", "int", "int", "int", "int"},
		getFunc: memoryInfoEx,
	},
	{
		column:  MEMORYINFO,
		enable:  true,
		names:   []string{"RSS", "VMS", "Data", "Stack", "locked", "Swap"},
		types:   []string{"int", "int", "int", "int", "int", "int"},
		getFunc: memoryInfo,
	},
	{
		column:  COMMAND,
		enable:  true,
		names:   []string{"COMMAND"},
		types:   []string{"text"},
		getFunc: cmdLine,
	},
}

func getPid(p *process.Process) []interface{} {
	return []interface{}{p.Pid}
}
func strWrap(v string, err error) []interface{} {
	if err != nil {
		return []interface{}{""}
	}
	return []interface{}{v}
}
func getName(p *process.Process) []interface{} {
	return strWrap(p.Name())
}
func status(p *process.Process) []interface{} {
	return strWrap(p.Status())
}
func cmdLine(p *process.Process) []interface{} {
	return strWrap(p.Cmdline())
}
func getUser(p *process.Process) []interface{} {
	return strWrap(p.Username())
}
func cwd(p *process.Process) []interface{} {
	return strWrap(p.Cwd())
}
func exe(p *process.Process) []interface{} {
	return strWrap(p.Exe())
}
func terminal(p *process.Process) []interface{} {
	return strWrap(p.Terminal())
}

func numWrap(v interface{}, err error) []interface{} {
	if err != nil {
		return []interface{}{0}
	}
	return []interface{}{v}
}
func cpuPercent(p *process.Process) []interface{} {
	return numWrap(p.CPUPercent())
}
func memPercent(p *process.Process) []interface{} {
	return numWrap(p.MemoryPercent())
}
func ioNice(p *process.Process) []interface{} {
	return numWrap(p.IOnice())
}
func nice(p *process.Process) []interface{} {
	return numWrap(p.Nice())
}
func numFDs(p *process.Process) []interface{} {
	return numWrap(p.NumFDs())
}
func numThreads(p *process.Process) []interface{} {
	return numWrap(p.NumThreads())
}
func ppid(p *process.Process) []interface{} {
	return numWrap(p.Ppid())
}
func tgid(p *process.Process) []interface{} {
	return numWrap(p.Tgid())
}
func sliceWrap(v []int32, err error) []interface{} {
	if err != nil {
		return []interface{}{0}
	}
	s := make([]string, len(v))
	for i, vv := range v {
		s[i] = fmt.Sprint(vv)
	}
	return []interface{}{strings.Join(s, ",")}
}
func uids(p *process.Process) []interface{} {
	return sliceWrap(p.Uids())
}
func gids(p *process.Process) []interface{} {
	return sliceWrap(p.Gids())
}
func createTime(p *process.Process) []interface{} {
	c, err := p.CreateTime()
	if err != nil {
		return []interface{}{0}
	}
	return []interface{}{time.Unix(c/1000, 0)}
}

func memoryInfoEx(p *process.Process) []interface{} {
	mx, err := p.MemoryInfoEx()
	if err != nil {
		return []interface{}{0, 0, 0, 0, 0, 0, 0}
	}
	return []interface{}{
		mx.RSS,
		mx.VMS,
		mx.Shared,
		mx.Text,
		mx.Lib,
		mx.Data,
		mx.Dirty,
	}
}
func memoryInfo(p *process.Process) []interface{} {
	mx, err := p.MemoryInfo()
	if err != nil {
		return []interface{}{0, 0, 0, 0, 0, 0}
	}
	return []interface{}{
		mx.RSS,
		mx.VMS,
		mx.Data,
		mx.Stack,
		mx.Locked,
		mx.Swap,
	}
}
