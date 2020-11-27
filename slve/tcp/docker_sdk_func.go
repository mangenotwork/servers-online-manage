// docker SDK 对应的功能实现
// 这里我简述为 通过 docker golang SDK 操作 docker

package tcp

import (
	"log"
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Run1(){
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	//ImagesRun(cli,ctx)
	//GetClientVersion(cli,ctx)
	GetConfigList(cli,ctx)
}

//docker images
func ImagesRun(cli *client.Client, ctx context.Context){
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		log.Println(err)
		return
	}

	for _, image := range images {
		log.Println(&image)
	}
}

//返回此客户机使用的API版本。
func GetClientVersion(cli *client.Client,  ctx context.Context){
	version := cli.ClientVersion()
	log.Println("回此客户机使用的API版本 = ", version)
}

//关闭客户端使用的传输
func ClientClose(cli *client.Client,  ctx context.Context){
	err := cli.Close()
	log.Println("关闭客户端使用的传输, err = ", err)
}

//返回配置的列表
func GetConfigList(cli *client.Client,  ctx context.Context){
	options := types.ConfigListOptions{}
	c , err := cli.ConfigList(ctx, options)
	log.Println("返回配置的列表  = ", c , err)
}

//ContainerCommit applies changes into a container and creates a new tagged image.
func ContainerCommitRun(cli *client.Client,  ctx context.Context){
	cli.ContainerCommit(ctx, "aaaa", types.ContainerCommitOptions{})
}

//创建一个新的容器
