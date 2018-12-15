package slackdump

import (
	"fmt"
	"gRPC-Remote-Procedure-Call-/slack-bot/log"
	db "gRPC-Remote-Procedure-Call-/slack-bot/postgres"
	slackbot "gRPC-Remote-Procedure-Call-/slack-bot/proto"
	"time"
)

type SlackDumpData struct {
	TableName     struct{}  `sql:"slack_dump" json:"-"`
	Id            string    `param:"id"`
	SenderMessage string    `param:"sender_message"`
	ChatBotReply  string    `param:"chatbot_reply"`
	CaptureTime   time.Time `param:"capture_time"`
}

func DumpingGroundSlackbot(req *slackbot.SlackDumpRequest) error {
	sqlStatement := `
	INSERT INTO slack_dump (sender_message, chatbot_reply, capture_time)
	VALUES ($1, $2, $3)
	RETURNING id`
	var id string
	err := db.GetClient().QueryRow(req.GetQuestionFromUser(), req.GetAnswerFromAi(), time.Now()).Scan(&id)
	fmt.Println("New Record ID is:", id)
	log.Info.Println("Dump Data :", sqlStatement)
	return err
}
