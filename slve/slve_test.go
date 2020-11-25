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

func TestProcNetSnmp(t *testing.T) {
	//ProcNetSnmp()
	//ProcPIDCmdline("221")
	//ProcPIDEnviron("11606")
	//ProcPIDLimits("11606")
	//ProcPIDMaps("11606")
	//ProcPIDStatus("11606")
	//ProcCrypto()
	//ProcModules()
	ProcUptime()
}

func TestHaveDocker(t *testing.T) {
	// isDocker, pid := HaveDocker()
	// t.Log(isDocker, pid)

	// version := CmdDockerVersion()
	// t.Log(version)

	// path := DockerFragmentPath()
	// t.Log(path)

	// file := CatDockerFragmentPath()
	// t.Log(file)

	// isopen, url := DockerIsOpenAPI()
	// t.Log(isopen, url)

	//OpenDockerAPI()
	CloseDockerAPI()
}

func TestDockerAPI(t *testing.T) {
	//data1,err := ImageList()
	//t.Log(string(data1), err)

	//ImageCreate()

	//ImageSearch()

	//ImageDelete()

	Run1()
}
