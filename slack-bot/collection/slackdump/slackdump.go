package slackdump

import (
	"context"
	"fmt"
	"gRPC-Remote-Procedure-Call-/slack-bot/mongodb"
	"gRPC-Remote-Procedure-Call-/slack-bot/proto"
	"log"
	"time"

	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SlackDumpData struct {
	ID            objectid.ObjectID `bson:"_id,omitempty"`
	SenderMessage string            `bson:"sender_message"`
	ChatBotReply  string            `bson:"chatbot_reply"`
	CaptureTime   time.Time         `bson:"capture_time"`
}

func DumpingGroundSlackbot(req *slackbot.SlackDumpRequest) error {
	data := SlackDumpData{
		SenderMessage: req.GetQuestionFromUser(),
		ChatBotReply:  req.GetAnswerFromAi(),
		CaptureTime:   time.Now(),
	}
	res, err := mongodb.CreateCollection("dump").InsertOne(context.Background(), data)
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintln("Internal MongoDb Error:", err))
	}

	oid, ok := res.InsertedID.(objectid.ObjectID)
	if !ok {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintln("Internal oid Error:", ok))
	}

	log.Println("oid: ", oid)
	return nil
}
