package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/mangenotwork/servers-online-manage/lib/global"
	"github.com/mangenotwork/servers-online-manage/lib/protocol"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
	"github.com/mangenotwork/servers-online-manage/lib/utils"
	"github.com/mangenotwork/servers-online-manage/slve/handler"
)

func main() {

	//初始话配置
	InitConf()

	//启动一个文件服务
	go FileServer()

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
		//5秒发送一次心跳
		heartBeatTick := time.Tick(60 * time.Second)
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
	utils.InitSlveConf()

	//检查空间
	os.Mkdir(global.SlveSpace, os.ModePerm)

}

//启动一个文件服务
func FileServer() {
	var (
		listen = flag.String("listen", ":18383", "listen address")
		dir    = flag.String("dir", "C://", "directory to serve")
	)
	log.Printf("listening on %q...", *listen)
	err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
	log.Fatalln(err)
}
