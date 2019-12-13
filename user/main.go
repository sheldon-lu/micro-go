package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"micro-go/common/database"
	"micro-go/user/handler"
	"micro-go/user/model"
	users "micro-go/user/model/user"
	user "micro-go/user/proto/user"
)

func main() {
	// database
	db, err := database.CreateConnection()
	defer db.Close()

	db.AutoMigrate(&model.User{})

	if err != nil {
		log.Fatalf("connection error : %v \n", err)
	}

	// Get db client


	repo := &users.User{db}

	// New Service
	service := micro.NewService(
		micro.Name("lu.micro.srv.user"),
		micro.Version("latest"),
	)

	// New Service

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), &handler.User{repo})

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("lu.micro.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("lu.micro.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
