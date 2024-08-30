package vo

// CPUUsage CPU使用情况
type CPUUsage struct {
	Time  string  `json:"time"`
	Value float64 `json:"value"`
}

// MemoryUsage 内存使用情况
type MemoryUsage struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

// DiskUsage 磁盘使用情况
type DiskUsage struct {
	Name string  `json:"name"`
	Used float64 `json:"used"`
	Free float64 `json:"free"`
}

// NetWorkUsage 网络使用情况
type NetWorkUsage struct {
	Time string `json:"time"`
	In   uint64 `json:"in"`
	Out  uint64 `json:"out"`
}

type SystemVO struct {
	CPUUsage     CPUUsage     `json:"cpu_usage"`
	MemoryUsage  MemoryUsage  `json:"memory_usage"`
	DiskUsage    []DiskUsage  `json:"disk_usage"`
	NetWorkUsage NetWorkUsage `json:"network_usage"`
}
