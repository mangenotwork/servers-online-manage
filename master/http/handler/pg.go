//	页面
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mangenotwork/servers-online-manage/lib/global"
)

//首页
func PGHome (c *gin.Context) {
	slist := make([]string, 0)
	for k, _ := range global.Slves {
		slist = append(slist, k)
	}

	c.JSON(200, gin.H{
		"version": global.Version,
		"slves":   slist,
	})
}

//测试上传文件页面
func PGUploadFileTest(c *gin.Context) {
	 c.HTML(200, "upload_file.html", gin.H{})
	 return
}