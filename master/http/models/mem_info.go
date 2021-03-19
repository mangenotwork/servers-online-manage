// 内存信息
package models

//Slve 内存信息
type MEMInfo struct {
	ID       int64   `gorm:"primary_key;column:id" json:"id"`
	SlveUUID string  `gorm:"column:uuid" json:"uuid"`
	Time     int64   `gorm:"column:time" json:"time"`       //采集时间
	Total    int64   `gorm:"column:total" json:"total"`     // 所有可用RAM大小
	Used     int64   `gorm:"column:used" json:"used"`       // 内存使用
	Free     int64   `gorm:"column:free" json:"free"`       // 可用
	Rate     float32 `gorm:"column:rate" json:"rate"`       // 使用率 单位%
	Buffers  int64   `gorm:"column:buffers" json:"buffers"` //给文件做缓冲大小
	Cached   int64   `gorm:"column:cached" json:"cached"`   //被高速缓冲存储器
}

func (s *MEMInfo) TableName() string {
	return "mem_info"
}
