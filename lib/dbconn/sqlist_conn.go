// sqlist 连接
//使用 github.com/jinzhu/gorm

package dbconn

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mangenotwork/servers-online-manage/lib/global"
)

func Conn() (db *gorm.DB) {
	var err error
	db, err = gorm.Open("sqlite3", global.SqlistDBPath)
	db.LogMode(true)
	if err != nil {
		log.Println(err)
		log.Println("连接数据库失败")
		return
	}
	return
}
