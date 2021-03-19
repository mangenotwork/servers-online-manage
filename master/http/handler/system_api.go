package handler

import (
	"fmt"
	"log"

	"github.com/mangenotwork/servers-online-manage/lib/loger"

	"github.com/gin-gonic/gin"
	"github.com/mangenotwork/servers-online-manage/lib/global"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
	"github.com/mangenotwork/servers-online-manage/master/http/dao"
	"github.com/mangenotwork/servers-online-manage/master/tcp"
)

//获取slve ip 列表
func GetSlveIPList(c *gin.Context) {
	slist := make([]string, 0)
	for k, _ := range global.Slves {
		slist = append(slist, k)
	}

	c.JSON(200, gin.H{
		"version": global.Version,
		"slves":   slist,
	})
}

//获取当前连接的slve 包含了基础信息
func GetSlveList(c *gin.Context) {

	data := make([]*structs.SlveBaseInfo, 0)

	for k, v := range global.Slves {
		log.Println("已经连接的客户端: ", k, &v.SlveInfo.Name)
		data = append(data, v.SlveInfo)
	}

	slveDao := new(dao.SlveBaseInfoDao)
	err := slveDao.Gets()
	slveDao.IsOnlines()
	loger.Debug(err)
	loger.Debug(slveDao.Datas)
	c.JSON(200, gin.H{
		"version": global.Version,
		"slves":   slveDao.Datas,
	})
}

//获取slve host 信息
func GetInfoTets(c *gin.Context) {
	slve := c.Query("slve")

	//获取tcp连接对象
	slveConn := global.Slves[slve]
	if slveConn == nil {
		c.JSON(200, gin.H{
			"data": fmt.Sprintf("%s 连接丢失", slve),
		})
		return
	}

	log.Println("slveConn = ", slveConn)

	//执行tcp方法
	tcp.GetSlveInfo(slveConn.Conn)
	//接收slv返回值
	data := <-slveConn.Rdata
	c.JSON(200, gin.H{
		"data": data,
	})
	return
}

// 使Slve执行命令
// BUG: ./ 这类命令会阻塞 slveConn.Rdata
func SendCMDTest(c *gin.Context) {
	slve := c.Query("slve")
	cmd := c.Query("cmd")
	slveConn := global.Slves[slve]
	if slveConn == nil {
		c.JSON(200, gin.H{
			"data": fmt.Sprintf("%s 连接丢失", slve),
		})
		return
	}
	tcp.SendSlveCmd(slveConn.Conn, cmd)
	data := <-slveConn.Rdata
	c.JSON(200, gin.H{
		"data": data,
	})

	// select {
	// case data := <-slveConn.Rdata:
	// 	c.IndentedJSON(200, gin.H{
	// 		"data": string(data),
	// 	})
	// }
	return
}

//测试上传文件
func UploadfilesTest(c *gin.Context) {
	// 单文件

	slve := c.PostForm("slve")
	log.Println(slve)

	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	f, err := file.Open()
	if err != nil {
		fmt.Println(err)
		c.String(200, "打开文件失败")
		return
	}
	defer f.Close()
	log.Println(f)
	slveConn := global.Slves[slve]
	if slveConn == nil {
		c.JSON(200, gin.H{
			"data": fmt.Sprintf("%s 连接丢失", slve),
		})
		return
	}
	tcp.SendFile2(slveConn.Conn, f, file.Size, file.Filename)

	// 上传文件到指定的路径
	// c.SaveUploadedFile(file, dst)

	c.String(200, fmt.Sprintf("'%s' uploaded!", file.Filename))
	return
}

func DockerImagesTest(c *gin.Context) {
	slve := c.Query("slve")
	log.Println("slve = ", slve)

	slveConn := global.Slves[slve]
	if slveConn == nil {
		c.JSON(200, gin.H{
			"data": fmt.Sprintf("%s 连接丢失", slve),
		})
		return
	}

	tcp.GetDockerImages(slveConn.Conn)
	data := <-slveConn.Rdata
	c.JSON(200, gin.H{
		"data": data,
	})
}

//GetEchartBaseData  获取Slve性能基础图表数据
func GetEchartBaseData(c *gin.Context) {
	uuid := c.Param("slveId")
	//获取cpu使用率 得到时间段
	cpuDao := new(dao.CPURateDao)
	cpuDao.GetListFromTimeMainCPU(uuid, 3)
	timeList, showTime, cpuShowData := cpuDao.EchartData(20)
	loger.Debug("timeList = ", timeList)
	//获取内存使用率
	memDao := new(dao.MEMInfoDao)
	memDao.GetFromTimes(timeList)
	memShowData := memDao.EchartData()
	//获取磁盘IO
	// diskDao := new(dao.DiskInfoDao)
	// diskDao.GetFromTimes(timeList)
	// diskShowData := diskDao.EchartData()
	//获取网络
	networkDao := new(dao.NetworkIODao)
	networkDao.GetFromTimes(timeList)
	txShowData, rxShowData := networkDao.EchartData()

	c.JSON(200, gin.H{
		"showTime":    showTime,
		"cpuShowData": cpuShowData,
		"memShowData": memShowData,
		"txShowData":  txShowData,
		"rxShowData":  rxShowData,
	})
}
