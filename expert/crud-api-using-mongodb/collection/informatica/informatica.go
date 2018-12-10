package informatica

import (
	"context"
	"fmt"
	"gRPC-Remote-Procedure-Call-/expert/crud-api-using-mongodb/mongodb"
	"gRPC-Remote-Procedure-Call-/expert/crud-api-using-mongodb/proto"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"

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
		//data := []*Informatica{}
		//use filter if you want to get data from a particular id
		//filter := bson.M("_id",data.id)
		res, err := mongodb.CreateCollection("informaticaData").Find(context.Background(), nil)
		if err != nil {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintln("No data found", err))
		}
		var items []*Informatica
		for res.Next(nil) {
			item := Informatica{}
			if err := res.Decode(&item); err != nil {
				return nil, status.Errorf(
					codes.Aborted,
					fmt.Sprintln("Data Can't be decoded", err))
			}
			items = append(items, &item)
		}

		return items, nil
	}
	return nil, status.Errorf(
		codes.Unauthenticated,
		fmt.Sprintln("Access Token is False"))
}

func UpdateDataInInformatica(req *informaticapb.UpdateInformaticaRequest) error {
	InsertDataInInformatica(req.GetInformatica())
	data := &Informatica{}
	filter := bson.M{"sequence": req.GetUpdateSequence()}

	res := mongodb.CreateCollection("informaticaData").FindOne(context.Background(), filter)

	//for single data decode
	if err := res.Decode(data); err != nil {
		fmt.Sprintln("Data without decode", res)
		return status.Errorf(
			codes.Aborted,
			fmt.Sprintln("Data Can't be decoded", err))
	}

	//update host name on particular sequence
	filter = bson.M{"_id": data.ID}
	data.HostName = req.GetHostName()

	_, err := mongodb.CreateCollection("informaticaData").ReplaceOne(context.Background(), filter, data)

	if err != nil {
		return status.Errorf(
			codes.Aborted,
			fmt.Sprintln("Data can't updated: ", err))
	}
	return nil
}

func DeleteInformaticaRow(req *informaticapb.UpdateInformaticaRequest) error {
	filter := bson.M{"sequence": req.GetUpdateSequence()}
	_, err := mongodb.CreateCollection("informaticaData").DeleteOne(context.Background(), filter)
	if err != nil {
		return status.Errorf(
			codes.Aborted,
			fmt.Sprintln("Some error occured while delete data: ", err))
	}
	return nil
}
