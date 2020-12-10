package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mangenotwork/servers-online-manage/lib/enum"
	"github.com/mangenotwork/servers-online-manage/lib/global"
	"github.com/mangenotwork/servers-online-manage/lib/utils"
	"github.com/mangenotwork/servers-online-manage/master/http/dao"
)

func GetAlarmList(c *gin.Context) {

	slve := c.Query("slve")
	ntype := utils.Str2Int64(c.Query("ntype"))
	state := utils.Str2Int64(c.Query("state"))
	start := utils.Str2Int64(c.Query("state"))
	end := utils.Str2Int64(c.Query("end"))
	pg := utils.Str2Int64(c.Query("page"))
	//pgSize := utils.Str2Int64(c.Query("page_size"))

	if pg <= 0 {
		pg = 1
	}

	d := &dao.NotificationDao{
		SlveName:  slve,
		Notiftype: ntype,
		State:     state,
		StartTime: start,
		EndTime:   end,
		Page:      pg,
		PageSize:  2,
	}
	datas, count, err := d.Get()
	if err != nil {
		c.JSON(200, gin.H{
			"version": global.Version,
			"message": err.Error(),
			"code":    enum.Error,
		})
	}
	log.Println(datas)
	log.Println("len = ", count)
	log.Println("err  = ", err)

	c.JSON(200, gin.H{
		"version": global.Version,
		"data":    datas,
		"code":    enum.Success,
	})
}
