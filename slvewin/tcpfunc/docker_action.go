// docker 交互的路由
// 优先使用 docker sdk进行交互， 其次使用 cmd进行交互, 最后使用 docker Remote api 进行交互；

package tcpfunc

import (
	"encoding/json"
	"log"

	"github.com/mangenotwork/servers-online-manage/lib/structs"

	"github.com/mangenotwork/servers-online-manage/lib/docker"
	"github.com/mangenotwork/servers-online-manage/slve/tcpfunc/software"
)

func Images(action string) (data []byte, err error) {
	switch action {
	case "get_images_list":
		return docker.ImageList()
	}
	return
}

func GetInfo() (data []byte, err error) {
	isHave, _ := software.IsDocker()
	isOpen, dockerServers := software.IsOpenDockerServer()
	versionInfo := software.CmdDockerVersion()
	info := software.CmdDockerInfo()
	imagesNum := 0
	containerNum := 0

	//是否开启了docker api
	isAPI := false
	ctx, cli, err := docker.GetClient()
	if err == nil {
		isAPI = true
	}

	//开启了走SDK, 否则走cmd
	if !isAPI {

		//获取镜像列表
		images, _ := docker.ImagesRun(cli, ctx)
		imagesNum = len(images)

		//获取容器列表
		container, _ := docker.ContainerList(cli, ctx)
		containerNum = len(container)

		//使用完了就关闭连接
		cli.Close()
	} else {
		//获取镜像列表
		_, imagesNum = new(software.Image).Get()

		//获取容器列表
		_, containerNum = new(software.Holder).Get()
	}

	//log.Println(isHave, isOpen, dockerServers, versionInfo, info, imagesNum, containerNum)

	baseInfo := structs.DockerBaseInfo{
		IsHave:       isHave,
		IsOpen:       isOpen,
		OpenServers:  dockerServers,
		Version:      versionInfo,
		Info:         info,
		ImagesNum:    imagesNum,
		ContainerNum: containerNum,
		InstallPath:  "windows 占不支持获取安装路径",
	}
	log.Println(baseInfo)

	return json.Marshal(baseInfo)
}
