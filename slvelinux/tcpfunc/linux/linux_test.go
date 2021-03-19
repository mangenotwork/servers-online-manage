package linux

import (
	"testing"
)

func TestRun(t *testing.T) {
	//a,v := ProcVersion()
	//t.Log(a)
	//t.Log(v)

	//a1 := ProcCpuinfo()
	//for _, v1 := range a1 {
	//	t.Log(v1)
	//}

	GetCPUIDFromLinux()

	//a2 := ProcMeminfo()
	//t.Log(a2)

	//a3, a4 := GetSystemDF()
	//t.Log(a3)
	//t.Log(a4)

	//a5,a6 := ProcStat(500 * time.Millisecond)
	//t.Log(*a5, a6)

	//ProcDiskstats()

	//ProcNetDev(500 * time.Millisecond)

	//ProcNetSnmp()

	//a7 := GetTcpConnCount()
	//t.Log(a7)

	//a8 := GetProcessCount()
	//t.Log(a8)

	//a9 := GetProcessList()
	//t.Log(len(a9),a9)

	//GetSysEnv()
}
