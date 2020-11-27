//本地数据库初始化
//数据库采用 sqlist
//包含了初始化创建

package db

import (
	"github.com/mangenotwork/servers-online-manage/lib/global"
	"github.com/mangenotwork/servers-online-manage/lib/dbconn"
	"github.com/mangenotwork/servers-online-manage/master/http/models"
	"log"
	"os"
)

//检查数据sqlist 是否初始化
func CheckSqlitDB(dbFile string) bool {
	//没有设置就默认在根目录下
	if dbFile == ""{
		dbFile = "./db/base.db"
	}
	global.SqlistDBPath = dbFile
	_, err := os.Lstat(dbFile)
	//没有则创建
	if os.IsNotExist(err){
		f,err := os.Create(dbFile)
		defer f.Close()
		if err !=nil {
			log.Println(err.Error())
		}
	}

	db := dbconn.Conn()
	defer db.Close()
	log.Println("db conn -> ", db)

	//检查table
	//如果没有就创建
	if  !db.HasTable(&models.Notifincation{}){
		log.Println("Notifincation 不存在")
		db.Set("gorm:notifincation", "ENGINE=InnoDB").CreateTable(&models.Notifincation{})
	}
	if  !db.HasTable(&models.SlveConnLog{}){
		log.Println("SeleveConnLog 不存在")
		db.Set("gorm:slve_conn_log", "ENGINE=InnoDB").CreateTable(&models.SlveConnLog{})
	}
	//如果没有用户表，则创建，并生成root admin 账号，密码随机生成并生成一个txt在根目录下

	return true
}