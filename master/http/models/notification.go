//通知数据
//警报与通知业务

package models

import (
	"github.com/mangenotwork/servers-online-manage/lib/dbconn"
)

type Notifincation struct {
	ID    int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Slve  string `gorm:"column:slve" json:"slve"`   //slve的描述， 是ip或者是名称
	Type  int    `gorm:"column:type" json:"type"`   //通知类型   消息通知: 1，完成:2，失败:3，异常:4，警报:5
	State int    `gorm:"column:state" json:"state"` //通知状态   未读1，已读2
	Messg string `gorm:"column:messg" json:"messg"` //通知内容
	Time  int64  `gorm:"column:time" json:"time"`   //通知时间
}

func (this *Notifincation) TableName() string {
	return "notifincation"
}

//创建消息
func (this *Notifincation) Create() (err error) {
	db := dbconn.Conn()
	defer db.Close()
	err = db.Create(&this).Error
	return
}
