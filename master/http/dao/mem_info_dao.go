package dao

import (
	"github.com/mangenotwork/servers-online-manage/lib/dbconn"
	"github.com/mangenotwork/servers-online-manage/master/http/models"
)

type MEMInfoDao struct {
	Data  *models.MEMInfo
	Datas []*models.MEMInfo
}

// Create
func (d *MEMInfoDao) Create() (err error) {
	db := dbconn.Conn()
	defer db.Close()
	err = db.Model(models.MEMInfo{}).Create(&d.Data).Error
	return
}
