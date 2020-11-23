package slve

func Images(action string) (data []byte){
	switch action {
	case "get_images_list":
		return ImageList()
	}
	return []byte{}
}