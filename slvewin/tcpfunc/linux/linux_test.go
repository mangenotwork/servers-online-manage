package linux

import "testing"

func TestRun(t *testing.T){
	//a,v := ProcVersion()
	//t.Log(a)
	//t.Log(v)

	//a1 := ProcCpuinfo()
	//for _, v1 := range a1 {
	//	t.Log(v1)
	//}

	//GetCPUIDFromLinux()

	//a2 := ProcMeminfo()
	//t.Log(a2)

	GetSystemDF()
}