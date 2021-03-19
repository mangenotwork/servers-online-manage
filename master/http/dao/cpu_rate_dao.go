package dao

import (
	"time"

	"github.com/mangenotwork/servers-online-manage/lib/dbconn"
	"github.com/mangenotwork/servers-online-manage/lib/utils"
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

// 获取指定时间段所有数据
// 单位为小时
func (d *CPURateDao) GetListFromTimeMainCPU(uuid string, h int64) (err error) {
	db := dbconn.Conn()
	defer db.Close()
	uinx := time.Now().Unix() - h*3600
	err = db.Model(models.CPURate{}).Where("uuid=? and time>? and is_main=1", uuid, uinx).Find(&d.Datas).Error
	return
}

// 图表需要的数据,平均取n个
func (d *CPURateDao) EchartData(n int) ([]int64, []string, []float32) {
	number := len(d.Datas)
	showData := make([]float32, 0)
	timeList := make([]int64, 0)
	showTime := make([]string, 0)
	if number > n {
		addNumber := number / n
		for k, v := range d.Datas {
			if k%addNumber == 0 {
				timeList = append(timeList, v.Time)
				showTime = append(showTime, utils.DateTime(v.Time))
				showData = append(showData, v.UseRate)
			}
		}
	}
	return timeList, showTime, showData
}
