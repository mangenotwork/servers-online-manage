#include <stdio.h>
#include <windows.h>
#include <TlHelp32.h>
#include "str_list.h"

//获取进程
StringList* GetWindowsPID()
{	
	StringList *r = strlist_malloc();
	HANDLE hProcessSnap = NULL;
	PROCESSENTRY32 pe32 = {0};
 
	//在使用这个结构前，先设置它的大小
	pe32.dwSize = sizeof(PROCESSENTRY32);
 
	//给系统内所有的进程拍个快照
	hProcessSnap = CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS,0);
	if (hProcessSnap == INVALID_HANDLE_VALUE)
	{
		printf_s("CreatToolhelp32Snapshot error!\n");
		return r;
	}
 
	//遍历进程快照，轮流显示每个进程的信息
	BOOL bRet = Process32First(hProcessSnap,&pe32);
	while (bRet)
	{
		//printf_s("name: %s \n",pe32.szExeFile);   //这里得到的应该是宽字符，用%ls或者%S,不然无法正常打印
		//printf_s("pid:  %u\n\n",pe32.th32ProcessID);
		bRet = Process32Next(hProcessSnap,&pe32);
		char *pidStr = (char *) malloc(100);
		sprintf(pidStr,"%s:%u", pe32.szExeFile, pe32.th32ProcessID);
		strlist_add(r,pidStr);
	}
 
	//不要忘记清除掉snapshot对象
	CloseHandle(hProcessSnap);
	return r;
}