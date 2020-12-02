package windows

import (
	"testing"
)

func TestGetCPUUse(t *testing.T) {
	//a := GetCPUUse()
	//t.Log(a)
	//
	//a2 := GetComputerName()
	//t.Log(a2)
	//
	//a3 := GetEnvironment("PATH")
	//t.Log(a3)

	//GetSystemInfo()
	//RunMetrics()

	//GETSystemPowerStatus()

	//WindowsGetUserDefaultLangID()

	c1 := GetCpuVendorId()
	t.Log(c1)

	c2 := GetCpuId()
	t.Log(c2)

	c3 := GetCpuName()
	t.Log(c3)
}