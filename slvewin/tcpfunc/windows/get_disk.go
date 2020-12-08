package windows

/*
#ifdef WIN32
#include "diskinfo.h"
#endif
*/
import "C"

import (
	"fmt"
	_ "log"
	"strings"
	"unsafe"

	"github.com/mangenotwork/servers-online-manage/lib/structs"

	"github.com/mangenotwork/servers-online-manage/lib/utils"
)

//磁盘类型枚举
var DiskType = map[int]string{
	0: "未知磁盘",
	1: "未知的磁盘类型",
	2: "路径无效",
	3: "可移动磁盘",
	4: "固定磁盘",
	5: "网络磁盘",
	6: "光驱",
	7: "内存映射盘",
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
func WindowsGetDiskUse(diskName string) (diskUse *structs.DiskUseInfo) {
	diskUse = &structs.DiskUseInfo{}
	cs := C.CString(diskName)
	rc := C.GoString(C.GetDrivesFreeSpace(cs))
	fmt.Println(rc)
	d_list := strings.Split(rc, "@")
	if len(d_list) != 2 {
		return
	}
	total := utils.Num2Int(d_list[0])
	free := utils.Num2Int(d_list[1])
	if total == 0 {
		return
	}
	rate := 100 - (float32(free)/float32(total))*100

	fmt.Println(total, free, rate)
	diskUse.Total = total
	diskUse.Free = free
	diskUse.Rate = rate
	return diskUse
}

//返回磁盘整体信息结构数据
func WindowsGetDiskInfo() (datas []*structs.DiskInfo) {
	datas = make([]*structs.DiskInfo, 0)
	for _, v := range WindowsGetDiskNameList() {

		diskType := WindowsGetDiskType(v)
		diskUse := WindowsGetDiskUse(v)
		diskTotal := diskUse.Total
		datas = append(datas, &structs.DiskInfo{
			DiskName:    v,
			DistType:    diskType,
			DistUse:     diskUse,
			DistTotalMB: fmt.Sprintf("%d MB", diskTotal),
		})
	}
	return
}

//TODO: 获取磁盘的IO率
