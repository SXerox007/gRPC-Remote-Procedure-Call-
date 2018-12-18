package main

import (
	"gRPC-Remote-Procedure-Call-/oauth2/proto"

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
	client, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error in connection : ", err)
	}
	defer client.Close()

	msg := oauthpb.NewOAuthServiceClient(client)

	log.Println("Final Data ", msg)
}
