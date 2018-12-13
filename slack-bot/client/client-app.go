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
	fmt.Println("Here 1 point")
	slackclient.InitSlack()
	rtm := slackclient.CreateRTM()
	fmt.Println("Here 1 point")
	//manage the connection
	go rtm.ManageConnection()
	//incoming message
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			//if len(ev.BotID) == 0 {
			fmt.Println("Here 2 point")
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
	_, _, err := slackclient.GetSlackClient().PostMessage(ev.User, slack.MsgOptionPostMessageParameters(
		slack.PostMessageParameters{
			AsUser:         true,
			ReplyBroadcast: true,
			Parse:          "Parser msg",
			Channel:        ev.Channel,
			User:           ev.User,
		},
	))

	if err != nil {
		log.Error.Println("Error Slack:", err)
	}

	return err, "Dummy Message From Bot"
}

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
