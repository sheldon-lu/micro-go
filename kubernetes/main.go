package main

import (
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/util/log"
    "micro-go/common/database"
    "micro-go/kubernetes/handler"
    `micro-go/kubernetes/model/k8smodel`

    kubernetes "micro-go/kubernetes/proto/kubernetes"
)

func main() {
    db, err := database.CreateConnection()
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

    repo := &k8smodel.KubeModel{db}

    // New Service
    service := micro.NewService(
        micro.Name("go.micro.srv.kubernetes"),
        micro.Version("latest"),
    )

    // Initialise service
    service.Init()

    // Register Handler
    kubernetes.RegisterKubernetesHandler(service.Server(), &handler.Kubernetes{Repo: repo})

    // Register Struct as Subscriber
    // micro.RegisterSubscriber("go.micro.srv.kubernetes", service.Server(), new(subscriber.Kubernetes))

    // Register Function as Subscriber
    // micro.RegisterSubscriber("go.micro.srv.kubernetes", service.Server(), subscriber.Handler)

    // Run service
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
