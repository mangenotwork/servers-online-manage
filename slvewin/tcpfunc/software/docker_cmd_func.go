//docker 命令交互
// 通过命令，由系统自己执行然后进行操作docker
package software

import (
	"fmt"
	"github.com/mangenotwork/servers-online-manage/lib/cmd"
	"strings"

)

//判断本机是否存在Docker
func IsDocker() (bool, string) {
	cmds := []string{"docker", "-v"}
	dockerstr := cmd.WindowsSendCommand(cmds)
	fmt.Println(dockerstr)
	if strings.Index(dockerstr, "Docker") != -1 || strings.Index(dockerstr, "version") != -1 {
		return true, dockerstr
	}
	return false, ""
}

// ========================   Images

//本机镜像
type Image struct {
	ID         string `json:"image_name"`       //镜像ID
	Repository string `json:"image_repository"` //镜像名
	Tag        string `json:"image_tag"`        //镜像Tag
	Digest     string `json:"image_id"`         //镜像 digest
	CreatedAt  string `json:"image_created"`    //镜像创建
	Size       string `json:"image_size"`       //镜像大小
}

//执行的Cmd docker ps
func (image *Image) CmdDockerImages() string {
	//执行的命令
	imageData := "{{.Repository}}|{{.Tag}}|{{.Digest}}|{{.ID}}|{{.CreatedAt}}|{{.Size}}"
	return fmt.Sprintf("docker images -a --format \"%s\"", imageData)
}

//获取本机当前的镜像
func (image *Image) Get() ([]*Image, int) {
	cmds := []string{"cmd", "/C", image.CmdDockerImages()}
	imageinfos := cmd.WindowsSendCommand(cmds)
	fmt.Println(imageinfos)
	return image.OutPutImages(imageinfos)
}

//输出为当前的镜像信息
func (image *Image) OutPutImages(imagesstr string) ([]*Image, int) {
	dockerpsInfo := strings.Split(imagesstr, "\n")
	ImagesNow := make([]*Image, 0)
	for _, dockerpsInfos := range dockerpsInfo {
		dockerpsInfos = strings.Replace(dockerpsInfos, "\"", "", -1)

		fmt.Println(dockerpsInfos)
		imagesInfos := strings.Split(dockerpsInfos, "|")
		fmt.Println(imagesInfos)
		if len(imagesInfos) == 6 {

			imageopt := &Image{
				ID:         string(imagesInfos[0]),
				Repository: string(imagesInfos[1]),
				Tag:        string(imagesInfos[2]),
				Digest:     string(imagesInfos[3]),
				CreatedAt:  string(imagesInfos[4]),
				Size:       string(imagesInfos[5]),
			}
			ImagesNow = append(ImagesNow, imageopt)
		}
		fmt.Println(len(imagesInfos))
		fmt.Println("\n\n")
	}
	fmt.Println(ImagesNow)
	for number, hoder := range ImagesNow {
		fmt.Println(number, hoder)
	}
	return ImagesNow, len(ImagesNow)
}

//删除镜像  执行  docker rmi id
func (image *Image) RMI(id string) bool {
	cmds := []string{"cmd", "/C", fmt.Sprintf("docker rmi -f %s", id)}
	psinfos := cmd.WindowsSendCommand(cmds)
	fmt.Println("\n\n\n\n删除 -- 》", psinfos)
	if psinfos == "" {
		return false
	}
	return true
}


//  =============================   container (容器)

//PS 结构体
//本机容器
type Holder struct {
	ID         string `json:"holder_id"`       //容器ID
	Names      string `json:"holder_name"`     //容器名称
	CreatedAt  string `json:"holder_create"`   //容器创建时间
	Status     string `json:"holder_status"`   //容器状态
	RunningFor string `json:"holder_running"`  //容器运行时间
	Ports      string `json:"holder_port"`     //容器使用的端口
	Size       string `json:"holder_size"`     //容器大小
	Command    string `json:"holder_command"`  //运行容器使用的命令
	Image      string `json:"holder_image"`    //运行容器基于的镜像
	Labels     string `json:"holder_labels"`   // 容器 Labels
	Mounts     string `json:"holder_mounts"`   // 容器  Mounts
	Networks   string `json:"holder_networks"` // 容器 Networks
}

//获取本机当前正在运行的容器
func (holder *Holder) Get() ([]*Holder, int) {
	cmds := []string{"cmd", "/C", holder.CmdDockerPsAll()}
	psinfos := cmd.WindowsSendCommand(cmds)
	fmt.Println(psinfos)
	return holder.OutPutHolder(psinfos)
}

//执行的Cmd docker ps -a
func (holder *Holder) CmdDockerPsAll() string {
	//执行的命令
	holderData := "{{.ID}}|{{.Names}}|{{.CreatedAt}}|{{.Status}}|{{.RunningFor}}|{{.Ports}}|" +
		"{{.Size}}|{{.Command}}|{{.Image}}|{{.Labels}}|{{.Mounts}}|{{.Networks}}"
	return fmt.Sprintf("docker ps -a --format \"%s\"", holderData)
}

// 将输入的名片 输出为当前的容器信息
func (holder *Holder) OutPutHolder(hoderstr string) ([]*Holder, int) {
	dockerpsInfo := strings.Split(hoderstr, "\n")
	holderNow := make([]*Holder, 0)
	for _, dockerpsInfos := range dockerpsInfo {
		dockerpsInfos = strings.Replace(dockerpsInfos, "\"", "", -1)

		fmt.Println(dockerpsInfos)
		holderInfos := strings.Split(dockerpsInfos, "|")
		fmt.Println(holderInfos)
		if len(holderInfos) == 12 {

			holderopt := &Holder{
				ID:        string(holderInfos[0]),
				Names:     string(holderInfos[1]),
				CreatedAt: string(holderInfos[2]),
				//Status:     HolderIsExited(string(holderInfos[3])), //是否关闭
				Status:     string(holderInfos[3]),
				RunningFor: string(holderInfos[4]),
				Ports:      string(holderInfos[5]),
				Size:       string(holderInfos[6]),
				Command:    string(holderInfos[7]),
				Image:      string(holderInfos[8]),
				Labels:     string(holderInfos[9]),
				Mounts:     string(holderInfos[10]),
				Networks:   string(holderInfos[11]),
			}
			holderNow = append(holderNow, holderopt)
		}
		fmt.Println(len(holderInfos))
		fmt.Println("\n\n")
	}
	fmt.Println(holderNow)
	for number, hoder := range holderNow {
		fmt.Println(number, hoder)
	}
	return holderNow, len(holderNow)
}

//判断容器是否被关闭
func HolderIsExited(holderState string) string {
	isInt := strings.Index(holderState, "Exited")
	if isInt > 0 {
		return holderState
	}
	return "已停止"
}

//执行的Cmd docker rm id
func (holder *Holder) RM(id string) bool {
	cmds := []string{"cmd", "/C", fmt.Sprintf("docker rm %s", id)}
	psinfos := cmd.WindowsSendCommand(cmds)
	fmt.Println("\n\n\n\n删除 -- 》", psinfos)
	if psinfos == "" {
		return false
	}
	return true
}

//关闭容器  coder kill id
func (holder *Holder) KILL(id string) bool {
	cmds := []string{"cmd", "/C", fmt.Sprintf("docker kill %s", id)}
	psinfos := cmd.WindowsSendCommand(cmds)
	fmt.Println("\n\n\n关闭容器 -- 》", psinfos)
	if psinfos == "" {
		return false
	}
	return true
}
