package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	"github.com/mangenotwork/servers-online-manage/lib/enum"
	"github.com/mangenotwork/servers-online-manage/lib/global"
	pk "github.com/mangenotwork/servers-online-manage/lib/packet"
	"github.com/mangenotwork/servers-online-manage/lib/protocol"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
	"github.com/mangenotwork/servers-online-manage/master/db"
	"github.com/mangenotwork/servers-online-manage/master/http"
	"github.com/mangenotwork/servers-online-manage/master/http/models"
	"github.com/mangenotwork/servers-online-manage/master/tcp"
)

func main() {

	//初始话配置文件
	InitConf()

	//启动Master Http服务
	go http.Httpserver()

	//运行Master
	RunMasterTCP()
}

//运行Master
func RunMasterTCP() {
	//类似于初始化套接字，绑定端口
	hawkServer, err := net.ResolveTCPAddr("tcp", global.MasterHost)
	checkErr(err)

	//侦听
	listen, err := net.ListenTCP("tcp", hawkServer)
	checkErr(err)

	//关闭
	defer listen.Close()

	tcpServer := &structs.TcpServer{
		Listener:   listen,
		HawkServer: hawkServer,
	}
	log.Println("start Master TCP server successful.")

	//接收请求
	for {
		//来自客户端的连接
		conn, err := tcpServer.Listener.Accept()
		checkErr(err)
		log.Println("accept tcp client ", conn.RemoteAddr().String(), conn)

		//创建新的客户端实例
		newCli := &structs.Cli{
			Conn:  conn,
			Rdata: make(chan interface{}),
		}

		//新客户端连接写入Slves
		ip := strings.Split(conn.RemoteAddr().String(), ":")[0]
		global.AddSlve(ip, newCli)
		//打印当前所有Slve
		global.PrintSlves()
		//生成通知
		notif := &models.Notifincation{
			Slve:  ip,
			Type:  enum.AlarmMessage,
			State: enum.AlarmUnread,
			Messg: fmt.Sprintf("IP:%s的Slve连接成功!", ip),
			Time:  time.Now().Unix(),
		}
		notif.Create()

		//客户端连接成功给他颁发一个Token
		packet := structs.Packet{
			PacketType:    pk.SET_SLVE_TOKEN_PACKET,
			PacketContent: []byte(getRandString()),
		}
		tcp.SendData(conn, packet)

		// 每次建立一个连接就放到单独的协程内做处理
		go Handle(newCli)
	}
}

func Handle(conn *structs.Cli) {
	// 连接断开的处理
	defer delete(global.Slves, conn.Conn.RemoteAddr().String())
	defer conn.Conn.Close()

	//Master接收的具体业务
	err := protocol.DePackSendDataMater(conn, tcp.MasterTcpFunc)
	if err != nil {
		log.Println(err)
		notif := &models.Notifincation{
			Slve:  strings.Split(conn.Conn.RemoteAddr().String(), ":")[0],
			Type:  enum.AlarmWarning,
			State: enum.AlarmUnread,
			Messg: err.Error(),
			Time:  time.Now().Unix(),
		}
		notif.Create()
	}
}

//处理错误
func checkErr(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}

//初始化配置
func InitConf() {
	//读取配置文件
	file, _ := os.Open("conf/master_conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	masterconf := structs.MasterConf{}
	err := decoder.Decode(&masterconf)
	if err != nil {
		log.Println("Error:", err)
	}
	log.Println("masterconf = ", &masterconf)

	//给全局变量赋值
	global.Version = masterconf.Version
	global.MasterHost = masterconf.MasterHost

	db.CheckSqlitDB(masterconf.SqlistDBFile)
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
