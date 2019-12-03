package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"micros/kubernetes/handler"
	"micros/kubernetes/subscriber"

	kubernetes "micros/kubernetes/proto/kubernetes"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.kubernetes"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	kubernetes.RegisterKubernetesHandler(service.Server(), new(handler.Kubernetes))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.kubernetes", service.Server(), new(subscriber.Kubernetes))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.kubernetes", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
