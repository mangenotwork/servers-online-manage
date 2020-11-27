package structs

//Images func request
//Master发起 docker 镜像请求方法
//Slve返回结果走的结构体
type DockerImagesAction struct {
	//方法名，方法的标识
	Action string
	//数据包
	Packet []byte
	//错误
	Error error
}