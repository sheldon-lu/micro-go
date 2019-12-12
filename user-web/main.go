package main

import (
    "github.com/micro/go-micro/util/log"
    "github.com/micro/go-micro/web"
    "github.com/gin-gonic/gin"
    "micro-go/user-web/handler"
    "github.com/micro/cli"
)

func main() {
    // create new web service
    service := web.NewService(
        web.Name("lu.micro.web.user"),
        web.Version("latest"),
    )

    // initialise service
    if err := service.Init(
        web.Action(
            func(c *cli.Context) {
                // 初始化handler
                handler.Init()
            }),
    ); err != nil {
        log.Fatal(err)
    }

    // register html handler
    // service.Handle("/", http.FileServer(http.Dir("html")))

    // register call handler
    // Create RESTful handler (using Gin)
    // Init new(User)
    user := new(handler.User)
    router := gin.Default()
    router.GET("/user", user.List)
    router.POST("/user/login", user.Login)
    router.POST("/user/register", user.Register)
    router.PUT("/user/updatewd", user.UpdatePassword)

    service.Handle("/", router)

    // run service
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
