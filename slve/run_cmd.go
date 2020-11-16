//slve 执行命令
package slve

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunInLinux(cmd string) string {
	fmt.Println("Running Linux cmd:", cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result))
}
