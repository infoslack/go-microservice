package main

import (
	"fmt"
	"log"

	pb "github.com/infoslack/go-microservice/user-service/proto/auth"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/mdns"
)

func main() {

	// Create a database connection
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Couldn't connect to DB: %v", err)
	}

	// Automatically migrates the user struct into DB
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	// Create a new service
	srv := micro.NewService(
		micro.Name("auth"),
		micro.Version("latest"),
	)

	srv.Init()

	publisher := micro.NewPublisher("user.created", srv.Client())

	pb.RegisterAuthHandler(srv.Server(), &service{repo, tokenService, publisher})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
