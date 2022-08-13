package controller

import (
	"github.com/prometheus/client_golang/prometheus"
	"metrics/model/service"
	"sync"
)

// Metrics metrics指标
type Metrics struct {
	metrics map[string]*prometheus.Desc
	mutex   sync.Mutex
}

type SystemInfo struct {
	CPU, LOAD, MEM, DISK, NET map[string]float64
}

const (
	CPU_METRICS  = "node_cpu_metrics"
	MEM_METRICS  = "node_mem_metrics"
	LOAD_METRICS = "node_load_metrics"
	DISK_METRICS = "node_disk_metrics"
	NET_METRICS  = "node_net_metrics"

	DefaultPerNic    bool   = true
	DefaultMetricsNS string = "prom"
)

func Registry() prometheus.Gatherer {
	registry := prometheus.NewRegistry()
	registry.MustRegister(NewMetrics(DefaultMetricsNS))
	return registry
}

// 初始化metrics
func newMetric(namespace string, metricName string, docString string, labels []string) *prometheus.Desc {
	return prometheus.NewDesc(namespace+"_"+metricName, docString, labels, nil)
}

// NewMetrics 初始化metrics
func NewMetrics(namespace string) *Metrics {
	return &Metrics{
		metrics: map[string]*prometheus.Desc{
			NET_METRICS:  newMetric(namespace, "server_network", "The description of gauge_metric", []string{"name"}),
			CPU_METRICS:  newMetric(namespace, "server_cpu", "The description of cpu_metric", []string{"name"}),
			LOAD_METRICS: newMetric(namespace, "server_load", "The description of load_metric", []string{"name"}),
			DISK_METRICS: newMetric(namespace, "server_disk", "The description of cpu_metric", []string{"name"}),
			MEM_METRICS:  newMetric(namespace, "server_mem", "The description of cpu_metric", []string{"name"}),
		},
	}
}

// Describe 注册metric结构体，然后将metric指标信息放入chan队列中
func (c *Metrics) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range c.metrics {
		ch <- m
	}
}

// Collect 采集metric指标数据，然后将监控指标数据存放到对应的chan队列中
func (c *Metrics) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock() // 加锁
	defer c.mutex.Unlock()

	allMetrics, err := fetchAllMetrics()
	if err != nil {
		panic(err)
	}

	for k, v := range allMetrics.CPU {
		ch <- prometheus.MustNewConstMetric(c.metrics[CPU_METRICS], prometheus.CounterValue, v, k)
	}
	for k, v := range allMetrics.LOAD {
		ch <- prometheus.MustNewConstMetric(c.metrics[LOAD_METRICS], prometheus.CounterValue, v, k)
	}
	for k, v := range allMetrics.MEM {
		ch <- prometheus.MustNewConstMetric(c.metrics[MEM_METRICS], prometheus.CounterValue, v, k)
	}
	for k, v := range allMetrics.DISK {
		ch <- prometheus.MustNewConstMetric(c.metrics[DISK_METRICS], prometheus.CounterValue, v, k)
	}
	for k, v := range allMetrics.NET {
		ch <- prometheus.MustNewConstMetric(c.metrics[NET_METRICS], prometheus.CounterValue, v, k)
	}
}

// 获取所有信息
func fetchAllMetrics() (SystemInfo, error) {
	cpuInfo, err := service.FetchCPU()
	if err != nil {
		return SystemInfo{}, err
	}

	loadInfo, err := service.FetchLoad()
	if err != nil {
		return SystemInfo{}, err
	}

	memInfo, err := service.FetchMEM()
	if err != nil {
		return SystemInfo{}, err
	}

	diskInfo, err := service.FetchDisk()
	if err != nil {
		return SystemInfo{}, err
	}

	netInfo, err := service.FetchNet(DefaultPerNic)
	if err != nil {
		return SystemInfo{}, err
	}

	return SystemInfo{
		CPU: map[string]float64{
			"usage": cpuInfo[0],
		},
		MEM: map[string]float64{
			"total": float64(memInfo.Total / 1024 / 1024),
			"free":  float64(memInfo.Free / 1024 / 1024),
			"used":  float64(memInfo.Used / 1024 / 1024),
			"usage": memInfo.UsedPercent,
		},
		LOAD: map[string]float64{
			"load1":  loadInfo.Load1,
			"load5":  loadInfo.Load5,
			"load15": loadInfo.Load15,
		},
		DISK: map[string]float64{
			"total": float64(diskInfo.Total / 1024 / 1024),
			"free":  float64(diskInfo.Free / 1024 / 1024),
			"usage": diskInfo.UsedPercent,
		},
		NET: map[string]float64{
			"BytesRecv": float64(netInfo[0].BytesRecv),
			"bytesSent": float64(netInfo[0].BytesSent),
		},
	}, nil
}
