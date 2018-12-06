package collection

import (
	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
)

type Informatica struct {
	ID       objectid.ObjectID `bson:"_id,omitempty"`
	Sequence int32             `bson:"sequence"`
	Title    string            `bson:"title"`
	Info     string            `bson:"info"`
	HostName string            `bson:"host_name"`
}
