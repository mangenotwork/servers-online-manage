// linux平台上的docker相关的命令，文件内容等操作
package software

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/mangenotwork/servers-online-manage/lib/cmd"
	"github.com/mangenotwork/servers-online-manage/lib/utils"
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
	path = utils.DeletePreAndSufSpace(rStrList[1])
	path = strings.Trim(path, "\n")
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
//修改 docker 配置的api tcp地址
//1. 打开docker配置文件
//2. 找到 ExecStart 所在的行，在行尾追加 -H tcp://0.0.0.0:5678
//3. 重新加载配置文件，重启docker daemon
// a) sudo systemctl daemon-reload
// b) sudo systemctl restart docker
//4. 检查省份开启 docker -H localhost:5678 version
func OpenDockerAPI() {
	port := "12225"
	//检查断开是否占用
	rStr := cmd.LinuxSendCommand("lsof -i:" + port)
	log.Println("rStr : ", rStr)
	if rStr != "" {
		log.Println("端口已经被占用")
		return
	}

	s := " -H tcp://0.0.0.0:" + port
	ModifyDockerExecStart(s)
	RestartDocker()
	//检查是否生效
	//docker -H 0.0.0.0:12221 info
	//没有 Cannot connect to the Docker daemo 则成功
	testopen := cmd.LinuxSendCommand("docker -H 0.0.0.0:" + port + " info")
	if t := strings.Contains(testopen, "Cannot connect to the Docker daemo"); t {
		log.Println("docker Api 启动失败！")
	} else {
		log.Println("docker Api 启动成功！")
	}

}

//关闭docker api
func CloseDockerAPI(){
	ModifyDockerExecStart("")
	RestartDocker()

}

//重启docker
func RestartDocker(){
	//sudo systemctl daemon-reload
	reloadStr := cmd.LinuxSendCommand("sudo systemctl daemon-reload")
	log.Println("reloadStr : ", reloadStr)

	//sudo systemctl restart docker
	restartStr := cmd.LinuxSendCommand("sudo systemctl restart docker")
	log.Println("reloadStr : ", restartStr)
}

//修改docker的配置文件中的 ExecStart
//主要用于 开启docker api, 改修docker api, 关闭docker api
func ModifyDockerExecStart(s string){
	//docker的配置文件
	//dockerFileName := "/lib/systemd/system/docker.service"
	dockerFileName := DockerFragmentPath()
	file, err := os.Open(dockerFileName)
	if err != nil {
		log.Println("open file fail:", err)
		return
	}
	defer file.Close()

	//新建一个接收文件
	outFilename := dockerFileName + ".new"
	out, err := os.OpenFile(outFilename, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		log.Println("Open write file fail:", err)
		os.Exit(-1)
	}
	defer out.Close()

	br := bufio.NewReader(file)
	index := 1
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("read err:", err)
			return
		}
		lineStr := string(line)
		log.Println(lineStr, index)

		if len(lineStr) > 9 && lineStr[0:9] == "ExecStart" {

			//判断是否已经有 tcp://
			vList := strings.Split(lineStr, " ")
		Loop:
			for i := 0; i < len(vList); i++ {
				if vList[i] == "-H" && strings.Contains(vList[i+1], "tcp://") {
					vList = append(vList[:i], vList[i+2:]...)
					goto Loop
				}
			}
			newStr := strings.Join(vList, " ")
			log.Println("追加")
			lineStr = newStr + s
			log.Println(lineStr)
		}

		_, err = out.WriteString(lineStr + "\n")
		if err != nil {
			log.Println("write to file fail:", err)
			return
		}

		index++
	}

	//删除
	err = os.Remove(dockerFileName)
	if err != nil {
		log.Println("删除失败:", err)
		return
	}
	//改名
	err = os.Rename(outFilename, dockerFileName)
	if err != nil {
		log.Println("改名失败:", err)
		return
	}
}