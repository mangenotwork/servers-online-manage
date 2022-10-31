package dao

import (
	"github.com/mangenotwork/servers-online-manage/lib/dbconn"
	"github.com/mangenotwork/servers-online-manage/master/http/models"
)

type PerformanceCountDao struct {
	Data  *models.PerformanceCount
	Datas []*models.PerformanceCount
}

// Create
func (d *PerformanceCountDao) Create() (err error) {
	db := dbconn.Conn()
	defer db.Close()
	err = db.Model(models.PerformanceCount{}).Create(&d.Data).Error
	return
}
