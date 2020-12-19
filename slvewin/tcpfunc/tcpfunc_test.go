package tcpfunc

import (
	"testing"
)

func TestDocker(t *testing.T) {
	a1, err := GetInfo()
	t.Log(a1, err)
}
