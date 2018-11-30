package main

import (
	"log"

	"google.golang.org/grpc"
)

func main() {
	Init()
}

func Init() {
	ClientSetup()
}

func ClientSetup() {
	//we don't have ssl certificate right now so we use grpc.withInsecure
	client, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error in connection : ", err)
	}
	defer client.Close()

	// msg := heavyload.NewHeavyLoadDataContainerServiceClient(client)

	// stream, err := msg.HeavyLoadData(context.Background())

	// if err != nil {
	// 	log.Fatal("Some Error Occured: ", err)
	// }

	//ClientStream(stream)

}
