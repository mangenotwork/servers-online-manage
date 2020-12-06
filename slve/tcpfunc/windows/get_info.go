package windows

/*
#include <windows.h>
#include <stdio.h>
#include <conio.h>

double FileTimeToDouble(FILETIME* pFiletime)
{
	return (double)((*pFiletime).dwHighDateTime * 4.294967296E9) + (double)(*pFiletime).dwLowDateTime;
}

double m_fOldCPUIdleTime;
double m_fOldCPUKernelTime;
double m_fOldCPUUserTime;


char* ComputerNameStr;//计算机名称
char* EnvironmentStr;//计算机环境变量

BOOL Initialize()
{
	FILETIME ftIdle, ftKernel, ftUser;
	BOOL flag = FALSE;
	if (flag = GetSystemTimes(&ftIdle, &ftKernel, &ftUser))
	{
		m_fOldCPUIdleTime = FileTimeToDouble(&ftIdle);
		m_fOldCPUKernelTime = FileTimeToDouble(&ftKernel);
		m_fOldCPUUserTime = FileTimeToDouble(&ftUser);

	}
	return flag;
}

//获取cpu使用
int GetCPUUseRate()
{
	int nCPUUseRate = -1;
	FILETIME ftIdle, ftKernel, ftUser;
	if (GetSystemTimes(&ftIdle, &ftKernel, &ftUser))
	{
		double fCPUIdleTime = FileTimeToDouble(&ftIdle);
		double fCPUKernelTime = FileTimeToDouble(&ftKernel);
		double fCPUUserTime = FileTimeToDouble(&ftUser);
		nCPUUseRate= (int)(100.0 - (fCPUIdleTime - m_fOldCPUIdleTime) / (fCPUKernelTime - m_fOldCPUKernelTime + fCPUUserTime - m_fOldCPUUserTime)*100.0);
		m_fOldCPUIdleTime = fCPUIdleTime;
		m_fOldCPUKernelTime = fCPUKernelTime;
		m_fOldCPUUserTime = fCPUUserTime;
	}
	return nCPUUseRate;
}

//获取cpu使用
int cpu()
{
	if (!Initialize())
	{
		getch();
		return -1;
	}
	else
	{
		Sleep(1000);
		return GetCPUUseRate();
	}
	return -1;
}

//获取计算机名称
char* WindowsGetComputerName()
{
	unsigned long size=255;
	ComputerNameStr = (char *)malloc(size);
 	GetComputerName(ComputerNameStr,&size);
    return ComputerNameStr;
}

//获取环境变量
char* GetEnvironment(char *input)
{
	unsigned long size=1024;
	EnvironmentStr = (char *)malloc(size);
 	GetEnvironmentVariableA(input,EnvironmentStr,size);
    return EnvironmentStr;
}


//取得与底层硬件平台有关的信息
SYSTEM_INFO WindowsGetSystemInfo()
{
	SYSTEM_INFO SystemInfo;
	GetSystemInfo(&SystemInfo);
	return SystemInfo;
}

//获得与当前系统电源状态有关的信息。对便携式计算机来说，这些信息特别有用。
//在那些地方，可用这个函数了解有关电源和电池组的情况
SYSTEM_POWER_STATUS GETSystemPowerStatus()
{
	SYSTEM_POWER_STATUS lpSystemPowerStatus;
	int a;
	a = GetSystemPowerStatus(&lpSystemPowerStatus);
	printf("%d",a);
	return lpSystemPowerStatus;
}

//用于获取自windows启动以来经历的时间长度（毫秒）
long GET_TickCount(){
	long a;
	a = GetTickCount();
	return a;
}

//在一个TIME_ZONE_INFORMATION结构中载入与系统时区设置有关的信息
//http://www.office-cn.net/t/api/gettimezoneinformation.htm

//为当前用户取得默认语言ID
long WindowsGetUserDefaultLangID(){
	long a;
	a = GetUserDefaultLangID();
	return a;
}

//取得当前用户的默认“地方”设置
//http://www.office-cn.net/t/api/getuserdefaultlcid.htm

//取得当前用户的名字
//http://www.office-cn.net/t/api/getusername.htm

//判断当前运行的Windows和DOS版本
//http://www.office-cn.net/t/api/getversion.htm

//设置新的计算机名
long WindowsSetComputerName(char *input){
	long a;
	a = SetComputerNameA(input);
	return a;
}

//设置当前系统时间
long WindowsSetSystemTime(){
	int val=0;
    SYSTEMTIME system_time = {0};

	//先获取本地时间
    GetLocalTime(&system_time);

	//只修改年和月份
    system_time.wYear = 1988;
    system_time.wMonth = 8;

    val = SetLocalTime(&system_time);

	if(val == 0){
        printf("设置本地时间失败！\n");
    }
    else{
		printf("设置本地时间成功！\n");
	}

	return val;
}

//获取windows系统版本号相关信息
OSVERSIONINFO GetOsInfo()
{
    OSVERSIONINFO osver = {sizeof(OSVERSIONINFO)};
    GetVersionEx(&osver);
    return osver;
}

//获取内存信息
MEMORYSTATUSEX GetMemoryInfo()
{
    MEMORYSTATUSEX statusex;
    statusex.dwLength = sizeof(statusex);
    GlobalMemoryStatusEx(&statusex);
    return statusex;
}

*/
import "C"
import (
	"fmt"
	"log"
	"github.com/mangenotwork/servers-online-manage/lib/cmd"
)

//获取cpu使用率
func GetCPUUse() int {
	return int(C.cpu())

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
