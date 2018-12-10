package informatica

import (
	"context"
	"fmt"
	"gRPC-Remote-Procedure-Call-/expert/crud-api-using-mongodb/mongodb"
	"gRPC-Remote-Procedure-Call-/expert/crud-api-using-mongodb/proto"
	"log"

	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Informatica struct {
	ID       objectid.ObjectID `bson:"_id,omitempty"`
	Sequence int32             `bson:"sequence"`
	Title    string            `bson:"title"`
	Info     string            `bson:"info"`
	HostName string            `bson:"host_name"`
}

func InsertDataInInformatica(req *informaticapb.Informatica) error {
	data := Informatica{
		Sequence: req.GetSequence(),
		Title:    req.GetTitle(),
		Info:     req.GetInfo(),
		HostName: req.GetHostName(),
	}

	res, err := mongodb.CreateCollection("informaticaData").InsertOne(context.Background(), data)
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

func GetDataFromInformatica(accessToken bool, email string) ([]*Informatica, error) {
	if accessToken {
		data := []*Informatica{}
		//use filter if you want to get data from a particular id
		//filter := bson.M("_id",data.id)
		res, err := mongodb.CreateCollection("informaticaData").Find(context.Background(), nil)
		if err != nil {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintln("No data found", err))
		}

		if err := res.Decode(data); err != nil {
			return nil, status.Errorf(
				codes.Aborted,
				fmt.Sprintln("Data Can't be decoded", err))
		}

		return data, nil
	}
	return nil, status.Errorf(
		codes.Unauthenticated,
		fmt.Sprintln("Access Token is False"))
}
