// docker SDK 对应的功能实现
// 这里我简述为 通过 docker golang SDK 操作 docker

package docker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/docker/go-connections/nat"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func Run1() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	//ImagesRun(cli,ctx)
	//GetClientVersion(cli,ctx)
	//GetConfigList(cli,ctx)
	//RunContainer(cli,ctx)
	//ContainerInspect(cli,ctx)
	//ContainerKill(cli,ctx)
	//ContainerList(cli,ctx)
	//ContainerRemove(cli,ctx)
	//ContainerRename(cli,ctx)
	//ContainerRestart(cli,ctx)
	//ContainerStatPath(cli,ctx)
	//ContainerStop(cli,ctx)
	//ContainerTop(cli,ctx)
	//CopyFromContainer(cli,ctx)
	//CustomHTTPHeaders(cli,ctx)
	//DaemonHost(cli,ctx)
	//DiskUsage(cli,ctx)
	//DistributionInspect(cli,ctx)
	//SDKImageCreate(cli,ctx)
	//ImageInspectWithRaw(cli,ctx)
	//ImagePull(cli,ctx)
	//ImageRemove(cli,ctx)
	//SDKImageSearch(cli,ctx)
	//ImageTag(cli,ctx)
	//Info(cli,ctx)
	//NetworkList(cli,ctx)
	//NodeList(cli,ctx)
	//Ping(cli,ctx)
	//PluginInspectWithRaw(cli,ctx)
	//PluginList(cli,ctx)
	//RegistryLogin(cli,ctx)
	SecretList(cli, ctx)
}

func GetClient() (ctx context.Context, cli *client.Client, err error) {
	ctx = context.Background()
	cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	return
}

//docker images 列表
func ImagesRun(cli *client.Client, ctx context.Context) ([]types.ImageSummary, error) {
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, image := range images {
		log.Println(&image)
	}
	return images, err
}

//返回此客户机使用的API版本。
func GetClientVersion(cli *client.Client, ctx context.Context) {
	version := cli.ClientVersion()
	log.Println("回此客户机使用的API版本 = ", version)
}

//关闭客户端使用的传输
func ClientClose(cli *client.Client, ctx context.Context) {
	err := cli.Close()
	log.Println("关闭客户端使用的传输, err = ", err)
}

//返回配置的列表
func GetConfigList(cli *client.Client, ctx context.Context) {
	options := types.ConfigListOptions{}
	c, err := cli.ConfigList(ctx, options)
	log.Println("返回配置的列表  = ", c, err)
}

//ContainerCommit applies changes into a container and creates a new tagged image.
func ContainerCommitRun(cli *client.Client, ctx context.Context) {
	cli.ContainerCommit(ctx, "aaaa", types.ContainerCommitOptions{})
}

//创建一个新的容器
func RunContainer(cli *client.Client, ctx context.Context) {

	//80/tcp
	nats := nat.PortSet{
		"80/tcp": struct{}{},
	}

	//配置健康检查
	healthcheck := &container.HealthConfig{}

	config := &container.Config{
		Hostname:     "man_1", // Hostname
		Domainname:   "man_1", // Domainname
		User:         "",      // User that will run the command(s) inside the container, also support user:group   -> 在容器内运行命令的User，也支持User:group
		AttachStdin:  true,    // Attach the standard input, makes possible user interaction  -> 附加标准输入，实现用户交互
		AttachStdout: true,    // Attach the standard output   ->  附加标准输出
		AttachStderr: true,    // Attach the standard error   ->  附加标准错误
		ExposedPorts: nats,    // List of exposed ports     ->   暴露端口列表
		Tty:          true,    // Attach standard streams to a tty, including stdin if it is not closed.  -> 附加标准流到Tty，包括未关闭的stdin。
		OpenStdin:    true,    // Open stdin    ->  打开stdin
		StdinOnce:    true,    // If true, close stdin after the 1 attached client disconnects.   ->  如果为真，在1个附加的客户端断开连接后关闭stdin。
		//Env            : []string{"GO_MOD=dev"},           // List of environment variable to set in the container    ->  在容器中设置的环境变量的列表
		//Cmd          :   []string{"echo", "mange"},  // Command to run when starting the container     ->  命令在启动容器时运行
		Healthcheck: healthcheck, // Healthcheck describes how to check the container is healthy  ->  描述了如何检查容器是否健康
		//ArgsEscaped     bool                `json:",omitempty"` // True if command is already escaped (meaning treat as a command line) (Windows specific).   -> 如果命令已经被转义(意味着作为命令行处理)(特定于Windows)，则为True。
		Image: "nginx:1.10-alpine", // Name of the image as it was passed by the operator (e.g. could be symbolic)   ->   镜像本地镜像,
		//Volumes   :  map[string]struct{}, // List of volumes (mounts) used for the container    ->  容器使用的卷(挂载)列表
		//WorkingDir      string              // Current directory (PWD) in the command will be launched   ->   当前目录(PWD)命令将在WorkingDir字符串中启动
		//Entrypoint   :   []string{"node", "/workspace/test.js"},  // Entrypoint to run when starting the container     ->  启动容器时运行的StrSlice // Entrypoint
		NetworkDisabled: false, // Is network disabled    ->  是网络禁用
		//MacAddress      string              `json:",omitempty"` // Mac Address of the container   ->  容器的Mac地址
		//OnBuild         []string            // ONBUILD metadata that were defined on the image Dockerfile   ->  在图像Dockerfile中定义的OnBuild元数据
		//Labels          map[string]string   // List of labels set to this container    ->   设置到该容器的标签列表
		//StopSignal      string              `json:",omitempty"` // Signal to stop a container    ->   信号停止一个容器
		//StopTimeout     *int                `json:",omitempty"` // Timeout (in seconds) to stop a container     ->   Timeout(秒)停止容器
		//Shell           strslice.StrSlice   `json:",omitempty"` // Shell for shell-form of RUN, CMD, ENTRYPOINT   ->   Shell为Shell形式的运行，CMD，入口点
	}

	hostConfig := &container.HostConfig{
		Binds: []string{
			"/var/run/docker.sock:/var/run/docker.sock",
		},
		//0.0.0.0:14143->80/tcp
		PortBindings: nat.PortMap{
			nat.Port("80/tcp"): []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "14143",
				},
			},
		},
	}

	containerName := ""

	containerResp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, containerName)
	if err != nil {
		log.Println("err = ", err)
	}
	log.Println(containerResp)

	//启动一个容器
	if err = cli.ContainerStart(ctx, containerResp.ID, types.ContainerStartOptions{}); err != nil {
		log.Println(err)
	}
}

//获取运行容器
//返回docker主机中的容器列表。
func ContainerList(cli *client.Client, ctx context.Context) ([]types.Container, error) {
	options := types.ContainerListOptions{
		//Quiet   bool
		//Size    bool
		All: true, // ps -a
		//Latest  bool
		//Since   string
		//Before  string
		//Limit   int
		//Filters filters.Args
	}
	clist, err := cli.ContainerList(ctx, options)
	log.Println(clist, err)
	return clist, err
}

//func (cli *Client) ContainerDiff(ctx context.Context, containerID string) ([]container.ContainerChangeResponseItem, error)
//ContainerDiff   显示了容器文件系统自启动以来的差异。

//func (cli *Client) ContainerExecAttach(ctx context.Context, execID string, config types.ExecStartCheck) (types.HijackedResponse, error)
//ContainerExecAttach    连接到服务器中的exec进程。它返回一个类型。被劫持的连接与被劫持的连接和读取器进行输出。
//它由被呼叫者通过呼叫type . hijackedresponse . close来关闭被劫持的连接。

//func (cli *Client) ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error)
//ContainerExecCreate     创建一个新的exec配置来运行exec进程。

//func (cli *Client) ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error)
//ContainerExecInspect      返回关于docker主机上的特定exec进程的信息。

//func (cli *Client) ContainerExecResize(ctx context.Context, execID string, options types.ResizeOptions) error
//ContainerExecResize     更改在容器中运行的exec进程的tty的大小。

//func (cli *Client) ContainerExecStart(ctx context.Context, execID string, config types.ExecStartCheck) error
//ContainerExecStart     启动已经在docker主机中创建的exec进程。

//func (cli *Client) ContainerExport(ctx context.Context, containerID string) (io.ReadCloser, error)
//ContainerExport    检索容器的原始内容，并将其作为一个io.ReadCloser返回。应该由调用者来关闭流。

//ContainerInspect   返回容器信息。
func ContainerInspect(cli *client.Client, ctx context.Context) {
	containerID := "79ac5330534a"
	info, err := cli.ContainerInspect(ctx, containerID)
	log.Println(info, err)
	log.Println(info.ContainerJSONBase)
	log.Println(info.Config)
	log.Println(info.NetworkSettings)
}

//ContainerInspectWithRaw   返回容器信息及其原始表示。
func ContainerInspectWithRaw(cli *client.Client, ctx context.Context) {
	containerID := "79ac5330534a"
	info, b, err := cli.ContainerInspectWithRaw(ctx, containerID, true)
	log.Println(info, b, err)
}

// docker kill <cid>
//ContainerKill   终止容器进程，但不从docker主机上删除容器。
func ContainerKill(cli *client.Client, ctx context.Context) {
	containerID := "dc1f0e9834bd"
	err := cli.ContainerKill(ctx, containerID, "")
	log.Println(err)
}

//查看容器日志
//func (cli *Client) ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error)
/*
reader, err := client.ContainerLogs(ctx, "container_id", types.ContainerLogsOptions{})
if err != nil {
    log.Fatal(err)
}
_, err = io.Copy(os.Stdout, reader)
if err != nil && err != io.EOF {
    log.Fatal(err)
}
*/

//ContainerPause暂停给定容器的主进程，但不终止它。
//func (cli *Client) ContainerPause(ctx context.Context, containerID string) error
func ContainerPause(cli *client.Client, ctx context.Context) {
	containerID := "dc1f0e9834bd"
	err := cli.ContainerPause(ctx, containerID)
	log.Println(err)
}

//container remove终止并从docker主机中移除一个容器。
//func (cli *Client) ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error
func ContainerRemove(cli *client.Client, ctx context.Context) {
	containerID := "294ad2db9818"
	options := types.ContainerRemoveOptions{
		//RemoveVolumes bool
		//RemoveLinks   bool
		//Force         bool
	}
	err := cli.ContainerRemove(ctx, containerID, options)
	log.Println(err)
}

//ContainerRename更改给定容器的名称。
func ContainerRename(cli *client.Client, ctx context.Context) {
	containerID := "fdfddca5e7f0"
	newname := "test1"
	err := cli.ContainerRename(ctx, containerID, newname)
	log.Println(err)
}

//ContainerResize更改容器tty的大小。
//func (cli *Client) ContainerResize(ctx context.Context, containerID string, options types.ResizeOptions) error

//ContainerRestart停止并再次启动一个容器。它使守护进程在给定超时的特定时间内等待容器再次启动。
//func (cli *Client) ContainerRestart(ctx context.Context, containerID string, timeout *time.Duration) error
func ContainerRestart(cli *client.Client, ctx context.Context) {
	containerID := "fdfddca5e7f0"
	timeout := 100 * time.Second
	err := cli.ContainerRestart(ctx, containerID, &timeout)
	log.Println(err)
}

//ContainerStart向docker守护进程发送一个请求来启动一个容器。
func ContainerStart(cli *client.Client, ctx context.Context) {
	//启动一个容器
	//if err = cli.ContainerStart(ctx, containerResp.ID, types.ContainerStartOptions{}); err != nil {
	//	log.Println(err)
	//}
}

//ContainerStatPath  返回容器文件系统内路径的统计信息。  ls
func ContainerStatPath(cli *client.Client, ctx context.Context) {
	containerID := "fdfddca5e7f0"
	path := "."
	d, e := cli.ContainerStatPath(ctx, containerID, path)
	log.Println(d, e)
}

//ContainerStats  返回给定容器的接近实时的统计数据。应该由调用者来关闭io。ReadCloser返回。
//func (cli *Client) ContainerStats(ctx context.Context, containerID string, stream bool) (types.ContainerStats, error)

//容器statsoneshot从容器获取单个stat条目。它与“ContainerStats”的不同之处在于，API不应该等待初始状态
//func (cli *Client) ContainerStatsOneShot(ctx context.Context, containerID string) (types.ContainerStats, error)

//ContainerStop停止容器
func ContainerStop(cli *client.Client, ctx context.Context) {
	containerID := "fdfddca5e7f0"
	timeout := 100 * time.Second
	err := cli.ContainerStop(ctx, containerID, &timeout)
	log.Println(err)
}

//ContainerTop 显示容器内的流程信息。
func ContainerTop(cli *client.Client, ctx context.Context) {
	containerID := "79ac5330534a"
	arguments := []string{}
	d, e := cli.ContainerTop(ctx, containerID, arguments)
	log.Println(d, e)
}

//ContainerUnpause 在容器内恢复进程执行
func ContainerUnpause(cli *client.Client, ctx context.Context) {
	containerID := "79ac5330534a"
	err := cli.ContainerUnpause(ctx, containerID)
	log.Println(err)
}

//ContainerUpdate  更新容器的资源
func ContainerUpdate(cli *client.Client, ctx context.Context) {
	containerID := "79ac5330534a"
	updateConfig := container.UpdateConfig{}
	d, e := cli.ContainerUpdate(ctx, containerID, updateConfig)
	log.Println(d, e)
}

//ContainerWait   将一直等待，直到指定的容器处于给定条件所指示的特定状态，即“不运行”(默认)、“下一退出”或“已删除”。
//如果该客户机的API版本在1.30之前，则忽略条件
func ContainerWait(cli *client.Client, ctx context.Context) {
	_, errC := cli.ContainerWait(ctx, "container_id", "")
	if err := <-errC; err != nil {
		log.Fatal(err)
	}
}

//ContainersPrune   请求守护进程删除未使用的数据
func ContainersPrune(cli *client.Client, ctx context.Context) {
	pruneFilters := filters.Args{}
	d, e := cli.ContainersPrune(ctx, pruneFilters)
	log.Println(d, e)
}

//CopyFromContainer   从容器获取内容，并将其作为TAR存档的读取器返回，以便在主机中对其进行操作。要由打电话的人来结束阅读。
func CopyFromContainer(cli *client.Client, ctx context.Context) {
	containerID := "79ac5330534a"
	i, d, e := cli.CopyFromContainer(ctx, containerID, "/")
	log.Println(i, d, e)
}

//CopyToContainer  将内容复制到容器文件系统中。注意，“content”必须是TAR存档的阅读器
func CopyToContainer(cli *client.Client, ctx context.Context) {
	containerID := "79ac5330534a"
	dstPath := "/"
	content := new(bytes.Buffer)
	options := types.CopyToContainerOptions{}
	e := cli.CopyToContainer(ctx, containerID, dstPath, content, options)
	log.Println(e)
}

//CustomHTTPHeaders  返回客户端存储的自定义http头。
func CustomHTTPHeaders(cli *client.Client, ctx context.Context) {
	d := cli.CustomHTTPHeaders()
	log.Println(d)
}

//DaemonHost  返回客户端使用的主机地址
func DaemonHost(cli *client.Client, ctx context.Context) {
	d := cli.DaemonHost()
	log.Println(d)
}

//拨号接口返回一个被劫持的连接，该连接带有协商协议原型。
//func (cli *Client) DialHijack(ctx context.Context, url, proto string, meta map[string][]string) (net.Conn, error)

//为原始流连接返回一个拨号器，带有HTTP/1.1报头，可用于代理守护进程连接。由“docker dial-stdio”使用(docker/cli#889)。
//func (cli *Client) Dialer() func(context.Context) (net.Conn, error)

//DiskUsage  从守护进程请求当前数据使用情况
func DiskUsage(cli *client.Client, ctx context.Context) {
	d, e := cli.DiskUsage(ctx)
	log.Println(d, e)
}

//DistributionInspect   返回带有完整清单的镜像摘要
func DistributionInspect(cli *client.Client, ctx context.Context) {
	d, e := cli.DistributionInspect(ctx, "nginx:1.10-alpine", "")
	log.Println(d, e)
}

//Events 返回守护进程中的事件流。
//func (cli *Client) Events(ctx context.Context, options types.EventsOptions) (<-chan events.Message, <-chan error)

//HTTPClient  返回绑定到服务器的HTTP客户机的副本
func HTTPClient(cli *client.Client) {
	c := cli.HTTPClient()
	log.Println(c)
}

//ImageBuild  向守护进程发送构建映像的请求。响应的主体实现一个io。ReadCloser，它是由来电者关闭它。
func SDKImageBuild(cli *client.Client, ctx context.Context) {
	buildContext := new(bytes.Buffer)
	options := types.ImageBuildOptions{}
	d, e := cli.ImageBuild(ctx, buildContext, options)
	log.Println(d, e)
}

//ImageCreate  创建一个基于父选项的新图像。它返回响应主体中的JSON内容。
func SDKImageCreate(cli *client.Client, ctx context.Context) {
	options := types.ImageCreateOptions{
		RegistryAuth: "--username=samsong@1160745187615739 --password=tco99312^ ",
		Platform:     "registry.cn-shenzhen.aliyuncs.com",
	}
	d, e := cli.ImageCreate(ctx, "", options)
	log.Println(d, e)
}

//ImageHistory  以历史格式返回图像中的更改。
func ImageHistory(cli *client.Client, ctx context.Context) {
	imageID := ""
	d, e := cli.ImageHistory(ctx, imageID)
	log.Println(d, e)
}

//ImageImport  根据源选项创建一个新映像。它返回响应主体中的JSON内容。
func ImageImport(cli *client.Client, ctx context.Context) {
	source := types.ImageImportSource{}
	options := types.ImageImportOptions{}
	cli.ImageImport(ctx, source, "", options)
}

//ImageInspectWithRaw  返回图像信息及其原始表示。
func ImageInspectWithRaw(cli *client.Client, ctx context.Context) {
	imageID := "bf756fb1ae65"
	i, d, e := cli.ImageInspectWithRaw(ctx, imageID)
	log.Println(i, d, e)
}

//ImageLoad  加载图像的码头工人从客户端主机主机。应该由调用者来关闭io。此函数返回的ImageLoadResponse中的ReadCloser。
//func (cli *Client) ImageLoad(ctx context.Context, input io.Reader, quiet bool) (types.ImageLoadResponse, error)

//拉镜像
//ImagePull  请求docker主机从远程注册表中提取图像。如果操作未被授权，它将执行特权函数，然后再尝试一次。由调用者来处理io。仔细阅读，然后合上。
func ImagePull(cli *client.Client, ctx context.Context) {
	options := types.ImagePullOptions{
		All:          true, //会将所有的版本拉下来
		RegistryAuth: "",
		//PrivilegeFunc RequestPrivilegeFunc
		Platform: "",
	}
	imageName := "busybox:latest"
	events, err := cli.ImagePull(ctx, imageName, options)
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(events)
	type Event struct {
		Status         string `json:"status"`
		Error          string `json:"error"`
		Progress       string `json:"progress"`
		ProgressDetail struct {
			Current int `json:"current"`
			Total   int `json:"total"`
		} `json:"progressDetail"`
	}
	var event *Event
	for {
		if err := d.Decode(&event); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Printf("EVENT: %+v\n", event)
	}
	// Latest event for new image
	// EVENT: {Status:Status: Downloaded newer image for busybox:latest Error: Progress:[==================================================>]  699.2kB/699.2kB ProgressDetail:{Current:699243 Total:699243}}
	// Latest event for up-to-date image
	// EVENT: {Status:Status: Image is up to date for busybox:latest Error: Progress: ProgressDetail:{Current:0 Total:0}}
	if event != nil {
		if strings.Contains(event.Status, fmt.Sprintf("Downloaded newer image for %s", imageName)) {
			// new
			fmt.Println("new")
		}
		if strings.Contains(event.Status, fmt.Sprintf("Image is up to date for %s", imageName)) {
			// up-to-date
			fmt.Println("up-to-date")
		}
	}
}

// 推镜像
//ImagePush   请求docker主机将一个映像推送到远程注册表。如果操作未被授权，它将执行特权函数，然后再尝试一次。由调用者来处理io。
func ImagePush(cli *client.Client, ctx context.Context) {
	tag := "nginx"
	resp, err := cli.ImagePush(ctx, tag, types.ImagePushOptions{
		RegistryAuth: `{ \"username\": \"myusername\", \"password\": \"mypassword\", \"email\": \"myemail\" }`,
	})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, resp)
	resp.Close()

}

//删除镜像
//ImageRemove  从docker主机上删除一个图像。
func ImageRemove(cli *client.Client, ctx context.Context) {
	imageID := "a84c36ecc374"
	options := types.ImageRemoveOptions{
		Force:         true,
		PruneChildren: true,
	}
	d, e := cli.ImageRemove(ctx, imageID, options)
	log.Println(d, e)
}

//ImageSave  从docker主机中检索一个或多个图像作为io.ReadCloser。存储图像和关闭流由调用者决定。
func ImageSave(cli *client.Client, ctx context.Context) {
	cli.ImageSave(ctx, []string{"a84c36ecc374"})
}

//搜索镜像
//ImageSearch  使docker主机通过远程注册表中的术语进行搜索。结果列表没有以任何方式排序。
func SDKImageSearch(cli *client.Client, ctx context.Context) {
	term := "java"
	options := types.ImageSearchOptions{
		Limit: 100,
	}
	d, e := cli.ImageSearch(ctx, term, options)
	log.Println(d, e)
}

//ImageTag  标记docker主机中的一个图像
func ImageTag(cli *client.Client, ctx context.Context) {

	e := cli.ImageTag(ctx, "busybox:1.24.0", "new_busybox:latest")
	log.Println(e)
}

//ImagesPrune  请求守护进程删除未使用的数据
//func (cli *Client) ImagesPrune(ctx context.Context, pruneFilters filters.Args) (types.ImagesPruneReport, error)

//Info  返回关于docker服务器的信息。
func Info(cli *client.Client, ctx context.Context) {
	d, e := cli.Info(ctx)
	log.Println(d, e)
}

//NegotiateAPIVersion  查询API并更新版本以匹配API版本。任何错误都会被悄悄地忽略。如果有手动覆盖，
//无论是通过“DOCKER_API_VERSION”环境变量，还是用固定版本(“options . withversion (xx)”)初始化客户端，都不会执行任何协商。
func NegotiateAPIVersion(cli *client.Client, ctx context.Context) {
	cli.NegotiateAPIVersion(ctx)
}

//NegotiateAPIVersionPing  更新客户端版本以匹配Ping。如果ping版本小于默认版本，则调用APIVersion。如果有手动覆盖，
//无论是通过“DOCKER_API_VERSION”环境变量，还是用固定版本(“options . withversion (xx)”)初始化客户端，都不执行协商。
func NegotiateAPIVersionPing(cli *client.Client, ctx context.Context) {
	p := types.Ping{
		//APIVersion     string
		//OSType         string
		//Experimental   bool
		//BuilderVersion BuilderVersion
	}
	cli.NegotiateAPIVersionPing(p)
}

//NetworkConnect  将容器连接到docker主机中的现有网络。
//func (cli *Client) NetworkConnect(ctx context.Context, networkID, containerID string, config *network.EndpointSettings) error

//NetworkCreate  在docker主机中创建一个新的网络。
//func (cli *Client) NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error)

//NetworkDisconnect  将容器与docker主机中的现有网络断开连接。
//func (cli *Client) NetworkDisconnect(ctx context.Context, networkID, containerID string, force bool) error

//NetworkInspect  返回在docker主机中配置的特定网络的信息。
//func (cli *Client) NetworkInspect(ctx context.Context, networkID string, options types.NetworkInspectOptions) (types.NetworkResource, error)

//NetworkInspectWithRaw  返回在docker主机中配置的特定网络的信息及其原始表示。
//func (cli *Client) NetworkInspectWithRaw(ctx context.Context, networkID string, options types.NetworkInspectOptions) (types.NetworkResource, []byte, error)

//NetworkList  返回docker主机中配置的网络列表。
func NetworkList(cli *client.Client, ctx context.Context) {
	options := types.NetworkListOptions{
		//Filters filters.Args
	}
	d, e := cli.NetworkList(ctx, options)
	log.Println(d, e)
}

//NetworkRemove  从docker主机上删除一个现有的网络。
//func (cli *Client) NetworkRemove(ctx context.Context, networkID string) error

//NetworksPrune  请求守护进程删除未使用的网络
//func (cli *Client) NetworksPrune(ctx context.Context, pruneFilters filters.Args) (types.NetworksPruneReport, error)

//func (cli *Client) NewVersionError(APIrequired, feature string) error
//NewVersionError 如果所需的APIVersion小于当前支持的版本，则NewVersionError返回一个错误

//func (cli *Client) NodeInspectWithRaw(ctx context.Context, nodeID string) (swarm.Node, []byte, error)
//NodeInspectWithRaw 返回节点信息。
func NodeInspectWithRaw(cli *client.Client, ctx context.Context) {
	i, d, e := cli.NodeInspectWithRaw(ctx, "")
	log.Println(i, d, e)
}

//NodeList  返回节点列表。
func NodeList(cli *client.Client, ctx context.Context) {
	options := types.NodeListOptions{}
	d, e := cli.NodeList(ctx, options)
	log.Println(d, e)
}

//func (cli *Client) NodeRemove(ctx context.Context, nodeID string, options types.NodeRemoveOptions) error
//NodeRemove 删除节点。

//func (cli *Client) NodeUpdate(ctx context.Context, nodeID string, version swarm.Version, node swarm.NodeSpec) error
//NodeUpdate 更新一个节点

//func (cli *Client) Ping(ctx context.Context) (types.Ping, error)
//Ping ping服务器并返回“Docker-Experimental”、“build - version”、“OS-Type”和“API-Version”头文件的值。
//它尝试在端点上使用HEAD请求，但如果守护进程不支持HEAD，则返回到GET。
func Ping(cli *client.Client, ctx context.Context) {
	d, e := cli.Ping(ctx)
	log.Println(d, e)
}

//func (cli *Client) PluginCreate(ctx context.Context, createContext io.Reader, createOptions types.PluginCreateOptions) error
//PluginCreate 创建一个插件

//func (cli *Client) PluginDisable(ctx context.Context, name string, options types.PluginDisableOptions) error
//PluginDisable 禁用插件

//func (cli *Client) PluginEnable(ctx context.Context, name string, options types.PluginEnableOptions) error
//PluginEnable 使一个插件

//func (cli *Client) PluginInspectWithRaw(ctx context.Context, name string) (*types.Plugin, []byte, error)
//PluginInspectWithRaw 检查现有插件
func PluginInspectWithRaw(cli *client.Client, ctx context.Context) {
	i, d, e := cli.PluginInspectWithRaw(ctx, "")
	log.Println(i, d, e)
}

//func (cli *Client) PluginInstall(ctx context.Context, name string, options types.PluginInstallOptions) (rc io.ReadCloser, err error)
//PluginInstall 安装插件

//func (cli *Client) PluginList(ctx context.Context, filter filters.Args) (types.PluginsListResponse, error)
//PluginList 返回已安装的插件
func PluginList(cli *client.Client, ctx context.Context) {
	filter := filters.Args{}
	d, e := cli.PluginList(ctx, filter)
	log.Println(d, e)
}

//func (cli *Client) PluginPush(ctx context.Context, name string, registryAuth string) (io.ReadCloser, error)
//PluginPush 推一个插件到注册表

//func (cli *Client) PluginRemove(ctx context.Context, name string, options types.PluginRemoveOptions) error
//PluginRemove removes a plugin

//func (cli *Client) PluginSet(ctx context.Context, name string, args []string) error
//PluginSet 修改现有插件的设置

//func (cli *Client) PluginUpgrade(ctx context.Context, name string, options types.PluginInstallOptions) (rc io.ReadCloser, err error)
//PluginUpgrade upgrades a plugin

//远程私有库登录
//func (cli *Client) RegistryLogin(ctx context.Context, auth types.AuthConfig) (registry.AuthenticateOKBody, error)
//RegistryLogin 使用给定的docker注册表验证docker服务器。当身份验证失败时，它返回unauthorizedError。
func RegistryLogin(cli *client.Client, ctx context.Context) {
	auth := types.AuthConfig{
		Username: "samsong@1160745187615739",
		Password: "tco99312^",
		Auth:     "",
		//// Email is an optional value associated with the username.
		//// This field is deprecated and will be removed in a later
		//// version of docker.
		//Email string `json:"email,omitempty"`
		ServerAddress: "registry.cn-shenzhen.aliyuncs.com",
		//// IdentityToken is used to authenticate the user and get
		//// an access token for the registry.
		//IdentityToken string `json:"identitytoken,omitempty"`
		//// RegistryToken is a bearer token to be sent to a registry
		//RegistryToken string `json:"registrytoken,omitempty"`
	}
	d, e := cli.RegistryLogin(ctx, auth)
	log.Println(d, e)
}

//func (cli *Client) SecretCreate(ctx context.Context, secret swarm.SecretSpec) (types.SecretCreateResponse, error)
//SecretCreate creates a new Secret.

//func (cli *Client) SecretInspectWithRaw(ctx context.Context, id string) (swarm.Secret, []byte, error)
//SecretInspectWithRaw 返回带有原始数据的秘密信息

//func (cli *Client) SecretList(ctx context.Context, options types.SecretListOptions) ([]swarm.Secret, error)
//SecretList returns the list of secrets.
func SecretList(cli *client.Client, ctx context.Context) {
	options := types.SecretListOptions{}
	d, e := cli.SecretList(ctx, options)
	log.Println(d, e)
}

//
