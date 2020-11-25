package cmd

import (
	"io/ioutil"
	"log"
	"os/exec"
)

//Linux Send Command Linux执行命令
//command 要执行的命令
func LinuxSendCommand(command string) (opStr string) {
	cmd := exec.Command("/bin/bash", "-c", command)
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
