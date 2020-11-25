package routers

import (
	"github.com/gin-gonic/gin"
	"time"
)


var staticVer = int(time.Now().Unix())

func PGBaseActive() gin.HandlerFunc{
	return func(c *gin.Context) {

	}
}