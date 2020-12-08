package sys2go

import (
	"fmt"
	"github.com/mangenotwork/servers-online-manage/lib/cmd"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
	"github.com/mangenotwork/servers-online-manage/lib/utils"
	"log"
	"net"
	"os"
	"regexp"
	"runtime"
	"time"
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

//获取系统架构
func GetSysArch() string {
	return runtime.GOARCH
}

//获取host 命名
func GetHostName() string {
	name, err := os.Hostname()
	if err != nil {
		name = "null"
	}
	return name
}

//获取cpu核心数
func GetCpuCoreNumber() string {
 return fmt.Sprintf("%d核", runtime.GOMAXPROCS(0))
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

//获取网卡信息
func GetNetInfo() error{
	intf, err := net.Interfaces()
	if err != nil {
		log.Println("get network info failed: ", err)
		return err
	}
	for _,v := range intf{
		log.Println("最大传输单元 = ",v.MTU)
		log.Println("Name = ",v.Name)
		log.Println("硬件地址 = ", v.HardwareAddr)
		log.Println("接口的属性 = ", v.Flags)
		/*
		"up",  接口在活动状态
			"broadcast",   接口支持广播
			"loopback",  接口是环回的
			"pointtopoint",  接口是点对点的
			"multicast",  接口支持组播
		 */
		ips, err := v.Addrs()
		if err != nil {
			log.Println("get network addr failed: ", err)
			return err
		}
		log.Println("ips = ", ips)
		mips, err := v.MulticastAddrs()
		if err != nil {
			log.Println("get network addr failed: ", err)
			return err
		}
		log.Println("mips = ", mips)
	}
	return nil
}

//执行 netstat -e  返回接收字节和发送字节
func RunNetstatE() (input int64,output int64){
	input,output = 0,0
	cmds := []string{"netstat", "-e"}
	rStr := cmd.WindowsSendCommand(cmds)
	reg := regexp.MustCompile(`[0-9]+`)
	sList := reg.FindAllString(rStr, -1)
	if len(sList) >= 2 {
		input = utils.Str2Int64(sList[0])
		output = utils.Str2Int64(sList[1])
	}
	log.Println("input = ", input)
	log.Println("output = ", output)
	return
}

//采集网络带宽
//通过netstat -e  命令进行采集
func GetNetIOFromCMD(){
	t1 := time.Now().UnixNano()
	input1,output1 := RunNetstatE()
	time.Sleep(500*time.Millisecond)
	t2 := time.Now().UnixNano()
	input2,output2 := RunNetstatE()
	t := (t2 - t1)/1000/1000
	input := input2 - input1
	output := output2 - output1
	inputms := ((float32(input)/1024)/float32(t))*1000
	outputms := ((float32(output)/1024)/float32(t))*1000
	log.Println("t = ", t ," ms")
	log.Println("input = ", input)
	log.Println("output = ", output)
	log.Println("inputms = ", inputms, "kb/s")
	log.Println("outputms = ", outputms, "kb/s")
}

