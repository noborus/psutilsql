package psutilsql

import (
	"fmt"
	"strconv"
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
	IOCOUNTERS
	FOREGROUND
	BACKGROUND
	ISRUNNING
	COMMAND
)

type processColumn struct {
	names   []string
	types   []string
	getFunc func(p *process.Process) []interface{}
}

var ProcessColumn = map[ColumnNum]processColumn{
	PID: {
		names:   []string{"pid"},
		types:   []string{"int"},
		getFunc: getPid,
	},
	NAME: {
		names:   []string{"name"},
		types:   []string{"text"},
		getFunc: getName,
	},
	CPU: {
		names:   []string{"CPU"},
		types:   []string{"float"},
		getFunc: cpuPercent,
	},
	MEM: {
		names:   []string{"MEM"},
		types:   []string{"float"},
		getFunc: memPercent,
	},
	STATUS: {
		names:   []string{"STATUS"},
		types:   []string{"text"},
		getFunc: status,
	},
	START: {
		names:   []string{"START"},
		types:   []string{"timestamp"},
		getFunc: createTime,
	},
	USER: {
		names:   []string{"USER"},
		types:   []string{"text"},
		getFunc: getUser,
	},
	CWD: {
		names:   []string{"Cwd"},
		types:   []string{"text"},
		getFunc: cwd,
	},
	EXE: {
		names:   []string{"Exe"},
		types:   []string{"text"},
		getFunc: exe,
	},
	TERMINAL: {
		names:   []string{"Terminal"},
		types:   []string{"text"},
		getFunc: terminal,
	},
	IONICE: {
		names:   []string{"IONice"},
		types:   []string{"int"},
		getFunc: ioNice,
	},
	NICE: {
		names:   []string{"Nice"},
		types:   []string{"int"},
		getFunc: nice,
	},
	NUMFDS: {
		names:   []string{"NumFDs"},
		types:   []string{"int"},
		getFunc: numFDs,
	},
	NUMTHREADS: {
		names:   []string{"NumThreads"},
		types:   []string{"int"},
		getFunc: numThreads,
	},
	PPID: {
		names:   []string{"pPid"},
		types:   []string{"int"},
		getFunc: ppid,
	},
	TGID: {
		names:   []string{"Tgid"},
		types:   []string{"int"},
		getFunc: tgid,
	},
	UIDS: {
		names:   []string{"Uids"},
		types:   []string{"text"},
		getFunc: uids,
	},
	GIDS: {
		names:   []string{"Gids"},
		types:   []string{"text"},
		getFunc: gids,
	},
	MEMORYINFOEX: {
		names:   []string{"RSS", "VMS", "Shared", "Text", "Lib", "Data", "Dirty"},
		types:   []string{"int", "int", "int", "int", "int", "int", "int"},
		getFunc: memoryInfoEx,
	},
	MEMORYINFO: {
		names:   []string{"RSS", "VMS", "Data", "Stack", "locked", "Swap"},
		types:   []string{"int", "int", "int", "int", "int", "int"},
		getFunc: memoryInfo,
	},
	IOCOUNTERS: {
		names:   []string{"ReadCount", "WriteCount", "ReadBytes", "WriteBytes"},
		types:   []string{"int", "int", "int", "int"},
		getFunc: ioCounters,
	},
	FOREGROUND: {
		names:   []string{"Foreground"},
		types:   []string{"bool"},
		getFunc: foreGround,
	},
	BACKGROUND: {
		names:   []string{"Background"},
		types:   []string{"bool"},
		getFunc: backGround,
	},
	ISRUNNING: {
		names:   []string{"IsRunning"},
		types:   []string{"bool"},
		getFunc: isRunning,
	},

	COMMAND: {
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
