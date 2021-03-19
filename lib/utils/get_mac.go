// 获取Mac
package utils

import (
	"log"
	"net"
)

func GetMAC() (mac string) {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	log.Println(interfaces)
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	for _, v := range interfaces {
		mac = v.HardwareAddr.String()
		if mac != "" {
			log.Println("MAC = ", mac)
			return
		}
	}
	return
}

func GetMD5MAC() (mac string) {
	mac = GetMAC()
	mac = Str2MD5(mac)
	return
}
