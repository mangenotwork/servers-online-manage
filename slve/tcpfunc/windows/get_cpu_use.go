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
*/
import "C"

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