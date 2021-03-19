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

//
func (d *MEMInfoDao) GetFromTimes(t []int64) (err error) {
	db := dbconn.Conn()
	defer db.Close()
	err = db.Model(models.MEMInfo{}).Where("time in (?)", t).Find(&d.Datas).Error
	return
}

// 图表需要的数据,平均取n个
func (d *MEMInfoDao) EchartData() []float32 {
	showData := make([]float32, 0)
	for _, v := range d.Datas {
		showData = append(showData, v.Rate)
	}
	return showData
}
