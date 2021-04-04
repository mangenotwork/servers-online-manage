package tcp

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/mangenotwork/servers-online-manage/lib/global"
	"github.com/mangenotwork/servers-online-manage/lib/utils"

	"github.com/mangenotwork/servers-online-manage/lib/loger"
	pk "github.com/mangenotwork/servers-online-manage/lib/packet"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
	"github.com/mangenotwork/servers-online-manage/master/http/dao"
	"github.com/mangenotwork/servers-online-manage/master/http/models"
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
		slveKey := strings.Split(conn.Conn.RemoteAddr().String(), ":")[0]
		//获取slve， global.Slves[slveKey]
		beatPacket.SlveKey = slveKey
		beatPacket.HostIP = conn.Conn.RemoteAddr().String()
		beatPacket.ConnTime = utils.DateTime(time.Now().Unix())
		global.Slves[slveKey].SlveInfo = &beatPacket
		//TODO 保存基础信息，如果数据库存在更新信息
		slveBaseDao := new(dao.SlveBaseInfoDao)
		slveBaseDao.Data = &models.SlveBaseInfo{
			SlveUUID:        beatPacket.SlveUUID,
			Name:            beatPacket.Name,
			SetName:         beatPacket.SetName,
			HostIP:          beatPacket.HostIP,
			SysType:         beatPacket.SysType,
			SlveVersion:     beatPacket.SlveVersion,
			LastConn:        beatPacket.ConnTime,
			OsName:          beatPacket.SysInfo.OsName,
			SysArchitecture: beatPacket.SysInfo.SysArchitecture,
			CpuCoreNumber:   beatPacket.SysInfo.CpuCoreNumber,
			CpuName:         beatPacket.SysInfo.CpuName,
			CpuID:           beatPacket.SysInfo.CpuID,
			BaseBoardID:     beatPacket.SysInfo.BaseBoardID,
			MemTotal:        beatPacket.SysInfo.MemTotal,
			DiskTotal:       beatPacket.SysInfo.DiskTotal,
		}
		if !slveBaseDao.IsHave(beatPacket.SlveUUID) {
			slveBaseDao.Create()
		} else {
			slveBaseDao.Update()
		}

	//处理心跳
	case pk.HEART_BEAT_PACKET:
		//解析数据
		var beatPacket structs.HeartPacket
		json.Unmarshal(packet.PacketContent, &beatPacket)
		log.Printf("收到心跳数据 [%s] ,data is [%v]\n", conn.Conn.RemoteAddr().String(), beatPacket, *beatPacket.Performance)

		//处理数据，并保存

		//1.存cpu
		cpuDao := new(dao.CPURateDao)
		cpuDao.Datas = make([]*models.CPURate, 0)
		cpuDao.Datas = append(cpuDao.Datas, &models.CPURate{
			SlveUUID: beatPacket.SlveId,
			Time:     beatPacket.Timestamp,
			IsMain:   1,
			CPU:      beatPacket.Performance.CpuRate.CPU,
			UseRate:  beatPacket.Performance.CpuRate.UseRate,
		})
		//cpuDao.Create()
		for _, v := range beatPacket.Performance.CpucoreRate {
			cpuDao.Datas = append(cpuDao.Datas, &models.CPURate{
				SlveUUID: beatPacket.SlveId,
				Time:     beatPacket.Timestamp,
				IsMain:   0,
				CPU:      v.CPU,
				UseRate:  v.UseRate,
			})
		}
		cpuDao.Creates()

		//2.存磁盘
		diskDao := new(dao.DiskInfoDao)
		diskDao.Datas = make([]*models.DiskInfo, 0)
		for _, v := range beatPacket.Performance.DiskInfo {
			diskDao.Datas = append(diskDao.Datas, &models.DiskInfo{
				SlveUUID:    beatPacket.SlveId,
				Time:        beatPacket.Timestamp,
				DiskName:    v.DiskName,
				DistType:    v.DistType,
				DistTotalMB: v.DistTotalMB,
				Total:       v.DistUse.Total,
				Free:        v.DistUse.Free,
				Rate:        v.DistUse.Rate,
			})
		}
		diskDao.Creates()

		//3.存网络
		networkDao := new(dao.NetworkIODao)
		networkDao.Datas = make([]*models.NetworkIO, 0)
		for _, v := range beatPacket.Performance.NetworkIO {
			networkDao.Datas = append(networkDao.Datas, &models.NetworkIO{
				SlveUUID: beatPacket.SlveId,
				Time:     beatPacket.Timestamp,
				Name:     v.Name,
				Tx:       v.Tx,
				Rx:       v.Rx,
			})
		}
		networkDao.Creates()

		//4.存内存
		memDao := new(dao.MEMInfoDao)
		used := beatPacket.Performance.MemInfo.MemUsed
		total := beatPacket.Performance.MemInfo.MemTotal
		rate := (float32(used) / float32(total)) * 100
		loger.Debug("mem Rate = ", rate, used, total)
		memDao.Data = &models.MEMInfo{
			SlveUUID: beatPacket.SlveId,
			Time:     beatPacket.Timestamp,
			Total:    beatPacket.Performance.MemInfo.MemTotal,
			Used:     beatPacket.Performance.MemInfo.MemUsed,
			Free:     beatPacket.Performance.MemInfo.MemFree,
			Rate:     rate,
			Buffers:  beatPacket.Performance.MemInfo.MemBuffers,
			Cached:   beatPacket.Performance.MemInfo.MemCached,
		}
		memDao.Create()

		//TODO: 监控点

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

	//接收slve返回的 docker infos 数据
	case pk.Docker_Infos:
		var callbackPacket structs.DockerAction
		json.Unmarshal(packet.PacketContent, &callbackPacket)
		//将数据发送给chan
		conn.Rdata <- string(callbackPacket.Packet)
		return

	//接收slve返回的docker container 相关操作（方法）的包
	case pk.Docker_Container:
		var callbackPacket structs.DockerContainerAction
		json.Unmarshal(packet.PacketContent, &callbackPacket)
		//将数据发送给chan
		conn.Rdata <- string(callbackPacket.Packet)
		return

	//接收slave返回的pid list
	case pk.GET_SLVE_PID_LIST:
		log.Println("接收slve数据 = ", string(packet.PacketContent))
		//将数据发送给chan
		conn.Rdata <- string(packet.PacketContent)
		return

	//接收slave返回的环境变量
	case pk.GET_SLVE_ENV:
		log.Println("接收slve数据 = ", string(packet.PacketContent))
		//将数据发送给chan
		conn.Rdata <- string(packet.PacketContent)
		return

	}

}
