// docker SDK 对应的功能实现
// 这里我简述为 通过 docker golang SDK 操作 docker

package slve

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

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		log.Println(err)
		return
	}

	for _, image := range images {
		log.Println(&image)
	}
}