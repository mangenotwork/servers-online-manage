package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/mangenotwork/servers-online-manage/lib/global"
	pk "github.com/mangenotwork/servers-online-manage/lib/packet"
	"github.com/mangenotwork/servers-online-manage/lib/protocol"
	"github.com/mangenotwork/servers-online-manage/slve/tcp"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
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
	go protocol.DePackSendData(client.Connection, tcp.SlveTcpFunc)

	//发送心跳的goroutine
	go func() {
		heartBeatTick := time.Tick(3 * time.Second)
		for {
			select {
			case <-heartBeatTick:
				SendHeartPacket(client)
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

//发送数据包
//仔细看代码其实这里做了两次json的序列化，有一次其实是不需要的
func sendReportPacket(client *structs.TcpClient) {
	reportPacket := structs.ReportPacket{
		Content:   getRandString(),
		Timestamp: time.Now().Unix(),
		Rand:      rand.Int(),
	}
	packetBytes, err := json.Marshal(reportPacket)
	if err != nil {
		log.Println(err.Error())
	}
	//这一次其实可以不需要，在封包的地方把类型和数据传进去即可
	packet := structs.Packet{
		PacketType:    pk.REPORT_PACKET,
		PacketContent: packetBytes,
	}
	sendBytes, err := json.Marshal(packet)
	if err != nil {
		log.Println(err.Error())
	}
	//发送
	client.Connection.Write(protocol.EnPackSendData(sendBytes))
	//log.Println("Send metric data success!")
}

//发送心跳包，与发送数据包一样
func SendHeartPacket(client *structs.TcpClient) {
	heartPacket := structs.HeartPacket{
		Version:   global.SlveVersion,
		SlveId:    global.SlveToken,
		IP:        tcp.GetMyIP(),
		System:    tcp.GetSysType(),
		HostName:  tcp.GetHostName(),
		UseCPU:    "28%",
		UseMEM:    "28%",
		Timestamp: time.Now().Unix(),
	}
	packetBytes, err := json.Marshal(heartPacket)
	if err != nil {
		log.Println(err.Error())
	}
	packet := structs.Packet{
		PacketType:    pk.HEART_BEAT_PACKET,
		PacketContent: packetBytes,
	}
	sendBytes, err := json.Marshal(packet)
	if err != nil {
		log.Println(err.Error())
	}
	client.Connection.Write(protocol.EnPackSendData(sendBytes))
}

//拿一串随机字符
func getRandString() string {
	length := rand.Intn(50)
	strBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		strBytes[i] = byte(rand.Intn(26) + 97)
	}
	return string(strBytes)
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
