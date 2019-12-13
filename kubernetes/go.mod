module micros/kubernetes

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.11
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.10.0
	k8s.io/api v0.17.0
	k8s.io/client-go v11.0.0+incompatible // indirect
	k8s.io/utils v0.0.0-20191114200735-6ca3b61696b6 // indirect
)

replace (
	github.com/golang/lint => golang.org/x/lint v0.0.0-20190930215403-16217165b5de
	github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.9
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90
	micros/kubernetes => ../kubernetes
)
