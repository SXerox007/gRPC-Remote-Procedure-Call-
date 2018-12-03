package main

import (
	"context"
	"gRPC-Remote-Procedure-Call-/basic/unary-API/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	Init()
}

func Init() {
	ClientSetup()
}

func ClientSetup() {
	//with ssl
	certFile := "expert/ssl-security/ssl/ca.crt"
	//override of server name the blank arrgument
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	if sslErr != nil {
		log.Fatalln("Some Error occured: ", sslErr)
		return
	}

	client, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(creds))
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
