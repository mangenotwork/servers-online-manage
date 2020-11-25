package tcp

import (
	"encoding/json"
	"github.com/mangenotwork/servers-online-manage/lib/global"
	"github.com/mangenotwork/servers-online-manage/utils"
	"log"
	"strings"
	"time"

	pk "github.com/mangenotwork/servers-online-manage/lib/packet"
	"github.com/mangenotwork/servers-online-manage/structs"
)

//Master接收的具体业务
//conn: 客户端的连接实例
//packet: 接收客户端的包
func MasterTcpFunc(conn *structs.Cli, packet *structs.Packet) {
	//根据自定义的包类型进行不同业务的处理
	switch packet.PacketType {

	//Slve连接后的第一包
	case pk.FIRST_PACKET:
		//解析数据
		var beatPacket structs.SlveBaseInfo
		json.Unmarshal(packet.PacketContent, &beatPacket)
		log.Println("beatPacket = ", &beatPacket)
		slveKey := strings.Split(conn.Conn.RemoteAddr().String(),":")[0]
		//获取slve， global.Slves[slveKey]
		beatPacket.SlveKey = slveKey
		beatPacket.HostIP = conn.Conn.RemoteAddr().String()
		beatPacket.ConnTime = utils.DateTime(time.Now().Unix())
		global.Slves[slveKey].SlveInfo = &beatPacket

	//处理心跳
	case pk.HEART_BEAT_PACKET:
		//解析数据
		var beatPacket structs.HeartPacket
		json.Unmarshal(packet.PacketContent, &beatPacket)
		log.Printf("收到心跳数据 [%s] ,data is [%v]\n", conn.Conn.RemoteAddr().String(), beatPacket)

		//收到返回
		packet := structs.Packet{
			PacketType:    pk.REPLY_HEART_PACKET,
			PacketContent: []byte("成功收到心跳包！"),
		}
		SendData(conn.Conn, packet)
		return

	//处理数据包
	case pk.REPORT_PACKET:
		var reportPacket structs.ReportPacket
		json.Unmarshal(packet.PacketContent, &reportPacket)
		log.Printf("recieve report data from [%s] ,data is [%v]\n", conn.Conn.RemoteAddr().String(), reportPacket)
		conn.Conn.Write([]byte("Report data has recive\n"))
		return

	//接收slve数据
	case pk.RECEPTION_SLVE_PACKET:
		log.Println("接收slve数据 = ", string(packet.PacketContent))
		//将数据发送给chan
		conn.Rdata <- string(packet.PacketContent)
		return

	//slve接收文件成功
	case pk.SEND_FILE_COMPLETE_PACKET:
		log.Println("发送文件成功！")
		return

		//接收slve返回的docker images 相关操作（方法）的包
	case pk.Docker_Images:
		var callbackPacket structs.DockerImagesAction
		json.Unmarshal(packet.PacketContent, &callbackPacket)
		//将数据发送给chan
		conn.Rdata <- string(callbackPacket.Packet)
		return
	}


}
