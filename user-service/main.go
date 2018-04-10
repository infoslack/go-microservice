package main

import (
	"fmt"
	"log"

	pb "github.com/infoslack/go-microservice/user-service/proto/user"
	micro "github.com/micro/go-micro"
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

	// Create a new service
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
