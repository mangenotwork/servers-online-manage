//	页面
package handler

import (
	"github.com/gin-gonic/gin"
	//"github.com/mangenotwork/servers-online-manage/lib/global"
)

//首页
func PGHome (c *gin.Context) {
	// slist := make([]string, 0)
	// for k, _ := range global.Slves {
	// 	slist = append(slist, k)
	// }

	// c.JSON(200, gin.H{
	// 	"version": global.Version,
	// 	"slves":   slist,
	// })
	c.HTML(200, "home.html", gin.H{})
	return
}

// 服务器
func PGHostList(c *gin.Context) {
	c.HTML(200, "host_list.html", gin.H{})
	return
}

// 资产
func PGProperty(c *gin.Context) {
	c.HTML(200, "property.html", gin.H{})
	return
}

//部署
func PGRelease(c *gin.Context) {
	c.HTML(200, "release.html", gin.H{})
	return
}

//警报与通知
func PGAlarm(c *gin.Context) {
	c.HTML(200, "alarm.html", gin.H{})
	return
}

//设置
func  PGSettings(c *gin.Context) {
	c.HTML(200, "settings.html", gin.H{})
	return
}

//账号管理
func PGUserManage(c *gin.Context) {
	c.HTML(200, "user_manage.html", gin.H{})
	return
}

//帮助
func PGHelp(c *gin.Context) {
	c.HTML(200, "help.html", gin.H{})
	return
}

//测试上传文件页面
func PGUploadFileTest(c *gin.Context) {
	 c.HTML(200, "upload_file.html", gin.H{})
	 return
}