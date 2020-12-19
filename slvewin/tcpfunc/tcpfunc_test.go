package tcpfunc

import (
	"testing"
)

func TestDocker(t *testing.T) {
	// a1, err := GetInfo()
	// t.Log(a1, err)

	// a2, err := GETImageList()
	// t.Log(a2, err)

	a3, err := GETContainerList()
	t.Log(a3, err)
}
