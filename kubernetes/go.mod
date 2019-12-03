module micros/kubernetes

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.11
	github.com/micro/go-config v1.1.0
	github.com/micro/go-grpc v1.0.1
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.1.0
)

replace (
	github.com/golang/lint => golang.org/x/lint v0.0.0-20190930215403-16217165b5de
	github.com/micro/go-micro => github.com/micro/go-micro v1.1.0
	github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.9
	micros/kubernetes => ../kubernetes
)
