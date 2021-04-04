//tcp交互的实现, 提供给http调用
package tcp

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net"
	"time"

	"github.com/mangenotwork/servers-online-manage/lib/global"
	pk "github.com/mangenotwork/servers-online-manage/lib/packet"
	"github.com/mangenotwork/servers-online-manage/lib/protocol"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
)

//向客户端发送数据
func SendData(c net.Conn, packet structs.Packet) {
	sendBytes, err := json.Marshal(packet)
	if err != nil {
		log.Println(err.Error())
	}
	c.Write(protocol.EnPackSendData(sendBytes))
}

//向客户端发送数据包
func SendPacket(c net.Conn, s interface{}, pactetType byte) {
	packetBytes, err := json.Marshal(s)
	if err != nil {
		log.Println(err.Error())
	}
	packet := structs.Packet{
		PacketType:    pactetType,
		PacketContent: packetBytes,
	}
	SendData(c, packet)
}

//请求slve的host info
func GetSlveInfo(conn net.Conn) {
	log.Println("请求slve的host info")
	//发送请求
	packet := structs.Packet{
		PacketType:    pk.Get_SLVE_INFO_PACKET,
		PacketContent: []byte(""),
	}
	SendData(conn, packet)
}

//向 slve发送命令，slve收到命令并执行
func SendSlveCmd(conn net.Conn, cmd string) {
	packet := structs.Packet{
		PacketType:    pk.SET_SLVE_CMD_PACKET,
		PacketContent: []byte(cmd),
	}
	SendData(conn, packet)
}

//向 slve发送文件
//有bug的方案
func SendFile(conn net.Conn, f multipart.File, fileSize int64, filename string) {
	//最大包数量
	maxpacket := fileSize / global.FilePacketSize
	if fileSize%global.FilePacketSize > 0 {
		maxpacket += 2
	}
	log.Println("maxpacket = ", maxpacket)
	filePacketNum := 1
	var count int64
	for {
		log.Println("第 ", filePacketNum, " 包")
		filepacket := structs.SendFilePacket{
			FileName:      filename,
			FileSuffix:    "",
			FilePacketNum: filePacketNum,
			MaxPacketNum:  maxpacket,
		}
		buf := make([]byte, global.FilePacketSize)
		//读取文件内容
		n, err := f.Read(buf)
		if err != nil && io.EOF == err {
			log.Println("文件传输完成 ： ", err)
			filepacket.IsEnd = true
			packetBytes, err := json.Marshal(filepacket)
			if err != nil {
				log.Println(err.Error())
			}
			packet := structs.Packet{
				PacketType:    pk.SEND_FILE_PACKET,
				PacketContent: packetBytes,
			}
			SendData(conn, packet)
			return
		}
		filepacket.FilePacket = buf[:n]
		filepacket.IsEnd = false
		packetBytes, err := json.Marshal(filepacket)
		if err != nil {
			log.Println(err.Error())
		}
		packet := structs.Packet{
			PacketType:    pk.SEND_FILE_PACKET,
			PacketContent: packetBytes,
		}
		SendData(conn, packet)
		filePacketNum++
		count += int64(n)
		sendPercent := float64(count) / float64(fileSize) * 100
		value := fmt.Sprintf("%.2f", sendPercent)
		//打印上传进度
		log.Println("文件上传：" + value + "%")
	}
}

//确定方案，文档方案
//向 slve发送文件 2
func SendFile2(conn net.Conn, f multipart.File, fileSize int64, filename string) {
	log.Println("conn = ", conn)
	log.Println("f = ", f)
	log.Println("fileSize = ", fileSize)

	//包的最大数量
	maxpacket := 1
	//当前包计数
	filePacketNum := 1
	//用于保存分包
	sendList := make([]*structs.SendFilePacket, 0)

	//分包
	for {
		log.Println("第 ", filePacketNum, " 包")
		filepacket := &structs.SendFilePacket{
			FileName:      filename,
			FileSuffix:    "",
			FilePacketNum: filePacketNum,
		}
		buf := make([]byte, global.FilePacketSize)
		//读取文件内容
		n, err := f.Read(buf)
		if err != nil && io.EOF == err {
			log.Println("文件传输完成 ： ", err)
			filepacket.IsEnd = true
			sendList = append(sendList, filepacket)
			break
		}
		filepacket.FilePacket = buf[:n]
		filepacket.IsEnd = false
		sendList = append(sendList, filepacket)
		filePacketNum++
		maxpacket++
	}

	//发包
	for k, v := range sendList {
		v.MaxPacketNum = int64(maxpacket)
		packetBytes, err := json.Marshal(v)
		if err != nil {
			log.Println(err.Error())
		}
		packet := structs.Packet{
			PacketType:    pk.SEND_FILE_PACKET,
			PacketContent: packetBytes,
		}
		SendData(conn, packet)

		//打印一个百分比
		sendPercent := float64(k) / float64(maxpacket) * 100
		value := fmt.Sprintf("%.2f", sendPercent)
		//打印上传进度
		log.Println("文件上传：" + value + "%")
	}

	//等待2秒发送一个结束的标识
	time.Sleep(2 * time.Second)
	packet := structs.Packet{
		PacketType: pk.SEND_FILE_COMPLETE_PACKET,
	}
	SendData(conn, packet)
}

//获取Slve PID list
func GetSlvePIDList(conn net.Conn) {
	packet := structs.Packet{
		PacketType:    pk.GET_SLVE_PID_LIST,
		PacketContent: []byte(""),
	}
	SendData(conn, packet)
}

//获取Slve 环境变量
func GETSlveENV(conn net.Conn) {
	packet := structs.Packet{
		PacketType:    pk.GET_SLVE_ENV,
		PacketContent: []byte(""),
	}
	SendData(conn, packet)
}

//获取Slve Docker 镜像列表
func GetDockerImages(conn net.Conn) {
	active := structs.DockerImagesAction{
		Action: "get_images_list",
	}
	SendPacket(conn, active, pk.Docker_Images)
}

//获取 Slve Docker 基本信息
func GetDockerInfos(conn net.Conn) {
	SendPacket(conn, "", pk.Docker_Infos)
}

//获取 Slve Docker 镜像列表
func DockerImagesList(conn net.Conn) {
	active := structs.DockerImagesAction{
		Action: "list",
	}
	SendPacket(conn, active, pk.Docker_Images)
}

//获取 SLve Docker 容器列表
func DockerContainerList(conn net.Conn) {
	active := structs.DockerContainerAction{
		Action: "list",
	}
	SendPacket(conn, active, pk.Docker_Container)
}
