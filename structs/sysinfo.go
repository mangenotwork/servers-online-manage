package structs

//host信息
type HostInfo struct {
	HostName      string `json:"host_name"`
	SysType       string `json:"sys_type"`
	SysArch       string `json:"sys_architecture"`
	CpuCoreNumber string `json:"cpu_core_number"`
}
