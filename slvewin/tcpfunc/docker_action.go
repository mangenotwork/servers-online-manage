// docker 交互的路由
// 优先使用 docker sdk进行交互， 其次使用 cmd进行交互, 最后使用 docker Remote api 进行交互；

package tcpfunc

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/mangenotwork/servers-online-manage/lib/docker"
	"github.com/mangenotwork/servers-online-manage/lib/loger"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
	"github.com/mangenotwork/servers-online-manage/lib/utils"
	"github.com/mangenotwork/servers-online-manage/slve/tcpfunc/software"
)

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

func Images(action string) (data []byte, err error) {
	switch action {
	case "get_images_list":
		return docker.ImageList()
	case "list":
		return GETImageList()
	}
	return
}

func Container(action string) (data []byte, err error) {
	switch action {
	case "list":
		return GETContainerList()
	}
	return
}

//获取镜像列表
func GETImageList() (data []byte, err error) {
	imagesList := make([]*structs.ImageInfo, 0)
	ctx, cli, err := docker.GetClient()
	if err == nil {
		images, _ := docker.ImagesRun(cli, ctx)
		for _, v := range images {
			loger.Debug("images = ", v)
			i := &structs.ImageInfo{
				ID: v.ID,
			}
			i.Repository = strings.Join(v.RepoTags, ";")
			i.Digest = strings.Join(v.RepoDigests, ";")
			i.CreatedAt = utils.DateTime(v.Created)
			i.Size = fmt.Sprintf("%d M", 104120620/1000/1000) //按照 docker images 出来的结果为标准
			loger.Debug(i)
			imagesList = append(imagesList, i)
		}
		cli.Close()
	} else {
		images, imagesNum := new(software.Image).Get()
		loger.Debug(images, imagesNum)
		for _, v := range images {
			imagesList = append(imagesList, &structs.ImageInfo{
				ID:         v.ID,
				Repository: v.Repository,
				Tag:        v.Tag,
				Digest:     v.Digest,
				CreatedAt:  v.CreatedAt,
				Size:       v.Size,
			})
		}
	}
	return json.Marshal(imagesList)
}

//获取容器列表
func GETContainerList() (data []byte, err error) {
	containerList := make([]*structs.ContainerInfo, 0)
	ctx, cli, err := docker.GetClient()
	if err != nil {
		container, _ := docker.ContainerList(cli, ctx)
		for _, v := range container {
			loger.Debug("container = ", v)
			i := &structs.ContainerInfo{
				ID:        v.ID,
				Name:      strings.Join(v.Names, ";"),
				CreatedAt: utils.DateTime(v.Created),
				Image:     v.Image,
				ImageID:   v.ImageID,
				Command:   v.Command,
				State:     v.State,
				Status:    v.Status,
				Size:      fmt.Sprintf("%d", v.SizeRw),
				Networks:  v.HostConfig.NetworkMode,
			}
			for _, p := range v.Ports {
				i.Ports += fmt.Sprintf(" ip: %s; type: %s |", p.IP, p.Type)
			}
			for k, l := range v.Labels {
				i.Labels += fmt.Sprintf(" %s:%s |", k, l)
			}
			for _, m := range v.Mounts {
				i.Mounts += fmt.Sprintf(" %s |", m.Name)
			}
			loger.Debug("container = ", i)
			containerList = append(containerList, i)
		}
	} else {
		container, _ := new(software.Holder).Get()
		for _, v := range container {
			loger.Debug("container = ", v)
			containerList = append(containerList, &structs.ContainerInfo{
				ID:         v.ID,
				Name:       v.Names,
				CreatedAt:  v.CreatedAt,
				Image:      v.Image,
				Command:    v.Command,
				Status:     v.Status,
				RunningFor: v.RunningFor,
				Ports:      v.Ports,
				Size:       v.Size,
				Labels:     v.Labels,
				Mounts:     v.Mounts,
				Networks:   v.Networks,
			})
		}
	}
	return json.Marshal(containerList)
}
