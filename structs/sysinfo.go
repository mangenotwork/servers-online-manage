package structs

//host信息
type HostInfo struct {
	HostName string `json:"host_name"`
	SysType  string `json:"sys_type"`

	//系统架构
	SysArch string `json:"sys_architecture"`

	//cpu核心数
	CpuCoreNumber string `json:"cpu_core_number"`
}

//从 proc/meminfo 获取内存信息
type ProcMemInfo struct {
	//所有可用RAM大小 （即物理内存减去一些预留位和内核的二进制代码大小）
	MemTotal int64 `json:"mem_total"`

	//内存使用
	MemUsed int64 `json:"mem_used"`

	//LowFree与HighFree的总和，被系统留着未使用的内存
	MemFree int64 `json:"mem_free"`

	//用来给文件做缓冲大小
	MemBuffers int64 `json:"mem_buffers"`

	//被高速缓冲存储器（cache memory）用的内存的大小（等于diskcache minus SwapCache ）.
	MemCached int64 `json:"mem_cached"`
}

// 从/proc/stat 获取cpu信息计算后的输出结果
type ProcStatCPUData struct {
	Name string

	//总的cpu时间totalCpuTime
	Total int64
	Used  int64

	//从系统启动开始累计到当前时刻，除IO等待时间以外的其它等待时间。
	Idle int64
}

// 从/proc/<pid>/stat 获取某个进程cpu信息计算后的输出结果
type ProcessProcStatCPUData struct {
	Name           string
	ProcessCpuTime int64
	TaskState      string
	Ppid           string
	Pgid           string
	Sid            string
	NumThreads     string
}

//从 /proc/diskstats  -- 每块磁盘设备的磁盘I/O统计信息列表
type ProcDiskstatsData struct {
	//设备
	DiskName string
	//输入/输出操作花费的毫秒数
	IOTime int64
	//读完成次数
	ReadRequest int64
	//写完成次数
	WriteRequest int64
	//读操作花费毫秒数
	MsecRead int64
	//写操作花费的毫秒数
	MsecWrite int64
}

//从/proc/net/dev中读取  采集网卡信息
type ProcNetDevData struct {
	Name string
	Recv int64
	Send int64
}
