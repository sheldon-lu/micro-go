module micro-go/user

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.11
	github.com/micro/go-micro v1.10.0
	github.com/nats-io/nats-server/v2 v2.1.2 // indirect
	golang.org/x/crypto v0.0.0-20191108234033-bd318be0434a
	micro-go/common v0.0.0-00010101000000-000000000000
)

replace (
	github.com/golang/lint => golang.org/x/lint v0.0.0-20190930215403-16217165b5de
	github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.9
	micro-go/common => ../common
	micro-go/user/model/user => ../user/model/user
)
