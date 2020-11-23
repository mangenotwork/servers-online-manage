// docker api 对应的功能实现
package slve

import (
	"io/ioutil"
	"log"
	"net/http"
)

//  ============    容器

//列表容器 https://docs.docker.com/engine/api/v1.40/#tag/Container
// GET  /containers/json

//创建一个容器 Create a container
//doc: https://docs.docker.com/engine/api/v1.40/#operation/ContainerList
// POST  /containers/create

//检查一个容器 Inspect a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerInspect
// GET  /containers/{id}/json

//列出在容器内运行的进程   List processes running inside a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerInspect
// GET  /containers/{id}/top

//得到容器日志  Get container logs
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerLogs
//GET   /containers/{id}/logs

//获取容器文件系统上的更改   Get changes on a container’s filesystem
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerChanges
//GET   /containers/{id}/changes

//将容器的内容导出为压缩包。  Export a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerExport
//GET   /containers/{id}/export

//基于资源使用的容器统计信息   Get container stats based on resource usage
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerStats
//GET  /containers/{id}/stats

//调整容器TTY的大小  Resize a container TTY
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerResize
//POST  /containers/{id}/resize

//启动一个容器  Start a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerStart
//POST  /containers/{id}/start

//停止一个容器  Stop a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerStop
//POST  /containers/{id}/stop

//重启一个容器  Restart a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerRestart
//POST  /containers/{id}/restart

//杀死一个容器  Kill a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerKill
//POST  /containers/{id}/kill

//更新一个容器  Update a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerUpdate
//POST  /containers/{id}/update

//重命名一个容器  Rename a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerRename
//POST  /containers/{id}/rename

//暂停一个容器  Pause a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerPause
//POST  /containers/{id}/pause

//取消暂停一个容器  Unpause a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerUnpause
//POST  /containers/{id}/unpause

//Attach to a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerAttach
//POST  /containers/{id}/attach

//Attach to a container via a websocket
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerAttachWebsocket
//GET  /containers/{id}/attach/ws

//等待一个容器  Wait for a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerWait
//POST  /containers/{id}/wait

//删除一个容器  Remove a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerDelete
//DELETE  /containers/{id}

//获取容器中文件的信息  Get information about files in a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerArchiveInfo

//获取容器中文件系统资源的存档  Get an archive of a filesystem resource in a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ContainerArchive

//将文件或文件夹的存档解压到容器中的目录中  Extract an archive of files or folders to a directory in a container
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/PutContainerArchive

//删除停止容器  Delete stopped containers
//doc: https://docs.docker.com/engine/api/v1.40/#operation/ContainerPrune

//  =================   镜像

//镜像列表 images
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImageList
func ImageList() []byte{
	url := "http://0.0.0.0:12225" + "/images/json"
	resp, err := http.Get(url)
	log.Println(resp,resp.Body, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	return body
}

//打包镜像 Build an image
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImageBuild

//删除builder缓存  Delete builder cache
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/BuildPrune

//创建一个镜像  Create an image
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImageCreate

//检查一个镜像
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImageInspect

//获取镜像的历史记录
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImageHistory

//推送一个镜像
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImagePush

//给一个镜像添加tag
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImageTag

//删除一个镜像
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImageDelete

//搜索一个镜像
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImageSearch

//删除未使用的镜像
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImagePrune

//将指定容器创建一个新的镜像
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImageCommit

//导出一个镜像
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImageGet

//导出多个镜像
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImageGetAll

//导入一个镜像
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImageLoad

//网络列表
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/NetworkList

//检查网络
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/NetworkInspect

//删除网络
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/NetworkDelete

//创建一个网络
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/NetworkCreate

//将容器连接到网络
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/NetworkConnect

//
