package main

import (
	"fmt"
	"gRPC-Remote-Procedure-Call-/slack-bot/client/slackclient"

	"github.com/nlopes/slack"
)

func main() {
	Init()
}

func Init() {
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
		fmt.Println("Error Slack:", err)
	}
}
