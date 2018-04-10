package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/infoslack/go-microservice/consignment-service/proto/consignment"
	vesselProto "github.com/infoslack/go-microservice/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
)

const defaultHost = "localhost:27017"

func main() {

	// Database host from env var
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	defer session.Close()

	if err != nil {
		log.Panicf("Couldn't connect to datastore with host %s - %v", host, err)
	}

	// Create a new service
	srv := micro.NewService(

		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
