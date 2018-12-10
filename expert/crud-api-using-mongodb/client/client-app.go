package main

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

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

	//create the informatica
	CreateInformatica(msg)
	//read the informatica
	ReadInformatica(msg)
	//update the informatica
	UpdateInformatica(msg)
	//delete the informatica
	DeleteInformatica(msg)

}

func DeleteInformatica(msg informaticapb.InformaticaServiceClient) {
	req := &informaticapb.UpdateInformaticaRequest{
		Informatica: &informaticapb.Informatica{
			Sequence: generateRandomSequence(),
			Title:    strconv.Itoa(int(generateRandomSequence())) + " Informatica",
			Info:     "Informatica",
			HostName: "Unknown",
		},
		UpdateSequence: 2891,
		HostName:       "Black_Dreams",
	}
	res, err := msg.UpdateInformatica(context.Background(), req)
	if err == nil {
		log.Println("Data: ", res.GetCommonResponse())
	} else {
		log.Println("Error: ", err)
	}
}

func UpdateInformatica(msg informaticapb.InformaticaServiceClient) {
	req := &informaticapb.UpdateInformaticaRequest{
		Informatica: &informaticapb.Informatica{
			Sequence: generateRandomSequence(),
			Title:    strconv.Itoa(int(generateRandomSequence())) + " Informatica",
			Info:     "Informatica",
			HostName: "Unknown",
		},
		UpdateSequence: 2891,
		HostName:       "Black_Dreams",
	}
	res, err := msg.UpdateInformatica(context.Background(), req)
	if err == nil {
		log.Println("Data: ", res.GetCommonResponse())
	} else {
		log.Println("Error: ", err)
	}
}

func ReadInformatica(msg informaticapb.InformaticaServiceClient) {
	req := &informaticapb.CommonRequest{
		AccessToken: true,
		Email:       "sumitthakur769@gmail.com",
	}
	res, err := msg.GetInformatica(context.Background(), req)
	if err == nil {
		log.Println("Common Data: ", res.GetCommonResponse())
		log.Println("List of All Data: ", res.GetData())
	} else {
		log.Println("Error: ", err)
	}
}

func CreateInformatica(msg informaticapb.InformaticaServiceClient) {
	req := &informaticapb.InformaticaRequest{
		Informatica: &informaticapb.Informatica{
			Sequence: generateRandomSequence(),
			Title:    strconv.Itoa(int(generateRandomSequence())) + " Informatica",
			Info:     "Informatica",
			HostName: "Unknown",
		},
	}
	res, err := msg.CreateInformatica(context.Background(), req)
	if err == nil {
		log.Println("Data: ", res.GetCommonResponse())
	} else {
		log.Println("Error: ", err)
	}
}

func generateRandomSequence() int32 {
	rand.Seed(time.Now().UTC().UnixNano())
	return int32(rand.Intn(9999))
}
