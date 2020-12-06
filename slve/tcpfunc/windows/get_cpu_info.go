package windows
import "C"

/*
#include "cpuinfo.h"
#include "cpuuse.h"
*/
import "C"

//获取cpu的VendorId
func GetCpuVendorId() string{
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