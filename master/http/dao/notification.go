package dao

import (
	"log"

	"github.com/mangenotwork/servers-online-manage/lib/dbconn"
	"github.com/mangenotwork/servers-online-manage/master/http/models"
)

type NotificationDao struct {
	SlveName  string //slvename
	Notiftype int64  //消息类型
	State     int64  //消息状态
	StartTime int64  //起始时间
	EndTime   int64  //结束时间
	Page      int64  //页数
	PageSize  int64  //页数
}

//未读消息数量
func (this *NotificationDao) PendingCount() (count int, err error) {
	db := dbconn.Conn()
	defer db.Close()
	err = db.Model(&models.Notifincation{}).Where("state = 0").Count(&count).Error
	return
}

//获取消息
func (this *NotificationDao) Get() (datas []*models.Notifincation, count int, err error) {
	db := dbconn.Conn()
	defer db.Close()
	db = db.Model(&models.Notifincation{})
	if this.SlveName != "" {
		db = db.Where("slve = ?", this.SlveName)
	}
	if this.Notiftype > 0 {
		db = db.Where("type = ?", this.Notiftype)
	}
	if this.State > 0 {
		db = db.Where("state = ?", this.State)
	}
	if this.StartTime > 0 {
		db = db.Where("time >= ?", this.StartTime)
	}
	if this.EndTime > 0 {
		db = db.Where("time >= ?", this.EndTime)
	}
	log.Println("this.Page = ", this.Page, this.PageSize)
	limit := this.Page * this.PageSize
	if this.Page > 0 {
		this.Page = this.Page - 1
	}
	offset := this.Page * this.PageSize
	log.Println("offset = ", offset)
	err = db.Limit(limit).Offset(offset).Find(&datas).Error
	err = db.Count(&count).Error
	return
}

//TODO 警报标记为已读
