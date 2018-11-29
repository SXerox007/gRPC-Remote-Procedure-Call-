package main

import (
	"context"
	"gRPC-Remote-Procedure-Call-/basic/unary-API/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	Init()
}

func Init() {
	ClientSetup()
}

func ClientSetup() {
	//we don't have ssl certificate right now so we use grpc.withInsecure
	client, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error in connection : ", err)
	}
	defer client.Close()

	msg := search.NewCreateServiceClient(client)

	req := &search.SearchRequest{
		PhoneNumber: 1234567890,
	}
	data, err := msg.SearchContact(context.Background(), req)

	if err != nil {
		log.Fatal("Some Error Occured: ", err)
	}

	log.Println("Response Send by Server:", data)
}
