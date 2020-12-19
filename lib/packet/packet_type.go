package packet

//数据包类型
const (
	//slve 发送心跳包
	HEART_BEAT_PACKET = 0x00

	//slve 连接成功收到token，返回的第一包
	FIRST_PACKET = 0x99

	//slve 发送数据包
	REPORT_PACKET = 0x01

	//master 颁发token
	SET_SLVE_TOKEN_PACKET = 0x02

	//master 答复slve心跳包
	REPLY_HEART_PACKET = 0x03

	//获取slve host信息
	Get_SLVE_INFO_PACKET = 0x04

	//接收slve 发送数据包
	RECEPTION_SLVE_PACKET = 0x05

	//发送命令让slve执行
	SET_SLVE_CMD_PACKET = 0x06

	//发送文件
	SEND_FILE_PACKET = 0x07

	//发送文件完成
	SEND_FILE_COMPLETE_PACKET = 0x08

	//请求docker images相关
	Docker_Images = 0x09

	//请求docker 信息相关
	Docker_Infos = 0x10

	//请求docker container相关
	Docker_Container = 0x11
)
