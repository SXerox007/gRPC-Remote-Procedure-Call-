package main

import (
	"context"
	"gRPC-Remote-Procedure-Call-/expert/error-handling/proto"
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

	msg := errorpb.NewErrorHandlingServiceClient(client)

	req := &errorpb.EvenNumberRequest{
		Number: 1,
	}
	data, err := msg.ErrorHandlingService(context.Background(), req)

	if err != nil {
		log.Fatal("Some Error Occured: ", err)
	}

	log.Println("Response Send by Server:", data)
}
