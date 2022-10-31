// 磁盘信息
package models

//Slve 磁盘信息
type DiskInfo struct {
	ID          int64   `gorm:"primary_key;column:id" json:"id"`
	SlveUUID    string  `gorm:"column:uuid" json:"uuid"`
	Time        int64   `gorm:"column:time" json:"time"` //采集时间
	DiskName    string  `gorm:"column:name" json:"name"`
	DistType    string  `gorm:"column:type" json:"type"`
	DistTotalMB string  `gorm:"column:total_mb" json:"total_mb"` // 单位MB
	Total       int     `gorm:"column:total" json:"total"`       // 单位MB
	Free        int     `gorm:"column:free" json:"free"`         // 可用 单位MB
	Rate        float32 `gorm:"column:rate" json:"rate"`         // 使用率 单位%
}

func (s *DiskInfo) TableName() string {
	return "disk_info"
}
