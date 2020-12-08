package tcpfunc

import (
	"fmt"
	"log"

	"github.com/mangenotwork/servers-online-manage/lib/structs"
	"github.com/mangenotwork/servers-online-manage/slve/tcpfunc/sys2go"
	"github.com/mangenotwork/servers-online-manage/slve/tcpfunc/windows"
)

//获取系统信息
func SysInfos() (data structs.RetuenSysInfos) {
	data = structs.RetuenSysInfos{}

	//获取当前系统类型
	sysType := sys2go.GetSysType()

	switch sysType {
	case "windows":
		return WinSysInfo()
	default:
		log.Println("不支持的系统类型！")
	}

	return
}

//Windows 基础信息
func WinSysInfo() (data structs.RetuenSysInfos) {
	diskTotal := 0
	diskInfo := windows.WindowsGetDiskInfo()
	for _, v := range diskInfo {
		diskTotal += v.DistUse.Total
	}
	data = structs.RetuenSysInfos{
		SysType:         "windows",
		OsName:          windows.WindowsGetOsInfo(),
		SysArchitecture: sys2go.GetSysArch(),
		CpuCoreNumber:   sys2go.GetCpuCoreNumber(),
		CpuName:         windows.GetCpuName(),
		CpuID:           windows.GetCpuId(),
		BaseBoardID:     windows.GetBaseBoardID(),
		MemTotal:        windows.WindowsGetMemoryTotal(),
		Disk:            windows.WindowsGetDiskInfo(),
		DiskTotal:       fmt.Sprintf("%dMB", diskTotal),
	}
	return
}

//上报性能信息
//cpu 使用率， mem， disk .....
func GetPerformance() {
	if sys2go.GetSysType() != "windows" {
		return
	}

	//获取cpu 使用率, 和每个核心的使用率
	cpuUse := windows.GetCPUUse()
	cpuUseRate := &structs.CPUUseRate{
		CPU:     "cpu",
		UseRate: float32(cpuUse),
	}
	//TODO 每个核心的使用率
	cpucoreUseRate := make([]*structs.CPUUseRate, 0)
	log.Println(cpuUseRate, cpucoreUseRate)

	//获取内存
	memInfo := windows.WindowsGetMemoryInfo()
	log.Println(memInfo)

	//获取磁盘信息
	diskinfo := windows.WindowsGetDiskInfo()
	log.Println(diskinfo)

	//网络IO
	networkIO := windows.GetNetIOFromCMD()
	log.Println(networkIO)

	//有效连接数
	connCount := windows.GetTcpConnCount()
	log.Println(connCount)

	//进程数
	processCount, _ := windows.GetWindowsPIDInfo()
	log.Println(processCount)
}
