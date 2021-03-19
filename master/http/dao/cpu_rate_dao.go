package dao

import (
	"github.com/mangenotwork/servers-online-manage/lib/dbconn"
	"github.com/mangenotwork/servers-online-manage/master/http/models"
)

type CPURateDao struct {
	Data  *models.CPURate
	Datas []*models.CPURate
}

// Create
func (d *CPURateDao) Create() (err error) {
	db := dbconn.Conn()
	defer db.Close()
	err = db.Model(models.CPURate{}).Create(&d.Data).Error
	return
}

// Creates
func (d *CPURateDao) Creates() (err error) {
	db := dbconn.Conn()
	defer db.Close()
	for _, v := range d.Datas {
		err = db.Model(models.CPURate{}).Create(&v).Error
	}
	return
}
