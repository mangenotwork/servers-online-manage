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

//容器
type DockerContainerAction struct {
	//方法名，方法的标识
	Action string
	//数据包
	Packet []byte
	//错误
	Error error
}

//通用
type DockerAction struct {
	//方法名，方法的标识
	Action string
	//数据包
	Packet []byte
	//错误
	Error error
}

//获取docker info  & docker 版本  &  当前镜像数量 & 当前容器数量
type DockerBaseInfo struct {
	IsHave       bool     //当前slve是否有安装docker
	IsOpen       bool     //当前slve是否有启动docker
	OpenServers  []string //当前slve是否有启动docker的服务名
	Version      string   // docker version
	Info         string   // docker info
	ImagesNum    int      //docker 镜像数量
	ContainerNum int      //docker 容器数量
	InstallPath  string   //docker 按照路径
	//docker 配置文件位置
	//docker 进程ID
}
