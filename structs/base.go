package structs

import (
	"net"
)

//与服务器相关的资源都放在这里面
type TcpServer struct {
	Listener   *net.TCPListener
	HawkServer *net.TCPAddr
}

//与客户端相关的资源都放在这里面
type TcpClient struct {
	Connection *net.TCPConn
	HawkServer *net.TCPAddr
	StopChan   chan struct{}
}

//数据包
type Packet struct {
	PacketType    byte
	PacketContent []byte
}

//客户机心跳包，定时上传心跳数据
type HeartPacket struct {
	Version   string `json:"version"`
	SlveId    string `json:"slve_id"`
	IP        string `json:"slve_ip"` //ip+port
	System    string `json:"system"`
	HostName  string `json:"host_name"`
	UseCPU    string `json:"use_mem"`
	UseMEM    string `json:"use_mem"`
	Timestamp int64  `json:"timestamp"`
}

//数据包
type ReportPacket struct {
	Content   string `json:"content"`
	Rand      int    `json:"rand"`
	Timestamp int64  `json:"timestamp"`
}

//发送文件的数据结构
type SendFilePacket struct {
	// 文件名
	FileName string

	// 文件后缀
	FileSuffix string

	// 文件最大分包
	MaxPacketNum int64

	// 文件包计数
	FilePacketNum int

	// 文件包
	FilePacket []byte

	//结束通知 true结束
	IsEnd bool
}

//客户端
type Cli struct {
	Conn  net.Conn
	Rdata chan interface{}
}

//Master conf
type MasterConf struct {
	Version    string `json:"version"`
	MasterHost string `json:"master_host"`
}

//Slve conf
type SlveConf struct {
	Version    string `json:"version"`
	MasterHost string `json:"master_host"`
	SlveSpace  string `json:"slve_space"`
}
