package slackclient

import (
	"os"

	"github.com/nlopes/slack"
)

var slackClient *slack.Client

func InitSlack() {
	slackClient = slack.New(os.Getenv("SLACK_ACCESS_TOKEN"))
}

func CreateRTM() *slack.RTM {
	return slackClient.NewRTM()
}

func GetSlackClient() *slack.Client {
	return slackClient
}
