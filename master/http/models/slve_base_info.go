//selve基础信息
//记录连接过的seleve

package models

//Slve基础信息
type SlveBaseInfo struct {
	SlveUUID        string `gorm:"primary_key;column:uuid" json:"uuid"`
	Name            string `gorm:"column:host_name" json:"host_name"`
	SetName         string `grom:"column:name" json:"name"`
	HostIP          string `grom:"column:host_ip" json:"host_ip"`
	SysType         string `grom:"column:sys_type" json:"sys_type"`
	SlveVersion     string `grom:"column:slve_version" json:"slve_version"`
	LastConn        string `grom:"column:last_conn_time" json:"last_conn_time"`
	OsName          string `grom:"column:os_name" json:"os_name"`
	SysArchitecture string `grom:"column:sys_architecture" json:"sys_architecture"`
	CpuCoreNumber   string `grom:"column:cpu_core_number" json:"cpu_core_number"`
	CpuName         string `grom:"column:cpu_name" json:"cpu_name"`
	CpuID           string `grom:"column:cpu_id" json:"cpu_id"`
	BaseBoardID     string `grom:"column:board_id" json:"board_id"`
	MemTotal        string `grom:"column:mem_totle" json:"mem_totle"`
	DiskTotal       string `grom:"column:disk_totle" json:"disk_totle"`
	Note            string `grom:"column:note" json:"note"` //备注
	Online          bool   `grom:"-" json:"online"`
	MemUsed         string `grom:"column:mem_used" json:"mem_used"`   //以使用的内存
	DiskUsed        string `grom:"column:disk_used" json:"disk_used"` //磁盘使用
}

func (s *SlveBaseInfo) TableName() string {
	return "slve_base_info"
}
