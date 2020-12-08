/*
    获取cpu的使用率
    原理：GetSystemTimes,
*/

#include <stdio.h>
#include <windows.h>
#include <conio.h>

double FileTimeToDouble(FILETIME* pFiletime)
{
	return (double)((*pFiletime).dwHighDateTime * 4.294967296E9) + (double)(*pFiletime).dwLowDateTime;
}

double m_fOldCPUIdleTime;
double m_fOldCPUKernelTime;
double m_fOldCPUUserTime;

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