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
	getFunc func(p *process.Process) []interface{}
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

func boolWrap(v bool, err error) []interface{} {
	if err != nil {
		return []interface{}{""}
	}
	return []interface{}{strconv.FormatBool(v)}
}

func foreGround(p *process.Process) []interface{} {
	return boolWrap(p.Foreground())
}

func backGround(p *process.Process) []interface{} {
	return boolWrap(p.Background())
}

func isRunning(p *process.Process) []interface{} {
	return boolWrap(p.IsRunning())
}

func createTime(p *process.Process) []interface{} {
	c, err := p.CreateTime()
	if err != nil {
		return []interface{}{0}
	}
	return []interface{}{time.Unix(c/1000, 0)}
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

func ioCounters(p *process.Process) []interface{} {
	io, err := p.IOCounters()
	if err != nil {
		return []interface{}{0, 0, 0, 0}
	}
	return []interface{}{
		io.ReadCount,
		io.WriteCount,
		io.ReadBytes,
		io.WriteBytes,
	}
}

/*
func pageFaults(p *process.Process) []interface{} {
	pf, err := p.PageFaults
	if err != nil {
		return []interface{}{0, 0, 0, 0}
	}
	return []interface{}{
		pf.MinorFaults,
		pf.MajorFaults,
		pf.ChildMinorFaults,
		pf.ChildMajorFaults,
	}
}
*/
