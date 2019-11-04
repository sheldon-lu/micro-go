package main

import (
	"micros/user/handler"
	"micros/user/model"
	"micros/user/repository"

	"github.com/micro/go-log"

	"github.com/micro/go-micro"
	// "github.com/micro/go-micro/util/log"

	user "shopping/user/proto/user"
)

func main() {

	db, err := model.CreateConnection()
	defer db.Close()

	db.AutoMigrate(&model.User{})

	if err != nil {
		log.Fatalf("connection error : %v \n", err)
	}

	repo := &repository.User{db}

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserServiceHandler(service.Server(), &handler.User{repo})

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}