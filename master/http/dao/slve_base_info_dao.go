package dao

import (
	"strings"

	"github.com/mangenotwork/servers-online-manage/lib/dbconn"
	"github.com/mangenotwork/servers-online-manage/lib/global"
	"github.com/mangenotwork/servers-online-manage/master/http/models"
)

type SlveBaseInfoDao struct {
	Data  *models.SlveBaseInfo
	Datas []*models.SlveBaseInfo
}

// Create
func (d *SlveBaseInfoDao) Create() (err error) {
	db := dbconn.Conn()
	defer db.Close()
	err = db.Model(models.SlveBaseInfo{}).Create(&d.Data).Error
	return
}

// 是否存在
func (d *SlveBaseInfoDao) IsHave(uuid string) bool {
	db := dbconn.Conn()
	defer db.Close()
	var count int = 0
	db.Model(models.SlveBaseInfo{}).Where("uuid=?", uuid).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

// Update
func (d *SlveBaseInfoDao) Update() (err error) {
	db := dbconn.Conn()
	defer db.Close()
	err = db.Model(models.SlveBaseInfo{}).Update(&d.Data).Error
	return
}

// 获取slve list
func (d *SlveBaseInfoDao) Gets() (err error) {
	db := dbconn.Conn()
	defer db.Close()
	err = db.Model(models.SlveBaseInfo{}).Order("last_conn asc").Find(&d.Datas).Error
	return
}

// slve list 查看在线与数据整理
func (d *SlveBaseInfoDao) IsOnlines() {
	for _, v := range d.Datas {
		ip := strings.Split(v.HostIP, ":")[0]
		_, v.Online = global.Slves[ip]
	}
}

func (d *SlveBaseInfoDao) IsOnline() {
	ip := strings.Split(d.Data.HostIP, ":")[0]
	_, d.Data.Online = global.Slves[ip]

}

// 通过uuid 获取数据
func (d *SlveBaseInfoDao) GetFromUUID(uuid string) (err error) {
	db := dbconn.Conn()
	defer db.Close()
	d.Data = &models.SlveBaseInfo{}
	err = db.Model(d.Data).Where("uuid=?", uuid).First(&d.Data).Error
	return
}
