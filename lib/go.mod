module github.com/mangenotwork/servers-online-manage/lib

go 1.13

replace (
	github.com/docker/docker v1.13.1 => github.com/docker/engine v17.12.0-ce-rc1.0.20200204220554-5f6d6f3f2203+incompatible
	github.com/mangenotwork/servers-online-manage/lib => ./
)

require github.com/jinzhu/gorm v1.9.16
