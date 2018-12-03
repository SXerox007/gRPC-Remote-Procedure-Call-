package main

import (
	"context"
	"gRPC-Remote-Procedure-Call-/expert/deadline/proto"
	"log"
	"time"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

func main() {
	Init()
}

func Init() {
	ClientSetup()
}

//deadline is very important usecase gRPC recommend us to use deadline in every RPC service we use
func ClientSetup() {
	//we don't have ssl certificate right now so we use grpc.withInsecure
	client, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error in connection : ", err)
	}
	defer client.Close()

	msg := errorpb.NewDeadlineHandlingServiceClient(client)

	req := &errorpb.EvenNumberRequest{
		Number: 1,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	data, err := msg.DeadlineHandling(ctx, req)

	if err != nil {
		status, err := status.FromError(err)
		if status.Code() == codes.DeadlineExceeded {
			log.Fatal("Time limit Exceded: ", status)
		} else {
			log.Fatal("Some Error Occured: ", status, err)
		}
	}

	log.Println("Response Send by Server:", data)
}
