package service

import (
	"errors"
	"github.com/shirou/gopsutil/process"
	"metrics/model/core"
	"strings"
	"time"
)

type MonitorProcess struct {
	ID         int
	ProcessID  string
	CPUUsage   string
	MEMUsage   string
	UpdateTime int
}

// FetchAllRunningProcess 获取当前运行的所有进程
func FetchAllRunningProcess() ([]core.Process, error) {
	processes, err := process.Processes()
	if err != nil {
		return []core.Process{}, err
	}

	var processList []core.Process
	for _, p := range processes {

		user, err := p.Username()
		if err != nil {
			if err.Error() == "user: unknown userid 1000" {
				continue
			}
			return []core.Process{}, err
		}
		if user == "" {
			continue
		}

		cmdLine, err := p.Cmdline()
		if err != nil {
			return []core.Process{}, err
		}

		if cmdLine == "" {
			continue
		}

		if len(cmdLine) < 1 || cmdLine == "" {
			continue
		}

		cmdLineArr := strings.Split(cmdLine, "\"")
		cmd := cmdLineArr[0]
		if len(cmdLineArr) > 1 && cmdLineArr[0] == "" {
			cmd = cmdLineArr[1]
		}

		processList = append(processList, core.Process{
			PID:        int(p.Pid),
			User:       user,
			Cmd:        cmd,
			CmdLine:    cmdLine,
			UpdateTime: int(time.Now().Unix()),
		})
	}
	return processList, nil
}

// AutoCreateProcess 自动添加当前运行的所有进程
func AutoCreateProcess() error {
	processes, err := process.Processes()
	if err != nil {
		return err
	}

	var processList []core.Process
	var processItem core.Process
	for _, p := range processes {

		user, err := p.Username()
		if err != nil {
			if err.Error() == "user: unknown userid 1000" {
				continue
			}
			return err
		}
		if user == "" {
			continue
		}

		cmdLine, err := p.Cmdline()
		if err != nil {
			return err
		}

		if cmdLine == "" {
			continue
		}

		cmdLineArr := strings.Split(cmdLine, "\"")
		cmd := cmdLineArr[0]
		if len(cmdLineArr) > 1 && cmdLineArr[0] == "" {
			cmd = cmdLineArr[1]
		}

		processList = append(processList, core.Process{
			PID:        int(p.Pid),
			User:       user,
			Cmd:        cmd,
			CmdLine:    cmdLine,
			UpdateTime: int(time.Now().Unix()),
		})
	}

	return processItem.ProcessInsertMulti(processList)
}

// FetchAllMonitorProcess 获取监控的进程信息
func FetchAllMonitorProcess() ([]core.Process, error) {
	var err error
	var processes []core.Process
	var p core.Process = core.Process{}

	_, err = p.ProcessQuery().All(&processes, "id", "pid", "user", "cmd", "cmd_line", "update_time")
	return processes, err

}

// DeleteMonitorProcessByID 根据ID删除一个监控进程
func DeleteMonitorProcessByID(id int) error {
	var p core.Process = core.Process{}
	p.ID = id
	return p.ProcessDelete()
}

// UpdateMonitorProcess 更新一个监控进程
func UpdateMonitorProcess(p core.Process) error {
	return p.ProcessUpdate("pid", "user", "cmd", "cmd_line", "update_time")
}

// CreateMonitorProcess 新增一个监控进程
func CreateMonitorProcess(p core.Process) error {
	res := p.ProcessInsert()
	if res == 0 {
		return errors.New("insert data failed")
	}
	return nil
}
