// linux平台上的docker相关的命令，文件内容等操作
// TODO  需要改为适配 Windows平台的
package software

import (
	// "bufio"
	// "io"
	"log"

	// "os"
	// "regexp"
	"strings"

	"github.com/mangenotwork/servers-online-manage/lib/cmd"
	"github.com/mangenotwork/servers-online-manage/lib/docker"
	"github.com/mangenotwork/servers-online-manage/lib/utils"
)

//查看windows是否启动了docker 服务
//使用net start 命令
func IsOpenDockerServer() (bool, []string) {
	cmds := []string{"cmd", "/c", "net start"}
	rStr := cmd.WindowsSendCommand(cmds)
	//log.Println(rStr)
	rStrList := strings.Split(rStr, "\n")
	dockerServers := make([]string, 0)
	for _, v := range rStrList {
		if strings.Index(strings.ToLower(v), "docker") != -1 {
			dockerServers = append(dockerServers, utils.DeletePreAndSufSpace(v))
		}
	}
	have := false
	if len(dockerServers) > 0 {
		have = true
	}
	log.Println(have, dockerServers)
	return have, dockerServers
}

//通过 docker version 返回 docker 信息
func CmdDockerVersion() string {
	cmds := []string{"cmd", "/c", "docker version"}
	rStr := cmd.WindowsSendCommand(cmds)
	//log.Println(rStr)
	return rStr
}

//通过 docker info  获取信息
func CmdDockerInfo() string {
	cmds := []string{"cmd", "/c", "docker info"}
	rStr := cmd.WindowsSendCommand(cmds)
	return rStr
}

//Docker 是否可用
func IsOpenDockerAPI() bool {
	context, dockerC, err := docker.GetClient()
	defer dockerC.Close()
	log.Println(context, dockerC, err)
	if err != nil {
		return false
	}
	return true
}
