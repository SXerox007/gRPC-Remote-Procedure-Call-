package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"gRPC-Remote-Procedure-Call-/slack-bot/client/slackclient"
	"gRPC-Remote-Procedure-Call-/slack-bot/log"
	"gRPC-Remote-Procedure-Call-/slack-bot/proto"

	"github.com/nlopes/slack"
	"google.golang.org/grpc"
)

func initLogs() {
	logLevel := flag.Int("loglevel", 1, "an integer value (0-4)")
	flag.Parse()
	// save logs in log.txt
	log.SetLogLevel(log.Level(*logLevel), "client_logs.txt")
	err := errors.New("Start trace Client Important Error")
	//to trace the error
	log.Info.Println(err)
}

func main() {
	Init()
}

func Init() {
	initLogs()
	slackbotReciveMsgSetup()
}

func slackbotReciveMsgSetup() {
	slackclient.InitSlack()
	rtm := slackclient.CreateRTM()

	//manage the connection
	go rtm.ManageConnection()
	//incoming message
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			//if len(ev.BotID) == 0 {
			go incomingMessages(ev)
			//}
		}
	}
}

func incomingMessages(ev *slack.MessageEvent) {
	//print the slack incoming msg
	fmt.Println(ev)
	err, msg := msgBotSend(ev)
	if err == nil {
		log.Error.Fatalln("Error Slack:", err)
	}
	//For now bot msg as dummy
	InitGRPCClient(ev.Msg.Text, msg)
}

func msgBotSend(ev *slack.MessageEvent) (error, string) {
	_, _, err := slackclient.GetSlackClient().PostMessage(ev.Channel, slack.MsgOptionText("First Bot Msg", true))
	if err != nil {
		log.Error.Println("Error Slack:", err)
	}
	return err, "Dummy Message From Bot"
}

// func testFunc(ev *slack.MessageEvent) {
// 	//slackclient.GetSlackClient().PostEphemeral
// 	var attachment = slack.Attachment{
// 		Text:       "Which beer do you want? :beer:",
// 		Color:      "#f9a41b",
// 		CallbackID: "beer",
// 		Actions: []slack.AttachmentAction{
// 			{
// 				Name: actionSelect,
// 				Type: "select",
// 				Options: []slack.AttachmentActionOption{
// 					{
// 						Text:  "Asahi Super Dry",
// 						Value: "Asahi Super Dry",
// 					},
// 					{
// 						Text:  "Kirin Lager Beer",
// 						Value: "Kirin Lager Beer",
// 					},
// 					{
// 						Text:  "Sapporo Black Label",
// 						Value: "Sapporo Black Label",
// 					},
// 					{
// 						Text:  "Suntory Malt’s",
// 						Value: "Suntory Malts",
// 					},
// 				},
// 			},
// 			{
// 				Name:  actionCancel,
// 				Text:  "Cancel",
// 				Type:  "button",
// 				Style: "danger",
// 			},
// 		},
// 	}

// }

func InitGRPCClient(userMessage string, botReply string) {
	client, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Error.Fatal("Error in connection : ", err)
	}
	defer client.Close()

	msg := slackbot.NewSlackBotServiceClient(client)

	dumpSlackBotMsg(msg, userMessage, botReply)
}

func dumpSlackBotMsg(msg slackbot.SlackBotServiceClient, userMessage string, botReply string) {
	req := &slackbot.SlackDumpRequest{
		QuestionFromUser: userMessage,
		AnswerFromAi:     botReply,
	}

	res, err := msg.SlackDumpingGround(context.Background(), req)
	if err == nil {
		fmt.Println("Data: ", res.GetCommonResponse())
	} else {
		log.Error.Println("Slack Message Dumping Error: ", err)
	}
}
