/*
    获取windos 系统的信息
    原理: 主要通过 windows.h的使用
*/

#include <windows.h>
#include <stdio.h>
#include <conio.h>
#include "metrics.h"

char* ComputerNameStr;//计算机名称
char* EnvironmentStr;//计算机环境变量

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
