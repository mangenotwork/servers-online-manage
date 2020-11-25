package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/mangenotwork/servers-online-manage/master/http/handler"
)

var Router *gin.Engine

func Routers() *gin.Engine {

	Router = gin.Default()

	//静态目录配置
	// Router.Static("/static", "static")
	// Router.Static("/install/static", "static")
	Router.StaticFS("/static", http.Dir("./static"))

	//模板
	//Router.LoadHTMLGlob("templates/**/*") //多级目录
	Router.LoadHTMLGlob("templates/*") //单级目录

	//页面
	{
		Router.GET("/", handler.PGHome) //首页
		Router.GET("/sendfilepg",handler.PGUploadFileTest)//上传文件页面
	}

	//测试用的
	{
		Router.GET("/getinfo", handler.GetInfoTets)//获取slve host 信息
		Router.GET("/send",handler.SendCMDTest)////send cmd
		Router.POST("/uploadfiles",handler.UploadfilesTest)//接收上传文件
		Router.GET("/docker/images",handler.DockerImagesTest)
	}

	//接口
	API := Router.Group("/api")
	{
		//slve 相关的接口
		SlveAPI_V1 := API.Group("/slve/v1")
		{
			SlveAPI_V1.GET("/ip/list", handler.GetSlveIPList)//获取当前连接 slve IP
			SlveAPI_V1.GET("/list", handler.GetSlveList)//slve 列表, 包含了基本信息
		}

		//slve 信息相关的接口
		SlveInfoAPI_v1 := API.Group("/slve/info/v1")
		{
			SlveInfoAPI_v1.GET("/ping/:slveId", handler.DockerImagesTest)
		}

		//docker 相关的接口
		DockerAPI_V1 :=  API.Group("/docker/v1")
		{
			DockerAPI_V1.GET("/images", handler.DockerImagesTest) //docker images
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
		ctx.String(http.StatusInternalServerError, "500", "")
	})
	return Router
}