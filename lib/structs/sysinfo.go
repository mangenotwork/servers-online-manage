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

//磁盘信息
type DiskInfo struct {
	DiskName    string
	DistType    string
	DistTotalMB string
	DistUse     *DiskUseInfo
}

//磁盘使用的信息
type DiskUseInfo struct {
	Total int     //MB
	Free  int     //MB
	Rate  float32 //%
}

//返回的系统类型结构
type RetuenSysInfos struct {
	//系统平台
	SysType string `json:"sys_type"`
	//系统版本 os_name+版号
	OsName string `json:"os_name"`
	//系统架构
	SysArchitecture string `json:"sys_architecture"`
	//CPU核心数
	CpuCoreNumber string `json:"cpu_core_number"`
	//CPU name
	CpuName string `json:"cpu_name"`
	//CPU ID
	CpuID string `json:"cpu_id"`
	//主板ID
	BaseBoardID string `json:"board_id"`
	//内存总大小 MB
	MemTotal string `json:"mem_totle"`
	//磁盘信息
	Disk []*DiskInfo `json:"disk"`
	//磁盘总大小 MB
	DiskTotal string `json:"disk_totle"`
}

//环境变量
//Slve Env Info
type EnvInfos struct {
}

//cpu使用率
type CPUUseRate struct {
	CPU     string
	UseRate float32
}

//网络IO - 简单
//单位 (kb/sec)
type NetWorkIOSimple struct {
	Name string
	Tx   float32 //发送
	Rx   float32 //接收
}

//性能采集信息
type SlvePerformanceData struct {
	//获取cpu 使用率, 和每个核心的使用率
	CpuRate     *CPUUseRate
	CpucoreRate []*CPUUseRate

	//获取内存
	MemInfo *ProcMemInfo

	//获取磁盘信息
	DiskInfo []*DiskInfo

	//磁盘IO

	//网络IO
	NetworkIO []*NetWorkIOSimple

	//连接数
	TcpConnCount int

	//进程数
	PIDCount int
}
