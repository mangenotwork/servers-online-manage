package tcpfunc

import (
	"fmt"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
	"github.com/mangenotwork/servers-online-manage/slve/tcpfunc/linux"
	"github.com/mangenotwork/servers-online-manage/slve/tcpfunc/sys2go"
	"log"
	"time"
)



//获取系统信息
func SysInfos() (data structs.RetuenSysInfos) {
	data = structs.RetuenSysInfos{}
	//获取当前系统类型
	sysType := sys2go.GetSysType()
	switch sysType {
	case "linux":
		return LinuxSysInfo()
	default:
		log.Println("不支持的系统类型！")
	}
	return
}


//linux 基础信息
func LinuxSysInfo() (data structs.RetuenSysInfos) {
	osName,_ := linux.ProcVersion()
	cpuName := ""
	cpuInfoList := linux.ProcCpuinfo()
	if len(cpuInfoList) >= 2 {
		cpuName = cpuInfoList[1]["model name\t"]
	}
	memInfo := linux.ProcMeminfo()
	diskinfo,diskTotal := linux.GetSystemDF()
	data = structs.RetuenSysInfos{
		SysType: "linux",
		OsName: osName,
		SysArchitecture: sys2go.GetSysArch(),
		CpuCoreNumber: sys2go.GetCpuCoreNumber(),
		CpuName: cpuName,
		CpuID: linux.GetCPUIDFromLinux(),
		MemTotal: fmt.Sprintf("%d MB",memInfo.MemTotal/1024),
		Disk: diskinfo,
		DiskTotal: fmt.Sprintf("%dMB",diskTotal),
	}
	return
}

//上报性能信息
//cpu 使用率， mem， disk .....
func GetPerformance() *structs.SlvePerformanceData {
	if sys2go.GetSysType() != "linux" {
			return &structs.SlvePerformanceData{}
	}

	t := 500 * time.Millisecond
	//获取cpu 使用率, 和每个核心的使用率
	cpuUseRate, cpucoreUseRate := linux.ProcStat(t)
	log.Println(cpuUseRate,cpucoreUseRate)

	//获取内存
	memInfo := linux.ProcMeminfo()
	log.Println(memInfo)

	//获取磁盘信息
	diskinfo,_ := linux.GetSystemDF()
	log.Println(diskinfo)

	//网络IO
	networkIO := linux.ProcNetDev(t)
	log.Println(networkIO)

	//有效连接数
	connCount := linux.GetTcpConnCount()
	log.Println(connCount)

	//进程数
	processCount := linux.GetProcessCount()
	log.Println(processCount)

	return &structs.SlvePerformanceData{
		CpuRate:      cpuUseRate,
		CpucoreRate:  cpucoreUseRate,
		MemInfo:      memInfo,
		DiskInfo:     diskinfo,
		NetworkIO:    networkIO,
		TcpConnCount: connCount,
		PIDCount:     processCount,
	}
}

//获取slve详情信息
func GetDetailsInfo(){
	//基础信息
	//环境变量
	//进程列表

}