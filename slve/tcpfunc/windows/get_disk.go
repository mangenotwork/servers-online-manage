package windows

/*
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

*/
import "C"

import (
	"fmt"
	"unsafe"
	"strings"

	"github.com/mangenotwork/servers-online-manage/lib/utils"
)

//磁盘类型枚举
var DiskType = map[int]string{
	0:"未知磁盘",
	1:"未知的磁盘类型",
	2:"路径无效",
	3:"可移动磁盘",
	4:"固定磁盘",
	5:"网络磁盘",
	6:"光驱",
	7:"内存映射盘",
}

//磁盘信息
type DiskInfo struct {
	DiskName string
	DistType string
	DistTotalMB string
	DistUse *DiskUseInfo
}

//磁盘使用的信息
type DiskUseInfo struct {
	Total int //MB
	Free int //MB
	Rate float32 //%
}

//返回磁盘数量
func WindowsGetDiskCount() int {
	return int(C.GetLogicalDriveCount())
}

//返回磁盘名列表
func WindowsGetDiskNameList() []string {
	clist := C.GetDriveNameList()
	size := int(clist.size)
	var b *C.char
	ptrSize := unsafe.Sizeof(b)
	gostring := make([]string, size)
	if size > 0 {
		for i := 0; i < size; i++ {
			//这个方法是解析c 的字符串链表
			/*
			   typedef struct{
			       unsigned int size;        //子字符串数量
			       char **list;            //用字符串数组来存放字符串列表
			   }st_strlist;
			*/
			//(*C.char)(*(**C.char)(unsafe.Pointer(uintptr( unsafe.Pointer(job.exHosts)) + uintptr(1)*ptrSize ) ) ))
			element := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(clist.list)) + uintptr(i)*ptrSize))
			gostring[i] = C.GoString((*C.char)(*element))
		}
	}
	fmt.Println(gostring)
	return gostring
}

//更具磁盘名获取磁盘类型
func WindowsGetDiskType(diskName string) string {
	cs := C.CString(diskName)
	return DiskType[int(C.GetDrivesType(cs))]
}

//更具磁盘名获取磁盘使用信息
func WindowsGetDiskUse(diskName string) (diskUse *DiskUseInfo) {
	diskUse = &DiskUseInfo{}
	cs := C.CString(diskName)
	rc := C.GoString(C.GetDrivesFreeSpace(cs))
	fmt.Println(rc)
	d_list := strings.Split(rc,"@")
	if len(d_list) != 2{
		return
	}
	total :=  utils.Num2Int(d_list[0])
	free := utils.Num2Int(d_list[1])
	if total == 0{
		return
	}
	rate := 100 - (float32(free) / float32(total))*100

	fmt.Println(total,free,rate)
	diskUse.Total = total
	diskUse.Free = free
	diskUse.Rate = rate
	return diskUse
}

//返回磁盘整体信息结构数据
func WindowsGetDiskInfo() (datas []*DiskInfo) {
	datas = make([]*DiskInfo,0)
	for _,v := range WindowsGetDiskNameList(){

		diskType := WindowsGetDiskType(v)
		diskUse := WindowsGetDiskUse(v)
		diskTotal := diskUse.Total
		datas = append(datas,&DiskInfo{
			DiskName: v,
			DistType: diskType,
			DistUse: diskUse,
			DistTotalMB: fmt.Sprintf("%d MB",diskTotal),
		})
	}
	return
}