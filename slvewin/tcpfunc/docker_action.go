// docker 交互的路由
// 优先使用 docker sdk进行交互， 其次使用 cmd进行交互, 最后使用 docker Remote api 进行交互；

package tcpfunc

import (
	"github.com/mangenotwork/servers-online-manage/lib/docker"
)

func Images(action string) (data []byte, err error){
	switch action {
	case "get_images_list":
		return docker.ImageList()
	}
	return
}