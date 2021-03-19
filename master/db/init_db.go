//本地数据库初始化
//数据库采用 sqlist
//包含了初始化创建

package db

import (
	"log"
	"os"

	"github.com/mangenotwork/servers-online-manage/lib/dbconn"
	"github.com/mangenotwork/servers-online-manage/lib/global"
	"github.com/mangenotwork/servers-online-manage/master/http/models"
)

//检查数据sqlist 是否初始化
func CheckSqlitDB(dbFile string) bool {
	//没有设置就默认在根目录下
	if dbFile == "" {
		dbFile = "./db/base.db"
	}
	global.SqlistDBPath = dbFile
	_, err := os.Lstat(dbFile)
	//没有则创建
	if os.IsNotExist(err) {
		f, err := os.Create(dbFile)
		defer f.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}

	db := dbconn.Conn()
	defer db.Close()
	log.Println("db conn -> ", db)

	modelsMap := map[string]interface{}{
		"notifincation":     &models.Notifincation{},
		"slve_base_info":    &models.SlveBaseInfo{},
		"cpu_rate":          &models.CPURate{},
		"disk_info":         &models.DiskInfo{},
		"mem_info":          &models.MEMInfo{},
		"network_io":        &models.NetworkIO{},
		"performance_count": &models.PerformanceCount{},
	}
	// //检查table
	// //如果没有就创建
	for k, v := range modelsMap {
		if !db.HasTable(v) {
			log.Println(k, "不存在")
			db.Set("gorm:"+k, "ENGINE=InnoDB").CreateTable(v)
		}
	}

	return true
}
