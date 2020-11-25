//通知数据
//警报与通知业务

package models

type Notifincation struct {
	ID int64 `gorm:"primary_key;column:id" json:"id"`
	Type int  `gorm:"column:type" json:"type"`//通知类型   消息，完成，失败，异常，警报
	State int `gorm:"column:state" json:"state"`//通知状态   未读，已读
	Messg string `gorm:"column:messg" json:"messg"`//通知内容
	Time int64 `gorm:"column:time" json:"time"`//通知时间
}

func (s *Notifincation) TableName() string {
	return "notifincation"
}
