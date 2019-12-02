module micros/common

go 1.13

require (
	github.com/jinzhu/gorm v1.9.11
	github.com/micro/go-config v1.1.0
	github.com/micro/go-log v0.1.0
)

replace (
	micros/common => ../common
	micros/product => ../product
	micros/user => ../user
)

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.9

replace github.com/golang/lint => golang.org/x/lint v0.0.0-20190930215403-16217165b5de