package cmd

import (
	"io/ioutil"
	"log"
	"os/exec"
)

//Windows Send Command Linux执行命令
//command 要执行的命令
func WindowsSendCommand(command []string) (opStr string) {
	if len(command) < 1 {
		return ""
	}
	cmd := exec.Command(command[0], command[1:len(command)]...)
	stdout, stdoutErr := cmd.StdoutPipe()
	if stdoutErr != nil {
		log.Fatal("ERR stdout : ", stdoutErr)
	}
	defer stdout.Close()
	if startErr := cmd.Start(); startErr != nil {
		log.Fatal("ERR Start : ", startErr)
	}
	opBytes, opBytesErr := ioutil.ReadAll(stdout)
	if opBytesErr != nil {
		//log.Println(string(opBytes))
		opStr = ""
	}
	opStr = string(opBytes)
	//log.Println(opStr)
	return
}