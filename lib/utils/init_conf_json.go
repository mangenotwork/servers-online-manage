// 读取初始化配置
package utils

import (
	"encoding/json"
	"os"

	"github.com/mangenotwork/servers-online-manage/lib/global"
	"github.com/mangenotwork/servers-online-manage/lib/loger"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
)

//初始化配置
func InitMasterConf() *structs.MasterConf {
	//读取配置文件
	file, _ := os.Open("conf/master_conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	masterconf := &structs.MasterConf{}
	err := decoder.Decode(&masterconf)
	if err != nil {
		loger.Error("Error:", err)
	}
	loger.Debug("masterconf = ", &masterconf)

	//给全局变量赋值
	global.Version = masterconf.Version
	global.MasterHost = masterconf.MasterHost

	return masterconf
}

func InitSlveConf() {
	//读取配置文件
	file, _ := os.Open("conf/slve_conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	slveconf := structs.SlveConf{}
	err := decoder.Decode(&slveconf)
	if err != nil {
		loger.Error("Error:", err)
		loger.Error("读取配置文件 conf/slve_conf.json 失败 ！")
		os.Exit(1)
	}
	loger.Debug("slveconf = ", &slveconf)

	//给全局变量赋值
	global.SlveVersion = slveconf.Version
	global.MasterHost = slveconf.MasterHost
	global.SlveSpace = slveconf.SlveSpace
	global.SlveUUID = GetMD5MAC()
}
