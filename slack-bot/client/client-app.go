package main

import (
	"gRPC-Remote-Procedure-Call-/slack-bot/client/slackclient"
	"log"

	"github.com/nlopes/slack"
)

func main() {
	Init()
}

func Init() {
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
	log.Println(ev)
}
