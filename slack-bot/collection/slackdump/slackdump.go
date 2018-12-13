package slackdump

import (
	"gRPC-Remote-Procedure-Call-/slack-bot/proto"
	"time"

	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
)

type SlackDumpData struct {
	ID            objectid.ObjectID `bson:"_id,omitempty"`
	SenderMessage string            `bson:"sender_message"`
	ChatBotReply  string            `bson:"chatbot_reply"`
	CaptureTime   time.Time         `bson:"capture_time"`
}

func DumpingGroundSlackbot(req *slackbot.SlackDumpRequest) error {
	return nil
}
