package main

import (
    `github.com/gin-gonic/gin`
    "github.com/micro/go-micro/util/log"
    "github.com/micro/go-micro/web"
    "micro-go/kubernetes-web/handler"
)

func main() {
    // create new web service
    service := web.NewService(
        web.Name("lu.micro.web.v1.k8s"),
        web.Version("latest"),
        web.Address(":8881"),
    )

    // initialise service
    if err := service.Init(); err != nil {
        log.Fatal(err)
    }

    // register html handler
    // service.Handle("/", http.FileServer(http.Dir("html")))

    // register call handler
    kubes := new(handler.KuBernetes)
    router := gin.Default()
    v1 := router.Group("/v1")
    {
        v1.GET("/k8s/list", func(c *gin.Context) {
            c.JSON(200, "hi")
        })
        v1.POST("/k8s/detail", kubes.KubeDetail)
        v1.POST("/k8s/apps", kubes.Apps)
        v1.POST("/k8s/updates", kubes.UpdateS)
    }

    service.Handle("/", router)

    // run service
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
