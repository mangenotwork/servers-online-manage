package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"time"

	"github.com/mangenotwork/servers-online-manage/lib/global"
	"github.com/mangenotwork/servers-online-manage/lib/protocol"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
	"github.com/mangenotwork/servers-online-manage/slve/handler"
)

func main() {

	//初始话配置
	InitConf()

	//用于重连
Reconnection:

	//拿到服务器地址信息
	hawkServer, err := net.ResolveTCPAddr("tcp", global.MasterHost)
	if err != nil {
		log.Printf("hawk server [%s] resolve error: [%s]", global.MasterHost, err.Error())
		os.Exit(1)
	}

	//连接服务器
	connection, err := net.DialTCP("tcp", nil, hawkServer)
	if err != nil {
		log.Printf("connect to hawk server error: [%s]", err.Error())
		time.Sleep(1 * time.Second)
		goto Reconnection
	}

	//创建客户端实例
	client := &structs.TcpClient{
		Connection: connection,
		HawkServer: hawkServer,
		StopChan:   make(chan struct{}),
	}

	//启动接收,并执行slve的业务
	go protocol.DePackSendData(client.Connection, handler.SlveTcpFunc)

	//发送心跳的goroutine
	go func() {
		//测试是3秒
		//非测试则调整到 > 30 秒
		heartBeatTick := time.Tick(3 * time.Second)
		for {
			select {
			case <-heartBeatTick:
				handler.SendHeartPacket(client)
			case <-client.StopChan:
				return
			}
		}
	}()

	// //测试用的，开300个goroutine每秒发送一个包
	// for i := 0; i < 300; i++ {
	// 	go func() {
	// 		sendTimer := time.After(1 * time.Second)
	// 		for {
	// 			select {
	// 			case <-sendTimer:
	// 				client.sendReportPacket()
	// 				sendTimer = time.After(1 * time.Second)
	// 			case <-client.stopChan:
	// 				return
	// 			}
	// 		}
	// 	}()
	// }

	// 发送数据包测试
	// go func() {
	// 	sendTimer := time.After(3 * time.Second)
	// 	for {
	// 		select {
	// 		case <-sendTimer:
	// 			client.sendReportPacket()
	// 			sendTimer = time.After(1 * time.Second)
	// 		case <-client.stopChan:
	// 			return
	// 		}
	// 	}
	// }()

	for {
		select {
		case a := <-global.RConn:
			log.Println("global.RConn = ", a)
			goto Reconnection
		}
	}

	//等待退出
	<-client.StopChan
}


//初始化配置
func InitConf() {
	//读取配置文件
	file, _ := os.Open("conf/slve_conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	slveconf := structs.SlveConf{}
	err := decoder.Decode(&slveconf)
	if err != nil {
		log.Println("Error:", err)
		log.Println("读取配置文件 conf/slve_conf.json 失败 ！")
		os.Exit(1)
	}
	log.Println("slveconf = ", &slveconf)

	//给全局变量赋值
	global.SlveVersion = slveconf.Version
	global.MasterHost = slveconf.MasterHost
	global.SlveSpace = slveconf.SlveSpace

	//检查空间
	os.Mkdir(global.SlveSpace, os.ModePerm)

}
