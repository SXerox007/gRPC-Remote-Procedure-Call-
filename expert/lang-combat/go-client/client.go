package main

import (
	"context"
	"log"

	"gRPC-Remote-Procedure-Call-/expert/lang-combat/protos/login"

	"google.golang.org/grpc"
)

func main() {
	Init()
}

func Init() {
	ClientSetup()
}

func ClientSetup() {
	client, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error in connection : ", err)
	}
	defer client.Close()

	msg := login.NewLoginServiceClient(client)

	//Add block
	Login(msg)

}

func Login(msg login.LoginServiceClient) {
	req := &login.LoginRequest{
		LoginRequestData: &login.LoginRequestData{
			EmailOrPhone: "sumit@gmail.com",
			Password:     "12345678",
		},
	}
	res, err := msg.Login(context.Background(), req)
	if err == nil {
		log.Println("Success:", res)
	} else {
		log.Println("Error: ", err)
	}

}
