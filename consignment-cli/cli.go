package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/infoslack/go-microservice/consignment-service/proto/consignment"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {

	cmd.Init()

	cl := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Couldn't parse file: %v", err)
	}

	r, err := cl.CreateConsignment(context.TODO(), consignment)
	if err != nil {
		log.Fatalf("Couldn't create: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := cl.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Couldn't list consignments: %v", err)
	}

	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
