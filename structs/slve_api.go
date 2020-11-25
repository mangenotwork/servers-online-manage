package structs

//给slve设置一个名称Post传参结构
type SetSlveNamePost struct {
	Key string `json:"key"`
	Token string `json:"slve_token"`
	NewName string `json:"new_name"`
}
