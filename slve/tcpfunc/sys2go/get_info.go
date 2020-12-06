package sys2go

import (
	"fmt"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
	"log"
	"net"
	"os"
	"runtime"
)

//获取本机ip
func GetMyIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Println("[Error] :", err)
		return ""
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.String()
}

//获取host 系统类型
func GetSysType() string {
	return runtime.GOOS
}

//获取host 命名
func GetHostName() string {
	name, err := os.Hostname()
	if err != nil {
		name = "null"
	}
	return name
}

//获取系统信息
func SysInfo() {
	log.Println(`系统类型：`, runtime.GOOS)
	log.Println(`系统架构：`, runtime.GOARCH)
	log.Println(`CPU 核数：`, runtime.GOMAXPROCS(0))
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	log.Println(`电脑名称：`, name)
}

//获取host的基本信息
func GetHostInfo() *structs.HostInfo {
	sysType := runtime.GOOS
	//获取cpu信息

	return &structs.HostInfo{
		HostName:      GetHostName(),
		SysType:       sysType,
		SysArch:       runtime.GOARCH,
		CpuCoreNumber: fmt.Sprintf("cpu 核心数: %d", runtime.GOMAXPROCS(0)),
	}
}
