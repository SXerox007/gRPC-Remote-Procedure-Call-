package slackdump

import "time"

type SlackDumpData struct {
	TableName     struct{}  `sql:"slack_dump" json:"-"`
	Id            string    `param:"id"`
	SenderMessage string    `param:"sender_message"`
	ChatBotReply  string    `param:"chatbot_reply"`
	CaptureTime   time.Time `param:"capture_time"`
}
