//	页面
package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mangenotwork/servers-online-manage/lib/global"
	"github.com/mangenotwork/servers-online-manage/master/http/dao"
)

//首页
func PGHome(c *gin.Context) {

	//获取 host 连接的个数
	connHostCount := global.SlveLen()
	//获取 资产个数

	//获取 警报与通知个数
	notifincation, err := new(dao.NotificationDao).PendingCount()
	if err != nil {
		notifincation = 0
	}
	log.Println("notifincation = ", notifincation)

	c.HTML(200, "home.html", gin.H{
		"conn_count":          connHostCount,
		"notifincation_count": notifincation,
	})
	return
}

// 服务器
func PGHostList(c *gin.Context) {

	//获取服务列表

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
func PGSettings(c *gin.Context) {
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
