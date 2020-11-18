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
