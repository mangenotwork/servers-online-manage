package dao

import (
	"github.com/mangenotwork/servers-online-manage/lib/dbconn"
	"github.com/mangenotwork/servers-online-manage/master/http/models"
)

type NetworkIODao struct {
	Data  *models.NetworkIO
	Datas []*models.NetworkIO
}

// Create
func (d *NetworkIODao) Create() (err error) {
	db := dbconn.Conn()
	defer db.Close()
	err = db.Model(models.NetworkIO{}).Create(&d.Data).Error
	return
}

// Creates
func (d *NetworkIODao) Creates() (err error) {
	db := dbconn.Conn()
	defer db.Close()
	for _, v := range d.Datas {
		err = db.Model(models.NetworkIO{}).Create(&v).Error
	}
	return
}
