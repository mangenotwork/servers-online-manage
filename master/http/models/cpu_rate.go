// CPU 使用率
package models

//Slve CPU使用率
type CPURate struct {
	ID       int64   `gorm:"primary_key;column:id" json:"id"`
	SlveUUID string  `gorm:"column:uuid" json:"uuid"`
	Time     int64   `gorm:"column:time" json:"time"`       //采集时间
	IsMain   int64   `gorm:"column:is_main" json:"is_main"` // 记总  true: 1, false: 0
	CPU      string  `gorm:"column:cpu" json:"cpu"`         // 核心名,核心数
	UseRate  float32 `gorm:"column:use" json:"use"`         // 使用率
}

func (s *CPURate) TableName() string {
	return "cpu_rate"
}
