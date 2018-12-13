package slackdump

import "time"

type SlackDumpData struct {
	TableName     struct{}  `sql:"current_balances" json:"-"`
	Id            string    `param:"id" json:"id"`
	SenderMessage string    `param:"sender_message" json:"sender_message"`
	ChatBotReply  string    `param:"chatbot_reply" json:"chatbot_reply"`
	CaptureTime   time.Time `param:"capture_time" json:"capture_time"`
}
