package windows

/*
#ifdef WIN32
#include "cpuinfo.h"
#include "cpuuse.h"
#endif
*/
import "C"

import (
	"unsafe"
)

//获取cpu的VendorId
func GetCpuVendorId() string {

	//cStr := C.WindowsGetCpuVendorId()
	var out *C.char = C.WindowsGetCpuVendorId()
	a := C.GoString(out)
	C.free(unsafe.Pointer(out))
	return a
}

func GetCpuVendorId1() string {
	return C.GoString(C.WindowsGetCpuVendorId())
}

//获取cpu的 CpuId
func GetCpuId() string {
	return C.GoString(C.WindowsGetCpuId())
}

//获取cup的 CpuName
func GetCpuName() string {
	return C.GoString(C.WindowsGetCpuName())
}

//获取cpu使用率
func GetCPUUse() int {
	return int(C.cpu())

}
