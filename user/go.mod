module micros/user

go 1.13

require (
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/golang/protobuf v1.3.1
	github.com/jinzhu/gorm v1.9.11
	github.com/micro/go-config v1.1.0
	github.com/micro/go-grpc v1.0.0
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.1.0
	github.com/softlayer/softlayer-go v0.0.0-20190107163317-a40f6fdd659f
	golang.org/x/crypto v0.0.0-20190404164418-38d8ce5564a5
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
	micros/common v0.0.0-00010101000000-000000000000 // indirect
	sigs.k8s.io/structured-merge-diff v0.0.0-20190302045857-e85c7b244fd2 // indirect
)

replace (
	micros/common => ../common
	micros/product => ../product
	micros/user => ../user
)

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.9

replace github.com/golang/lint => golang.org/x/lint v0.0.0-20190930215403-16217165b5de
