package handler

import (
	"encoding/json"
	"log"
	"math/rand"
	"net"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/mangenotwork/servers-online-manage/lib/cmd"
	"github.com/mangenotwork/servers-online-manage/lib/global"
	pk "github.com/mangenotwork/servers-online-manage/lib/packet"
	"github.com/mangenotwork/servers-online-manage/lib/protocol"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
	"github.com/mangenotwork/servers-online-manage/slve/tcpfunc"
	"github.com/mangenotwork/servers-online-manage/slve/tcpfunc/sys2go"
)

var lock = &sync.Mutex{}

//具体执行的业务
func SlveTcpFunc(conn net.Conn, packet *structs.Packet) {
	switch packet.PacketType {

	//接收服务端颁发的Token
	case pk.SET_SLVE_TOKEN_PACKET:
		log.Println("接收服务端颁发的名称 : ", string(packet.PacketContent))
		global.SlveToken = string(packet.PacketContent)
		//回复Master的一包，包含所有信息
		hostinfo := sys2go.GetHostInfo()
		sysinfo := tcpfunc.SysInfos()

		//TODO 计算 uuid   由计算机唯一id得到

		packetData := &structs.SlveBaseInfo{
			Token:       global.SlveToken,
			Name:        hostinfo.HostName,
			SysType:     hostinfo.SysType,
			SysInfo:     sysinfo,
			SlveVersion: global.SlveVersion,
			SlveUUID:    global.SlveUUID,
		}
		SendPackat(conn, packetData, pk.FIRST_PACKET)
		return

	case pk.REPLY_HEART_PACKET:
		log.Println("Master 回复 : ", string(packet.PacketContent))
		return

	//返回slve信息给master
	case pk.Get_SLVE_INFO_PACKET:
		log.Println("返回slve信息给master ")
		hostinfo := sys2go.GetHostInfo()
		packetBytes, err := json.Marshal(hostinfo)
		if err != nil {
			log.Println(err.Error())
		}
		packet := structs.Packet{
			PacketType:    pk.RECEPTION_SLVE_PACKET,
			PacketContent: packetBytes,
		}
		sendBytes, err := json.Marshal(packet)
		if err != nil {
			log.Println(err.Error())
		}
		conn.Write(protocol.EnPackSendData(sendBytes))
		return

	//接收master 给的命令，执行后返回数据
	case pk.SET_SLVE_CMD_PACKET:
		cmdStr := string(packet.PacketContent)
		log.Println("接收服务端给的命令 : ", cmdStr)
		result := cmd.LinuxSendCommand(cmdStr)
		log.Println("执行命令后返回 : ", result)
		conn.Write(Str2Bytes(result, pk.RECEPTION_SLVE_PACKET))
		return

	//接收文件
	case pk.SEND_FILE_PACKET:

		//log.Println("接收文件数据包 : ", packet.PacketContent)
		//解包
		var filePacket structs.SendFilePacket
		json.Unmarshal(packet.PacketContent, &filePacket)
		log.Println("收到文件包 = ", filePacket.FileName, filePacket.MaxPacketNum, filePacket.FilePacketNum)

		//保存文件包
		//加上锁保证保存每一包都成功
		lock.Lock()
		global.FilePackets = append(global.FilePackets, &filePacket)
		lock.Unlock()

		//一个发送完成的标识符号
		if filePacket.IsEnd {
			global.FileEnd = true
		}

		log.Println("FilePackets len = ", int64(len(global.FilePackets)), filePacket.MaxPacketNum)
		//当包数相同与发送完成标识后就进行保存
		if int64(len(global.FilePackets)) == filePacket.MaxPacketNum && global.FileEnd {
			//保存的包按照包id排序
			sort.Slice(global.FilePackets, func(i, j int) bool {
				return global.FilePackets[i].FilePacketNum < global.FilePackets[j].FilePacketNum
			})
			//log.Println("接收完成！", FilePackets)

			//创建文件
			//配置读取文件夹
			f, err := os.Create(global.SlveSpace + filePacket.FileName)
			if err != nil {
				log.Println(err)
				return
			}
			for _, v := range global.FilePackets {
				log.Println("写入文件包 : ", *v)
				f.Write(v.FilePacket)
			}
			f.Close()

			//清空 FilePackets
			global.FilePackets = global.FilePackets[:0:0]
			//FilePacketMap = make([]*structs.SendFilePacket, 0)

			global.FileEnd = false

			//返回一个完成
			conn.Write(Str2Bytes("文件接收完成", pk.SEND_FILE_COMPLETE_PACKET))

		}
		return

	//如果是文件传输结束标志则清空 FilePackets
	case pk.SEND_FILE_COMPLETE_PACKET:
		global.FilePackets = global.FilePackets[:0:0]
		return

	//请求docker 信息相关
	case pk.Docker_Infos:
		data, err := tcpfunc.GetInfo()
		packetData := &structs.DockerAction{
			Action: "获取docker infos",
			Packet: data,
			Error:  err,
		}
		SendPackat(conn, packetData, pk.Docker_Infos)
		return

	//请求docker images相关
	case pk.Docker_Images:
		var dockerImagesPacket structs.DockerImagesAction
		json.Unmarshal(packet.PacketContent, &dockerImagesPacket)
		log.Println("收到Action = ", dockerImagesPacket.Action)
		data, err := tcpfunc.Images(dockerImagesPacket.Action)
		packetData := &structs.DockerImagesAction{
			Action: dockerImagesPacket.Action,
			Packet: data,
			Error:  err,
		}
		SendPackat(conn, packetData, pk.Docker_Images)
		return

	//请求docker container相关
	case pk.Docker_Container:
		var dockerContainerpacket structs.DockerContainerAction
		json.Unmarshal(packet.PacketContent, &dockerContainerpacket)
		log.Println("收到Action = ", dockerContainerpacket.Action)
		data, err := tcpfunc.Container(dockerContainerpacket.Action)
		packetData := &structs.DockerContainerAction{
			Action: dockerContainerpacket.Action,
			Packet: data,
			Error:  err,
		}
		SendPackat(conn, packetData, pk.Docker_Images)
		return
	}

}

func Str2Bytes(s string, packetType byte) []byte {
	packetBytes, err := json.Marshal(s)
	if err != nil {
		log.Println(err.Error())
	}
	packet := structs.Packet{
		PacketType:    packetType,
		PacketContent: packetBytes,
	}
	sendBytes, err := json.Marshal(packet)
	if err != nil {
		log.Println(err.Error())
	}
	return protocol.EnPackSendData(sendBytes)
}

func SendPackat(c net.Conn, s interface{}, packetType byte) {
	packetBytes, err := json.Marshal(s)
	if err != nil {
		log.Println(err.Error())
	}
	packet := structs.Packet{
		PacketType:    packetType,
		PacketContent: packetBytes,
	}
	sendBytes, err := json.Marshal(packet)
	if err != nil {
		log.Println(err.Error())
	}
	c.Write(protocol.EnPackSendData(sendBytes))
}

//发送心跳包，与发送数据包一样
//心跳包包含了当前slve的 资源使用情况
func SendHeartPacket(client *structs.TcpClient) {

	heartPacket := structs.HeartPacket{
		Version:   global.SlveVersion,
		SlveId:    global.SlveUUID,
		IP:        sys2go.GetMyIP(),
		System:    sys2go.GetSysType(),
		HostName:  sys2go.GetHostName(),
		Timestamp: time.Now().Unix(),
	}
	//采集性能
	performance := tcpfunc.GetPerformance()
	heartPacket.UseCPU = performance.CpuRate.UseRate
	heartPacket.UseMEM = performance.MemInfo.MemUsed
	heartPacket.Performance = performance

	log.Println(*heartPacket.Performance)
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
