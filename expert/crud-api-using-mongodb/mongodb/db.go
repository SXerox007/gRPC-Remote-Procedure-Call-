package mongodb

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func InitDB() error {
	client, err := mongo.NewClient("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return err
}
