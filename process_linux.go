//go:build linux
// +build linux

package psutilsql

import (
	"github.com/shirou/gopsutil/process"
)

var processColumn = map[pColumnNum]pColumn{
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

func memoryInfoEx(p *process.Process) []any {
	mx, err := p.MemoryInfoEx()
	if err != nil {
		return []any{0, 0, 0, 0, 0, 0, 0}
	}
	return []any{
		mx.RSS,
		mx.VMS,
		mx.Shared,
		mx.Text,
		mx.Lib,
		mx.Data,
		mx.Dirty,
	}
}
