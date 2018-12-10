package main

import (
	"context"
	"log"
	"math/rand"

	"gRPC-Remote-Procedure-Call-/expert/crud-api-using-mongodb/proto"

	"google.golang.org/grpc"
)

func main() {
	Init()
}

func Init() {
	ClientSetup()
}

func ClientSetup() {
	client, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error in connection : ", err)
	}
	defer client.Close()

	msg := informaticapb.NewInformaticaServiceClient(client)

	req := &informaticapb.InformaticaRequest{
		Informatica: &informaticapb.Informatica{
			Sequence: generateRandomSequence(),
			Title:    string(generateRandomSequence()) + " Informatica",
			Info:     "Informatica",
			HostName: "Unknown",
		},
	}
	stream, err := msg.CreateInformatica(context.Background(), req)
	if err == nil {
		log.Println("Data: ", stream.GetCommonResponse())
	} else {
		log.Println("Error: ", err)
	}

}

func generateRandomSequence() int32 {
	return int32(rand.Intn(999) * 33 / (2 - rand.Intn(5)))
}
