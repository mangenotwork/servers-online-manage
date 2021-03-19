// 网络IO
package models

//Slve 网络IO
type NetworkIO struct {
	ID       int64   `gorm:"primary_key;column:id" json:"id"`
	SlveUUID string  `gorm:"column:uuid" json:"uuid"`
	Time     int64   `gorm:"column:time" json:"time"` //采集时间
	Name     string  `gorm:"column:name" json:"name"` //网卡名
	Tx       float32 `gorm:"column:tx" json:"tx"`     //发送
	Rx       float32 `gorm:"column:rx" json:"rx"`     //接收
}

func (s *NetworkIO) TableName() string {
	return "network_io"
}
