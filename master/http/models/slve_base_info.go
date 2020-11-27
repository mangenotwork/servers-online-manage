//selve基础信息
//记录连接过的seleve

package models

//Slve连接日志
type SlveConnLog struct {
	ID int64 `gorm:"primary_key;column:id" json:"id"`
}

func (s *SlveConnLog) TableName() string {
	return "slve_conn_log"
}
