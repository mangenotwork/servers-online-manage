/*
    获取系统的磁盘信息
    原理: GetDriveType, GetDiskFreeSpaceEx, GetLogicalDrives, GetLogicalDriveStrings
*/

#include <stdio.h>
#include <Windows.h>
#include "str_list.h"

//获取磁盘类型
int GetDrivesType(const char *input){
    UINT uDriverType = GetDriveType(input);
    int type_id = 0;
    switch(uDriverType) {
    case DRIVE_UNKNOWN: type_id = 1; break;
    case DRIVE_NO_ROOT_DIR: type_id = 2; break;
    case DRIVE_REMOVABLE: type_id = 3; break;
    case DRIVE_FIXED: type_id = 4; break;
    case DRIVE_REMOTE: type_id = 5; break;
    case DRIVE_CDROM: type_id = 6; break;
    case DRIVE_RAMDISK: type_id = 7; break;
    default:
        break;
    }
    return type_id;
}

//获取磁盘的使用信息
char* GetDrivesFreeSpace(const char* lpRootPathName){
    char *space = (char *) malloc(100);
    unsigned long long available,total,free;
    if(GetDiskFreeSpaceEx(lpRootPathName,(ULARGE_INTEGER*)&available,(ULARGE_INTEGER*)&total,(ULARGE_INTEGER*)&free)){
        //printf("Drives %s | total = %lld MB,available = %lld MB,free = %lld MB\n",
        //        lpRootPathName,total>>20,available>>20,free>>20);
        sprintf(space,"%lld@%lld", total>>20, free>>20);
    }
    return space;
}

//获取磁盘个数
int GetLogicalDriveCount(){
    int driveCount = 0;
    char szDriveInfo[16 + 1] = {0};
    DWORD driveInfo = GetLogicalDrives();
    itoa(driveInfo, szDriveInfo, 2);
    while (driveInfo) {
        if (driveInfo & 1) {
            driveCount++;
        }
        driveInfo >>= 1;
    }
    return driveCount;
}

//获取磁盘名
StringList* GetDriveNameList(){
    int driveCount = GetLogicalDriveCount();
    StringList *r = strlist_malloc();
    DWORD dwSize = MAX_PATH;
    char szLogicalDrives[MAX_PATH] = {0};
    //获取逻辑驱动器号字符串
    DWORD dwResult = GetLogicalDriveStrings(dwSize,szLogicalDrives);
    if (dwResult > 0 && dwResult <= MAX_PATH) {
        char* szSingleDrive = szLogicalDrives;  //从缓冲区起始地址开始
        int n = 0;
        while(n<driveCount) {
            strlist_add(r,szSingleDrive);
            // 获取下一个驱动器号起始地址
            szSingleDrive += strlen(szSingleDrive) + 1;
            n++;
        }
    }
    return r;
}