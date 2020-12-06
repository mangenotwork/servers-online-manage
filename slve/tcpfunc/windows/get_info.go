package windows

/*
#include "getwininfo.h"
*/
import "C"

import (
	"fmt"
	"log"
	"github.com/mangenotwork/servers-online-manage/lib/cmd"
)

//获取屏幕尺寸
func RunMetrics() string{
	return  C.GoString(C.GET_CXSCREEN())
}

//获取计算机名称
func GetComputerName() string {
	return C.GoString(C.WindowsGetComputerName())
}

//获取计算机环境变量
/*
%WINDIR%                 {系统目录- C:\WINDOWS}
%SYSTEMROOT%             {系统目录- C:\WINDOWS}
%SYSTEMDRIVE%            {系统根目录- C:}
%HOMEDRIVE%              {当前用户根目录- C:}
%USERPROFILE%            {当前用户目录- C:\Documents and Settings\wy}
%HOMEPATH%               {当前用户路径- \Documents and Settings\wy}
%TMP%                    {当前用户临时文件夹- C:\DOCUME~1\wy\LOCALS~1\Temp}
%TEMP%                   {当前用户临时文件夹- C:\DOCUME~1\wy\LOCALS~1\Temp}
%APPDATA%                {当前用户数据文件夹- C:\Documents and Settings\wy\Application Data}
%PROGRAMFILES%           {程序默认安装目录- C:\Program Files}
%COMMONPROGRAMFILES%     {文件通用目录- C:\Program Files\Common Files}
%USERNAME%               {当前用户名- wy}
%ALLUSERSPROFILE%        {所有用户文件目录- C:\Documents and Settings\All Users}
%OS%                     {操作系统名- Windows_NT}
%COMPUTERNAME%           {计算机名- IBM-B63851E95C9}
%NUMBER_OF_PROCESSORS%   {处理器个数- 1}
%PROCESSOR_ARCHITECTURE% {处理器芯片架构 - x86}
%PROCESSOR_LEVEL%        {处理器型号- 6}
%PROCESSOR_REVISION%     {处理器修订号- 0905}
%USERDOMAIN%             {包含用户帐号的域- IBM-B63851E95C9}
%COMSPEC%                {C:\WINDOWS\system32\cmd.exe}

%PATHEXT% {执行文件类型 -.COM;.EXE;.BAT;.CMD;.VBS;.VBE;.JS;.JSE;.WSF;.WSH;.pyo;.pyc;.py;.pyw}
%PATH%    {搜索路径}
 */
func GetEnvironment(input string) string{
	cs := C.CString(input)
	return C.GoString(C.GetEnvironment(cs))
}

//取得与底层硬件平台有关的信息
func GetSystemInfo(){
	a := C.WindowsGetSystemInfo()
	log.Println(a)
	// http://www.office-cn.net/t/api/system_info.htm
	log.Println(a.dwPageSize)
	//fmt.Println(a.dwOemID)
	log.Println(a.lpMinimumApplicationAddress)
	log.Println(a.lpMaximumApplicationAddress)
	log.Println(a.dwActiveProcessorMask)
	//fmt.Println(a.dwNumberOrfProcessors)
	log.Println(a.dwProcessorType)
	log.Println(a.dwAllocationGranularity)
	log.Println(a.wProcessorLevel)
	log.Println(a.wProcessorRevision)
}

func GETSystemPowerStatus(){
	a := C.GETSystemPowerStatus()

	fmt.Println(a)

	//ACLineStatus   交流电源状态
	//0  Offline
	//1		Online
	//255    Unknown status
	fmt.Println(a.ACLineStatus)

	//BatteryFlag   电池充电状态。 可以包含一或多个以下值
	// 1	高，电量大于66%
	// 2	低，小于33%
	// 4	极低，小于5%
	// 8	充电中
	// 128	没有电池
	// 255	未知，无法读取状态
	fmt.Println(a.BatteryFlag)

	//Reserved1   保留，必须为0
	fmt.Println(a.Reserved1)

	//BatteryLifeTime    秒为单位的电池剩余电量, 若未知则为-1
	fmt.Println(a.BatteryLifeTime)

	//BatteryFullLifeTime   秒为单位的电池充满电的电量，若未知则为-1
	fmt.Println(a.BatteryFullLifeTime)
}

//用于获取自windows启动以来经历的时间长度（毫秒）
func GET_TickCount(){
	a := C.GET_TickCount()
	log.Println(a)
}

//为当前用户取得默认语言ID
func WindowsGetUserDefaultLangID(){
	a := C.WindowsGetUserDefaultLangID()
	log.Println(a)
}

//获取 windows系统型号
func WindowsGetOsInfo(){
	osinfo := C.GetOsInfo()
	buildNumber := osinfo.dwBuildNumber
	osName := "Windows ?"
	versionNumber := fmt.Sprintf("%d.%d",osinfo.dwMajorVersion,osinfo.dwMinorVersion)
	switch versionNumber {
	case "6.2":
		osName = "Windows 10/Servers-2012"
	case "6.1":
		osName = "Windows 7/Servers-2008-R2"
	case "6.0":
		osName = "Windows Vista/Servers-2008"
	case "5.2":
		osName = "Windows XP-x64/Servers-2003"
	case "5.1":
		osName = "Windows XP"
	case "5.0":
		osName = "Windows 2000"
	}
	fmt.Println("Name  = ", osName)
	fmt.Println("versionNumber  = ", versionNumber)
	fmt.Println("BuildNumber  = ", buildNumber)
}

//获取 内存大小
func WindowsGetMemoryInfo(){
	mem := C.GetMemoryInfo()
	fmt.Println("内存占用率 = ", mem.dwMemoryLoad)
	fmt.Println("总物理内存 = ", mem.ullTotalPhys/1024/1024," MB" )
	fmt.Println("闲置物理内存 = ", mem.ullAvailPhys/1024/1024, " MB")
}

//获取主板ID
func GetBaseBoardID(){
	cmds := []string{"wmic", "baseboard", "get", "serialnumber"}
	boardId := cmd.WindowsSendCommand(cmds)
	log.Println(boardId)
}
