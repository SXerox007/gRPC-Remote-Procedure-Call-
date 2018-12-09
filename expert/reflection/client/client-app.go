package main

import (
	"context"
	"io"
	"log"
	"time"

	"gRPC-Remote-Procedure-Call-/expert/reflection/proto"

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

	msg := chat.NewChatServiceContainerClient(client)

	stream, err := msg.ChatService(context.Background())

	if err != nil {
		log.Fatal("Some Error Occured: ", err)
		return
	}

	ClientStream(stream)
}

func ClientStream(stream chat.ChatServiceContainer_ChatServiceClient) {
	req := []*chat.ChatServiceRequest{
		&chat.ChatServiceRequest{
			Message: "Hi there",
		},
		&chat.ChatServiceRequest{
			Message: "yolo",
		},
		&chat.ChatServiceRequest{
			Message: "Hahaha",
		},
		&chat.ChatServiceRequest{
			Message: "Nope",
		},
		&chat.ChatServiceRequest{
			Message: "Ammmmmm",
		},
	}

	waitc := make(chan struct{})

	go func() {
		//  send multiple messages
		for _, req := range req {
			stream.Send(req)
			time.Sleep(2000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	//Receive message
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error int receiving: %v", err)
				break
			}
			log.Println("Data Success: ", res.GetMessage())
		}
		close(waitc)
	}()

	//wait until everything is complete
	<-waitc
}
