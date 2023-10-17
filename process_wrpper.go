package psutilsql

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/process"
)

type pColumnNum int

const (
	PID pColumnNum = iota
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
	IOCOUNTERS
	FOREGROUND
	BACKGROUND
	ISRUNNING
	COMMAND
)

type pColumn struct {
	names   []string
	types   []string
	getFunc func(p *process.Process) []any
}

func getPid(p *process.Process) []any {
	return []any{p.Pid}
}

func strWrap(v string, err error) []any {
	if err != nil {
		return []any{""}
	}
	return []any{v}
}

func getName(p *process.Process) []any {
	return strWrap(p.Name())
}

func status(p *process.Process) []any {
	return strWrap(p.Status())
}

func cmdLine(p *process.Process) []any {
	return strWrap(p.Cmdline())
}

func getUser(p *process.Process) []any {
	return strWrap(p.Username())
}

func cwd(p *process.Process) []any {
	return strWrap(p.Cwd())
}

func exe(p *process.Process) []any {
	return strWrap(p.Exe())
}

func terminal(p *process.Process) []any {
	return strWrap(p.Terminal())
}

func numWrap(v any, err error) []any {
	if err != nil {
		return []any{0}
	}
	return []any{v}
}

func cpuPercent(p *process.Process) []any {
	return numWrap(p.CPUPercent())
}

func memPercent(p *process.Process) []any {
	return numWrap(p.MemoryPercent())
}

func ioNice(p *process.Process) []any {
	return numWrap(p.IOnice())
}

func nice(p *process.Process) []any {
	return numWrap(p.Nice())
}

func numFDs(p *process.Process) []any {
	return numWrap(p.NumFDs())
}

func numThreads(p *process.Process) []any {
	return numWrap(p.NumThreads())
}

func ppid(p *process.Process) []any {
	return numWrap(p.Ppid())
}

func tgid(p *process.Process) []any {
	return numWrap(p.Tgid())
}

func sliceWrap(v []int32, err error) []any {
	if err != nil {
		return []any{0}
	}
	s := make([]string, len(v))
	for i, vv := range v {
		s[i] = fmt.Sprint(vv)
	}
	return []any{strings.Join(s, ",")}
}

func uids(p *process.Process) []any {
	return sliceWrap(p.Uids())
}

func gids(p *process.Process) []any {
	return sliceWrap(p.Gids())
}

func boolWrap(v bool, err error) []any {
	if err != nil {
		return []any{""}
	}
	return []any{strconv.FormatBool(v)}
}

func foreGround(p *process.Process) []any {
	return boolWrap(p.Foreground())
}

func backGround(p *process.Process) []any {
	return boolWrap(p.Background())
}

func isRunning(p *process.Process) []any {
	return boolWrap(p.IsRunning())
}

func createTime(p *process.Process) []any {
	c, err := p.CreateTime()
	if err != nil {
		return []any{0}
	}
	return []any{time.Unix(c/1000, 0)}
}

func memoryInfo(p *process.Process) []any {
	mx, err := p.MemoryInfo()
	if err != nil {
		return []any{0, 0, 0, 0, 0, 0}
	}
	return []any{
		mx.RSS,
		mx.VMS,
		mx.Data,
		mx.Stack,
		mx.Locked,
		mx.Swap,
	}
}

func ioCounters(p *process.Process) []any {
	io, err := p.IOCounters()
	if err != nil {
		return []any{0, 0, 0, 0}
	}
	return []any{
		io.ReadCount,
		io.WriteCount,
		io.ReadBytes,
		io.WriteBytes,
	}
}

/*
func pageFaults(p *process.Process) []any {
	pf, err := p.PageFaults
	if err != nil {
		return []any{0, 0, 0, 0}
	}
	return []any{
		pf.MinorFaults,
		pf.MajorFaults,
		pf.ChildMinorFaults,
		pf.ChildMajorFaults,
	}
}
*/
