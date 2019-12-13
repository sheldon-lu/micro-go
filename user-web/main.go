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
        web.Name("lu.micro.web.v1.user"),
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
    user := new(handler.Users)
    router := gin.Default()
    v1 := router.Group("/v1")
    {
        v1.GET("/user", user.List)
        v1.POST("/user/login", user.Login)
        v1.POST("/user/register", user.Register)
        v1.PUT("/user/updatewd", user.UpdatePassword)
        v1.DELETE("/user/deleteuser", user.DeleUser)
    }


    service.Handle("/", router)

    // run service
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
