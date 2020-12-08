package windows

import (
	"testing"
)

func TestGetCPUUse(t *testing.T) {

	//
	// a2 := GetComputerName()
	// t.Log(a2)

	// a3 := GetEnvironment("PATH")
	// t.Log(a3)

	// GetSystemInfo()

	// GETSystemPowerStatus()

	// WindowsGetUserDefaultLangID()

	// c1 := GetCpuVendorId()
	// t.Log(c1)

	// c2 := GetCpuId()
	// t.Log(c2)

	// c3 := GetCpuName()
	// t.Log(c3)

	// a := GetCPUUse()
	// t.Log(a)

	// c4 := RunMetrics()
	// t.Log(c4)

	// WindowsGetOsInfo()
	// WindowsGetMemoryInfo()
	// GetBaseBoardID()

	// WindowsGetDiskCount()
	// WindowsGetDiskNameList()

	// WindowsGetDiskType(`C:\`)
	// WindowsGetDiskUse(`C:\`)

	d1 := WindowsGetDiskInfo()
	for _, v := range d1 {
		t.Log(v)
		t.Log(*v.DistUse)
	}

	//GetTcpConnCount()

	a5, a6 := GetWindowsPIDInfo()
	t.Log(a5, a6)
}
