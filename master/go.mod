module github.com/mangenotwork/servers-online-manage/master

go 1.13

replace (
	github.com/mangenotwork/servers-online-manage => ../../
	github.com/mangenotwork/servers-online-manage/lib => ../lib
	github.com/mangenotwork/servers-online-manage/master => ./
)

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/mangenotwork/servers-online-manage/lib v0.0.0-00010101000000-000000000000
)
