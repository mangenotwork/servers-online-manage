// docker Remote api 对应的功能实现
// 这里我简述为 通过http 操作 docker
package docker

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
func ImageList() (body []byte, err error){
	params := url.Values{}
	Url, err := url.Parse("http://0.0.0.0:12225/images/json")
	if err != nil {
		log.Println(err)
		return
	}

	//all  是否显示所有
	params.Set("all","true")
	//digests    摘要
	//默认 true
	params.Set("digests","true")

	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()

	resp, err := http.Get(urlPath)
	if err != nil{
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

//TODO
//打包镜像 Build an image
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImageBuild
func ImageBuild(){
	params := url.Values{}
	Url, err := url.Parse("http://0.0.0.0:12225/build")
	if err != nil {
		log.Println(err)
		return
	}

	//dockerfile
	//默认：Dockerfile
	//在构建上下文中路径到Dockerfile。如果指定了remote并指向外部Dockerfile，则忽略此选项。
	params.Set("dockerfile","Dockerfile")

	//t   添加一个 tag
	//默认： latest
	params.Set("t","latest")

	//extrahosts  要添加到/etc/hosts的额外主机

	//remote  string

	//q   bool
	//默认：false
	//按详细的构建输出。
	params.Set("q","false")

	//nocache  bool
	//默认：false
	//构建时不使用缓存。
	params.Set("nocache","false")

	//cachefrom  string
	//JSON array of images used for build cache resolution.

	//pull
	//即使本地存在较旧的镜像，也尝试拉取镜像。

	//rm
	//默认：true
	//在成功构建后删除中间容器。

	//forcerm
	//默认： false
	//始终删除中间容器，即使出现故障。

	//memory
	//内存限制。

	//memswap
	//Total memory (memory + swap). Set as -1 to disable swap.

	//cpushares
	//CPU共享(相对权重)

	//cpusetcpus
	//CPUs in which to allow execution (e.g., 0-3, 0,1).

	log.Println(Url)
}

//删除builder缓存  Delete builder cache
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/BuildPrune

//创建一个镜像  Create an image
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImageCreate
func ImageCreate() {

	params := url.Values{}
	Url, err := url.Parse("http://0.0.0.0:12225/images/create")
	if err != nil {
		log.Println(err)
		return
	}

	//fromImage
	//要提取的镜像名称。名称可以包括标签或摘要。此参数只能在导入映像时使用。如果HTTP连接关闭，拉取操作将被取消
	params.Set("fromImage","nginx")

	//fromSrc
	//要导入的源。该值可以是一个URL，可以从中检索图像，或者-从请求主体中读取图像。此参数只能在导入映像时使用。

	//repo
	//导入映像时提供给映像的存储库名称。回购可以包括一个标签。此参数只能在导入映像时使用。

	//tag
	//导入镜像的tag

	//message
	//为导入的镜像提交消息。

	//platform
	//Platform in the format os[/arch[/variant]]

	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	log.Println(urlPath)
	resp, err := http.Post(urlPath, "application/x-www-form-urlencoded", nil)
	log.Println(resp, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body),err)
}

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
func ImageDelete(){
	params := url.Values{}
	Url, err := url.Parse("http://0.0.0.0:12225/images/5a683654e5d5")
	if err != nil {
		log.Println(err)
		return
	}
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	log.Println(urlPath)

	req, _ := http.NewRequest("DELETE", urlPath, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
}

//搜索一个镜像
//doc:  https://docs.docker.com/engine/api/v1.40/#operation/ImageSearch
func ImageSearch(term,limit string){
	params := url.Values{}
	Url, err := url.Parse("http://0.0.0.0:12225/images/search")
	if err != nil {
		log.Println(err)
		return
	}

	//添加参数
	//term
	//搜索词
	params.Set("term",term)

	//limit
	if limit == ""{
		limit = "10"
	}
	params.Set("limit",limit)

	//filters
	//A JSON encoded value of the filters (a map[string][]string) to process on the images list. Available filters:
	//筛选参数:
	//is-automated=(true|false)
	//is-official=(true|false)
	//stars=<number> Matches images that has at least 'number' stars.

	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()

	resp, err := http.Get(urlPath)
	if err != nil{
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body), err)
}

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
func ImageLoad(){

}

//    =================   网络

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
