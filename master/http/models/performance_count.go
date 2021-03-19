// 系统计数, 进程数, 连接数等等
package models

//Slve 系统各种计数
type PerformanceCount struct {
	ID           int64  `gorm:"primary_key;column:id" json:"id"`
	SlveUUID     string `gorm:"column:uuid" json:"uuid"`
	Time         int64  `gorm:"column:time" json:"time"`                     //采集时间
	TcpConnCount int    `gorm:"column:tcp_conn_count" json:"tcp_conn_count"` //连接数
	PIDCount     int    `gorm:"column:pid_count" json:"pid_count"`           //进程数
}

func (s *PerformanceCount) TableName() string {
	return "performance_count"
}
