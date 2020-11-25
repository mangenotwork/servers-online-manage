package global

import (
	"log"
	"sync"

	"github.com/mangenotwork/servers-online-manage/structs"
)

var (
	//Master
	Version string = "0.1"

	//Slve
	SlveVersion string = "0.1"

	//Master TCP Host
	MasterHost string = "192.168.0.9:8555"

	//用于保存slve的连接
	Slves = make(map[string]*structs.Cli, 0)

	//添加slve连接的全局锁
	SlvesLock = &sync.RWMutex{}

	//文件包的每一包大小
	FilePacketSize int64 = 2048

	//master颁发给slve的token
	SlveToken string

	//客户端保存文件包
	FilePackets = make([]*structs.SendFilePacket, 0)

	//客户端文件传输是否接收完成
	FileEnd bool = false

	//重连chan
	RConn = make(chan bool)

	//客户端创建的空间
	SlveSpace = "/home/ManGe_Slve/"
)

//添加Slve
func AddSlve(key string, slve *structs.Cli) {
	SlvesLock.Lock()
	defer SlvesLock.Unlock()
	Slves[key] = slve

}

//打印当前所有Slve
func PrintSlves() {
	SlvesLock.RLock()
	defer SlvesLock.RUnlock()
	log.Println(Slves)
}

//返回 Slve 长度
func SlveLen() int {
	SlvesLock.RLock()
	defer SlvesLock.RUnlock()
	return len(Slves)
}

//删除 Slve
func DelSlve(key string){
	SlvesLock.RLock()
	defer SlvesLock.RUnlock()
	delete(Slves,key)
}