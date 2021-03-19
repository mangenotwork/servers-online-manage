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

//
func (d *NetworkIODao) GetFromTimes(t []int64) (err error) {
	db := dbconn.Conn()
	defer db.Close()
	err = db.Model(models.NetworkIO{}).Where("time in (?)", t).Find(&d.Datas).Error
	return
}

// 图表需要的数据,平均取n个
func (d *NetworkIODao) EchartData() (map[string][]float32, map[string][]float32) {
	txshowData := make(map[string][]float32, 0)
	rxshowData := make(map[string][]float32, 0)
	for _, v := range d.Datas {
		txshowData[v.Name] = append(txshowData[v.Name], v.Tx)
		rxshowData[v.Name] = append(txshowData[v.Name], v.Rx)
	}
	return txshowData, rxshowData
}
