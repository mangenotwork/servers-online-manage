package slve

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"sort"
	"sync"

	"github.com/mangenotwork/servers-online-manage/lib/cmd"
	"github.com/mangenotwork/servers-online-manage/lib/global"
	pk "github.com/mangenotwork/servers-online-manage/lib/packet"
	"github.com/mangenotwork/servers-online-manage/lib/protocol"
	"github.com/mangenotwork/servers-online-manage/structs"
)

var lock = &sync.Mutex{}

//具体执行的业务
func SlveTcpFunc(conn net.Conn, packet *structs.Packet) {
	switch packet.PacketType {

	case pk.SET_SLVE_TOKEN_PACKET:
		//接收服务端颁发的名称
		log.Println("接收服务端颁发的名称 : ", string(packet.PacketContent))
		global.SlveToken = string(packet.PacketContent)
		return

	case pk.REPLY_HEART_PACKET:
		log.Println("接收心跳包的回复 : ", string(packet.PacketContent))
		return

	//返回slve信息给master
	case pk.Get_SLVE_INFO_PACKET:
		log.Println("返回slve信息给master ")
		hostinfo := GetHostInfo()
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

		//请求docker images相关
	case pk.Docker_Images:
		var dockerImagesPacket structs.DockerImagesAction
		json.Unmarshal(packet.PacketContent, &dockerImagesPacket)
		log.Println("收到Action = ", dockerImagesPacket.Action)
		data := Images(dockerImagesPacket.Action)
		packetData := &structs.DockerImagesAction{
			Action: dockerImagesPacket.Action,
			Packet: data,
		}
		SendPackat(conn,packetData,pk.Docker_Images)
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

func SendPackat(c net.Conn,s interface{},packetType byte) {
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