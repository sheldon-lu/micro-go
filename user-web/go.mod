module micro-go/user-web

go 1.13

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/micro/go-micro v1.10.0
	micro-go/user v0.0.0-00010101000000-000000000000
)

replace (
	micro-go/common => ../common
	micro-go/user => ../user
)
