package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mangenotwork/servers-online-manage/lib/global"
	"github.com/mangenotwork/servers-online-manage/lib/loger"
	"github.com/mangenotwork/servers-online-manage/master/tcp"
)

//获取 selve docker 基本信息
func DockerInfos(c *gin.Context) {

	slve := c.Param("slveId")
	loger.Debug("slve = ", slve)

	slveConn := global.Slves[slve]
	if slveConn == nil {
		c.JSON(200, gin.H{
			"data": fmt.Sprintf("%s 连接丢失", slve),
		})
		return
	}

	tcp.GetDockerInfos(slveConn.Conn)
	data := <-slveConn.Rdata
	c.JSON(200, gin.H{
		"data": data,
	})
}

//获取slve images 列表
func DockerImagesList(c *gin.Context) {
	slve := c.Param("slveId")
	loger.Debug("slve = ", slve)
	slveConn := global.Slves[slve]
	if slveConn == nil {
		c.JSON(200, gin.H{
			"data": fmt.Sprintf("%s 连接丢失", slve),
		})
		return
	}
	tcp.DockerImagesList(slveConn.Conn)
	data := <-slveConn.Rdata
	c.JSON(200, gin.H{
		"data": data,
	})
}

//获取slve container  列表
func DockerContainerList(c *gin.Context) {
	slve := c.Param("slveId")
	loger.Debug("slve = ", slve)
	slveConn := global.Slves[slve]
	if slveConn == nil {
		c.JSON(200, gin.H{
			"data": fmt.Sprintf("%s 连接丢失", slve),
		})
		return
	}
	tcp.DockerContainerList(slveConn.Conn)
	data := <-slveConn.Rdata
	c.JSON(200, gin.H{
		"data": data,
	})
}
