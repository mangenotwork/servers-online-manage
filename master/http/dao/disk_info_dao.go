package dao

import (
	"github.com/mangenotwork/servers-online-manage/lib/dbconn"
	"github.com/mangenotwork/servers-online-manage/master/http/models"
)

type DiskInfoDao struct {
	Data  *models.DiskInfo
	Datas []*models.DiskInfo
}

// Create
func (d *DiskInfoDao) Create() (err error) {
	db := dbconn.Conn()
	defer db.Close()
	err = db.Model(models.DiskInfo{}).Create(&d.Data).Error
	return
}

// Creates
func (d *DiskInfoDao) Creates() (err error) {
	db := dbconn.Conn()
	defer db.Close()
	//目前 gorm 并不支持批量插入这一功能，但已经被列入 v2.0 的计划里面了,详细讨论你可以参见 issues-255
	//这里可采用拼接sql实现
	//为了方便我这里采用循环
	for _, v := range d.Datas {
		err = db.Model(models.DiskInfo{}).Create(&v).Error
	}
	return
}

//
func (d *DiskInfoDao) GetFromTimes(t []int64) (err error) {
	db := dbconn.Conn()
	defer db.Close()
	err = db.Model(models.DiskInfo{}).Where("time in (?)", t).Find(&d.Datas).Error
	return
}

// 图表需要的数据,平均取n个
func (d *DiskInfoDao) EchartData() []float32 {
	showData := make([]float32, 0)
	for _, v := range d.Datas {
		showData = append(showData, v.Rate)
	}
	return showData
}
