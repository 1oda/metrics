package service

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	net2 "github.com/shirou/gopsutil/net"
	"time"
)

// FetchCPU 获取CPU信息
func FetchCPU() ([]float64, error) {
	return cpu.Percent(time.Second, false)
}

// FetchLoad 获取LOAD信息
func FetchLoad() (*load.AvgStat, error) {
	return load.Avg()
}

// FetchMEM 获取内存信息
func FetchMEM() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

// FetchDisk 获取磁盘信息
func FetchDisk() (*disk.UsageStat, error) {
	return disk.Usage("/")
}

// FetchNet 获取网络信息
func FetchNet(defaultPerNic bool) ([]net2.IOCountersStat, error) {
	return net2.IOCounters(defaultPerNic)
}
