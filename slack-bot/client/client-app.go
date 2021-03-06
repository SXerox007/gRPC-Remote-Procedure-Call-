package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	gc "gRPC-Remote-Procedure-Call-/slack-bot/client/grpcclient"
	"gRPC-Remote-Procedure-Call-/slack-bot/client/slackclient"
	"gRPC-Remote-Procedure-Call-/slack-bot/log"
	"gRPC-Remote-Procedure-Call-/slack-bot/proto"

	"github.com/nlopes/slack"
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
	gc.InitGRPCClient()
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
			if len(ev.BotID) == 0 {
				go incomingMessages(ev)
			}
		}
	}
}

func incomingMessages(ev *slack.MessageEvent) {
	//print the slack incoming msg
	//fmt.Println(ev)
	err, msg := msgBotSend(ev)
	if err != nil {
		log.Error.Fatalln("Error Slack:", err)
	} else {
		dumpSlackBotMsg(ev.Text, msg)
	}
}

func msgBotSend(ev *slack.MessageEvent) (error, string) {
	_, _, err := slackclient.GetSlackClient().PostMessage(ev.Channel, slack.MsgOptionText("First Bot Msg", true))
	if err != nil {
		log.Error.Println("Error Slack:", err)
	}
	return err, "Dummy Message From Bot"
}

func msgAttachment(ev *slack.MessageEvent) {
	//slackclient.GetSlackClient().PostEphemeral
	var attachment = slack.Attachment{
		Text:       "Which branch you want to merger in staging?",
		Color:      "#f9a41b",
		CallbackID: "branch",
		Actions: []slack.AttachmentAction{
			{
				Name: "actionSelect",
				Type: "select",
				Options: []slack.AttachmentActionOption{
					{
						Text:  "master",
						Value: "master",
					},
					{
						Text:  "develop",
						Value: "develop",
					},
					{
						Text:  "testing",
						Value: "testing",
					},
					{
						Text:  "Live-2.0",
						Value: "Live-2.0",
					},
				},
			},
			{
				Name:  "actionCancel",
				Text:  "Cancel",
				Type:  "button",
				Style: "danger",
			},
		},
	}

	log.Info.Println("Attactchment msg smaple", attachment)
}

func dumpSlackBotMsg(userMessage string, botReply string) {
	req := &slackbot.SlackDumpRequest{
		QuestionFromUser: userMessage,
		AnswerFromAi:     botReply,
		MongodbEnable:    true,
		PostgresEnable:   true,
	}

	res, err := gc.GetServiceClient().SlackDumpingGround(context.Background(), req)
	if err == nil {
		fmt.Println("Data: ", res.GetCommonResponse())
	} else {
		log.Error.Println("Slack Message Dumping Error: ", err)
	}
}
