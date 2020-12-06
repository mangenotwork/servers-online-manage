package tcpfunc

import (
	"fmt"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
	"github.com/mangenotwork/servers-online-manage/slve/tcpfunc/sys2go"
	"github.com/mangenotwork/servers-online-manage/slve/tcpfunc/windows"
	"log"
)



//获取系统信息
func SysInfos() (data structs.RetuenSysInfos) {
	data = structs.RetuenSysInfos{}

	//获取当前系统类型
	sysType := sys2go.GetSysType()

	switch sysType {
	case "windows":
		return WinSysInfo()
	case "linux":
		break
	default:
		log.Println("不支持的系统类型！")
	}

	return
}

//Windows 基础信息
func WinSysInfo() (data structs.RetuenSysInfos) {
	diskTotal := 0
	diskInfo := windows.WindowsGetDiskInfo()
	for _,v := range diskInfo{
		diskTotal+=v.DistUse.Total
	}
	data = structs.RetuenSysInfos{
		SysType: "windows",
		OsName: windows.WindowsGetOsInfo(),
		SysArchitecture: sys2go.GetSysArch(),
		CpuCoreNumber: sys2go.GetCpuCoreNumber(),
		CpuName: windows.GetCpuName(),
		CpuID: windows.GetCpuId(),
		BaseBoardID: windows.GetBaseBoardID(),
		MemTotal: windows.WindowsGetMemoryInfo(),
		Disk: windows.WindowsGetDiskInfo(),
		DiskTotal: fmt.Sprintf("%dMB",diskTotal),
	}
	return
}