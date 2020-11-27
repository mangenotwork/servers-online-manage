package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mangenotwork/servers-online-manage/master/http/routers"
)

// gin http servers
func Httpserver() {
	gin.SetMode(gin.DebugMode)
	s := routers.Routers()
	port := "15555"
	s.Run(":"+port)
}
