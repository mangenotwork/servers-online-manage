package tcpfunc

import (
	"fmt"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
	"github.com/mangenotwork/servers-online-manage/slve/tcpfunc/linux"
	"github.com/mangenotwork/servers-online-manage/slve/tcpfunc/sys2go"
	"log"
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