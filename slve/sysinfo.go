//获取计算机的系统信息的实现
package slve

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"

	"github.com/mangenotwork/csdemo/structs"
)

//获取本机ip
func GetMyIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Println("[Error] :", err)
		return ""
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println(localAddr.String())
	// ip := strings.Split(localAddr.String(), ":")[0]
	return localAddr.String()

}

func GetSysType() string {
	return runtime.GOOS
}

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

func GetHostInfo() *structs.HostInfo {
	return &structs.HostInfo{
		HostName:      GetHostName(),
		SysType:       runtime.GOOS,
		SysArch:       runtime.GOARCH,
		CpuCoreNumber: fmt.Sprintf("cpu 核心数: %d", runtime.GOMAXPROCS(0)),
	}
}
