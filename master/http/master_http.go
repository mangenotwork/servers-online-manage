package http

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mangenotwork/servers-online-manage/lib/global"
	"github.com/mangenotwork/servers-online-manage/master/tcp"
)

// gin http servers
func Httpserver() {
	r := gin.Default()

	//模板
	r.LoadHTMLGlob("templates/*")

	//首页
	r.GET("/", func(c *gin.Context) {
		slist := make([]string, 0)
		for k, _ := range global.Slves {
			slist = append(slist, k)
		}

		c.JSON(200, gin.H{
			"version": global.Version,
			"slves":   slist,
		})
	})

	//获取slve host 信息
	r.GET("/getinfo", func(c *gin.Context) {
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
	})

	//send cmd
	//BUG: ./ 这类命令会阻塞 slveConn.Rdata
	r.GET("/send", func(c *gin.Context) {
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
	})

	//上传文件页面
	r.GET("/sendfilepg", func(c *gin.Context) {
		c.HTML(200, "upload_file.html", gin.H{})
		return
	})

	//接收上传文件
	r.POST("/uploadfiles", func(c *gin.Context) {
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
	})

	r.Run(":15555")
}
