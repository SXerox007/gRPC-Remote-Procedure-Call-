package grpcclient

import (
	"fmt"
	"gRPC-Remote-Procedure-Call-/slack-bot/log"
	"gRPC-Remote-Procedure-Call-/slack-bot/proto"

	"google.golang.org/grpc"
)

var sc slackbot.SlackBotServiceClient

func InitGRPCClient() {
	client, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Error.Fatal("Error in connection : ", err)
	}
	defer client.Close()

	sc = slackbot.NewSlackBotServiceClient(client)
	fmt.Println("Service Client:", sc)

}

func GetServiceClient() slackbot.SlackBotServiceClient {
	fmt.Println("Service Client:", sc)
	return sc
}

// client, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatal("Error in connection : ", err)
// 	}
// 	defer client.Close()

// 	msg := informaticapb.NewInformaticaServiceClient(client)

// 	//create the informatica
// 	CreateInformatica(msg)
