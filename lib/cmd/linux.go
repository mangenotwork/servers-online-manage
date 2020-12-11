package cmd

import (
	"io/ioutil"
	"log"
	"os/exec"
	"context"
)

//Linux Send Command Linux执行命令
//command 要执行的命令
func LinuxSendCommand(command string) (opStr string) {
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx,"/bin/bash", "-c", command)
	stdout, stdoutErr := cmd.StdoutPipe()
	if stdoutErr != nil {
		log.Fatal("ERR stdout : ", stdoutErr)
	}
	defer cancel()
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
	cmd.Wait()
	return
}
