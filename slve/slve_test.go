package slve

import (
	"testing"
)

func TestGetSystemUUID(t *testing.T) {
	a := GetSystemUUID()
	t.Log(a)
}

func TestProcCpuinfo(t *testing.T) {
	a := ProcCpuinfo()
	t.Log(a)
}

func TestProcMeminfo(t *testing.T) {
	a := ProcMeminfo()
	t.Log(*a)
}

func TestProcStat(t *testing.T) {
	// a := GetProcStat()
	// for _, v := range a {
	// 	t.Log(*v)
	// }
	ProcStat()
}

func TestProcessProcStat(t *testing.T) {
	//GetProcessProcStat()
	ProcStat()
	ProcessProcStat("1729")
}

func TestProcVersion(t *testing.T) {
	ProcDiskstats()

}

func TestProcNetDev(t *testing.T) {
	ProcNetDev()
}
