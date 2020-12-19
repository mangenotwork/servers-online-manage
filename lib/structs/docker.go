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

//镜像信息
type ImageInfo struct {
	ID         string `json:"id"`         //镜像ID
	Repository string `json:"repository"` //镜像名
	Tag        string `json:"tag"`        //镜像Tag
	Digest     string `json:"digest"`     //镜像 digest
	CreatedAt  string `json:"created"`    //镜像创建
	Size       string `json:"size"`       //镜像大小
}

//容器信息
type ContainerInfo struct {
	ID         string `json:"id"`          //容器id
	Name       string `json:"name"`        //容器名
	CreatedAt  string `json:"created"`     //容器创建时间
	Image      string `json:"image"`       //容器镜像
	ImageID    string `json:"image_id"`    //容器镜像id
	Command    string `json:"command"`     //执行容器的命令
	State      string `json:"state"`       //容器状态
	Status     string `json:"status"`      //容器状态
	RunningFor string `json:"running_for"` //容器运行时间
	Ports      string `json:"port_info"`   //容器使用的端口
	Size       string `json:"size"`        //容器大小
	Labels     string `json:"labels"`      //容器 Labels
	Mounts     string `json:"mounts"`      //容器  Mounts
	Networks   string `json:"networks"`    //容器 Networks
}
