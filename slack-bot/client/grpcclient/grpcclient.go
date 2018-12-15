package grpcclient

import (
	"fmt"
	"gRPC-Remote-Procedure-Call-/slack-bot/log"
	"gRPC-Remote-Procedure-Call-/slack-bot/proto"

	"google.golang.org/grpc"
)

var serviceClient slackbot.SlackBotServiceClient

func InitGRPCClient() {
	client, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Error.Fatal("Error in connection :", err)
	}
	defer client.Close()

	serviceClient = slackbot.NewSlackBotServiceClient(client)
	fmt.Println("Service Client:", serviceClient)

}

func GetServiceClient() slackbot.SlackBotServiceClient {
	fmt.Println("Service Client:", serviceClient)
	return serviceClient
}
