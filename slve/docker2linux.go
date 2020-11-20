// linux平台上的docker相关的命令，文件内容等操作
package slve

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/mangenotwork/csdemo/lib/cmd"
)

//host 是否安装docker
//执行 ps  -ef | grep docker 命令 如果返回有2条以上则安装
//ps  -aux | grep docker 也可以
//返回第二个参数是pid 字符串类型
func HaveDocker() (bool, string) {
	rStr := cmd.LinuxSendCommand("ps -e|grep docker")
	if rStr == "" {
		return false, ""
	}
	//log.Println(rStr)
	rStrList := strings.Split(rStr, "\n")
	rList := make([]string, 0)
	for _, v := range rStrList {
		if v != "" {
			rList = append(rList, v)
		}
	}
	//log.Println(rList, len(rList))
	if len(rList) >= 1 {
		pid := ""
		pidList := strings.Split(rList[0], " ")
		for _, v := range pidList {
			if v != "" {
				pid = v
				break
			}
		}
		return true, pid
	}
	return false, ""
}

//docker 版本  通过 docker version 命令获取
func CmdDockerVersion() (version string) {
	isDocker, _ := HaveDocker()
	if !isDocker {
		return
	}
	rStr := cmd.LinuxSendCommand("docker version")
	if rStr == "" {
		return
	}
	//log.Println(rStr)
	reg := regexp.MustCompile(`Version:(.*?)\n`)
	sList := reg.FindAllString(rStr, -1)
	if len(sList) > 1 {
		version = sList[0]
	}

	return
}

//docker 配置文件在哪里
//systemctl show --property=FragmentPath docker
func DockerFragmentPath() (path string) {
	isDocker, _ := HaveDocker()
	if !isDocker {
		return
	}
	rStr := cmd.LinuxSendCommand("systemctl show --property=FragmentPath docker")
	if rStr == "" {
		return
	}
	rStrList := strings.Split(rStr, "=")
	if len(rStrList) < 2 {
		return
	}
	path = rStrList[1]
	return
}

// 查看docker 配置文件
func CatDockerFragmentPath() string {
	path := DockerFragmentPath()
	if path == "" {
		return ""
	}
	return cmd.LinuxSendCommand("cat " + path)
}

//docker api 是否开启
//这里采用查看docker 配置文件ExecStart 的参数每个 -H
//如果有tcp 说明开启，并返回api地址
func DockerIsOpenAPI() (isOpen bool, url string) {
	isOpen = false
	url = ""
	conf := CatDockerFragmentPath()
	rStrList := strings.Split(conf, "\n")
	for _, v := range rStrList {
		if len(v) > 9 && v[0:9] == "ExecStart" {
			//v = "ExecStart=/usr/bin/dockerd -H unix:///var/run/docker.sock -H tcp://0.0.0.0:5678"
			vList := strings.Split(v, " ")
			for i := 0; i < len(vList); i++ {
				if vList[i] == "-H" {
					if strings.Contains(vList[i+1], "tcp://") {
						//log.Println(vList[i], vList[i+1])
						isOpen = true
						url = vList[i+1][6:len(vList[i+1])]
						return
					}
				}
			}
		}
	}
	return
}

//开启 docker http api
//1. 打开docker配置文件
//2. 找到 ExecStart 所在的行，在行尾追加 -H tcp://0.0.0.0:5678
//3. 重新加载配置文件，重启docker daemon
// a) sudo systemctl daemon-reload
// b) sudo systemctl restart docker
//4. 检查省份开启 docker -H localhost:5678 version
func OpenDockerAPI() {
	dockerFileName := "/home/mange/Desktop/go/src/github.com/mangenotwork/servers-online-manage/docker1.service"
	file, err := os.Open(dockerFileName)
	if err != nil {
		log.Println("open file fail:", err)
		return
	}
	defer file.Close()
}
