package main

import (
	"errors"
	"flag"
	"gRPC-Remote-Procedure-Call-/slack-bot/client/slackclient"
	"gRPC-Remote-Procedure-Call-/slack-bot/log"

	"github.com/nlopes/slack"
)

func initLogs() {
	logLevel := flag.Int("loglevel", 1, "an integer value (0-4)")
	flag.Parse()
	// save logs in log.txt
	log.SetLogLevel(log.Level(*logLevel), "logs.txt")
	err := errors.New("Start trace Important Error")
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
	//log.Println(ev)
	msgBotSend(ev)
}

func msgBotSend(ev *slack.MessageEvent) {
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
}
