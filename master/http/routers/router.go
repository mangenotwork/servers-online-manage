package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mangenotwork/servers-online-manage/master/http/handler"
)

var Router *gin.Engine

func Routers() *gin.Engine {

	Router = gin.Default()
	Router.Delims("<<", ">>")
	//静态目录配置
	// Router.Static("/static", "static")
	// Router.Static("/install/static", "static")
	Router.StaticFS("/static", http.Dir("./static"))

	//模板
	//Router.LoadHTMLGlob("templates/**/*") //多级目录
	Router.LoadHTMLGlob("templates/*") //单级目录

	//页面
	{
		Router.GET("/", handler.PGHome)                     //首页
		Router.GET("/sendfilepg", handler.PGUploadFileTest) //上传文件页面
		Router.GET("/home", handler.PGHome)
		Router.GET("/host", handler.PGHostList)                  // 服务器
		Router.GET("/property", handler.PGProperty)              // 资产
		Router.GET("/release", handler.PGRelease)                //部署
		Router.GET("/alarm", handler.PGAlarm)                    //警报与通知
		Router.GET("/settings", handler.PGSettings)              //设置
		Router.GET("/user", handler.PGUserManage)                //账号管理
		Router.GET("/help", handler.PGHelp)                      //帮助
		Router.Any("/SlveFile/*URL", handler.ZF)                 //转发测试
		Router.GET("/slve/details/:slveId", handler.SlveDetails) //Slve详细信息与管理
		Router.GET("/slve/docker/:slveId", handler.SlveDocker)   //Slve docker管理
	}

	//测试用的
	{
		Router.GET("/getinfo", handler.GetInfoTets)          //获取slve host 信息
		Router.GET("/send", handler.SendCMDTest)             ////send cmd
		Router.POST("/uploadfiles", handler.UploadfilesTest) //接收上传文件
		Router.GET("/docker/images", handler.DockerImagesTest)
	}

	//接口
	API := Router.Group("/api")
	{
		//slve 相关的接口
		SlveAPI := API.Group("/slve/v1")
		{
			SlveAPI.GET("/ip/list", handler.GetSlveIPList)                 //获取当前连接 slve IP
			SlveAPI.GET("/list", handler.GetSlveList)                      //slve 列表, 包含了基本信息
			SlveAPI.GET("/echart/base/:slveId", handler.GetEchartBaseData) //获取Slve性能基础图表数据
			SlveAPI.GET("/pidlist/:slveId", handler.GetPIDList)            //获取Slve 进程列表
			SlveAPI.GET("/envlist/:slveId", handler.GetENVList)            //获取Slve 环境变量
		}

		//警报与通知 	alarm  相关接口
		AlarmAPI := API.Group("/alarm/v1")
		{
			AlarmAPI.GET("/list", handler.GetAlarmList) //消息 列表
		}

		//slve 信息相关的接口
		SlveInfoAPI_v1 := API.Group("/slve/info/v1")
		{
			SlveInfoAPI_v1.GET("/ping/:slveId", handler.DockerImagesTest)
		}

		//docker 相关的接口
		DockerAPI_V1 := API.Group("/docker/v1")
		{
			DockerAPI_V1.GET("/images", handler.DockerImagesTest)                    //docker images
			DockerAPI_V1.GET("/infos/:slveId", handler.DockerInfos)                  //获取 selve docker 基本信息
			DockerAPI_V1.GET("/images/list/:slveId", handler.DockerImagesList)       //获取slve images 列表
			DockerAPI_V1.GET("/container/list/:slveId", handler.DockerContainerList) //获取slve container  列表
		}
	}

	//404
	Router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", "")
	})

	//401
	Router.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusUnauthorized, "未授权的访问")
	})

	//403
	Router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusForbidden, "403.html", "")
	})

	//500
	Router.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusInternalServerError, "500")
	})
	return Router
}
