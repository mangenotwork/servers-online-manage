package windows

import (
	"testing"
	"time"
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

	// c1 := GetCpuVendorId1()
	// t.Log(c1)

	for i := 1; i < 5000; i++ {
		c1_1 := GetCpuVendorId2()

		t.Log(c1_1)
		c1_1 = ""
		time.Sleep(100 * time.Millisecond)
	}

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

	// d1 := WindowsGetDiskInfo()
	// for _,v := range d1{
	// 	t.Log(v)
	// 	t.Log(*v.DistUse)
	// }
}
