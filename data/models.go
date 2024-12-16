package data

import (
	"time"

	"github.com/shirou/gopsutil/v4/process"
)

type ProcessItem struct {
	Pid        int32
	Name       string
	Exe        string
	Cmd        string
	CreateTime time.Time
}

func NewProcessItem(p *process.Process) ProcessItem {
	pid := p.Pid
	name, _ := p.Name()
	exe, _ := p.Exe()
	cmd, _ := p.Cmdline()
	createTimeUnix, _ := p.CreateTime()

	createTime := time.UnixMilli(createTimeUnix)

	return ProcessItem{pid, name, exe, cmd, createTime}
}

func (pi *ProcessItem) GetLogItem(event string) LogItem {
	var seconds = 0
	diff := time.Since(pi.CreateTime)
	if diff > time.Second {
		seconds = int(diff.Seconds())
	}

	return LogItem{
		Event:      event,
		Timestamp:  time.Now(),
		Pid:        pi.Pid,
		Name:       pi.Name,
		Exe:        pi.Exe,
		Cmd:        pi.Cmd,
		CreateTime: pi.CreateTime,
		Seconds:    seconds,
	}
}

type LogItem struct {
	Event      string
	Timestamp  time.Time
	Pid        int32
	Name       string
	Exe        string
	Cmd        string
	CreateTime time.Time
	Seconds    int
}
