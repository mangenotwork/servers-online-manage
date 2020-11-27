module github.com/mangenotwork/servers-online-manage/slve

go 1.13

replace (
	github.com/docker/docker v1.13.1 => github.com/docker/engine v17.12.0-ce-rc1.0.20200204220554-5f6d6f3f2203+incompatible
	github.com/mangenotwork/servers-online-manage => ../../
	github.com/mangenotwork/servers-online-manage/lib => ../lib
	github.com/mangenotwork/servers-online-manage/slve => ./
)

require (
	github.com/Microsoft/go-winio v0.4.15 // indirect
	github.com/containerd/containerd v1.4.2 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/mangenotwork/servers-online-manage/lib v0.0.0-00010101000000-000000000000
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.7.0 // indirect
	google.golang.org/grpc v1.33.2 // indirect
)
