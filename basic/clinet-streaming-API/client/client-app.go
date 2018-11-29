package main

import (
	"log"

	"google.golang.org/grpc"
)

func main() {

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

	// msg := messages.NewDataPacketMessageContainerServiceClient(client)

	// req := &messages.DataPacketRequest{
	// 	ReqDataPacket: true,
	// }
	// data, err := msg.DataPacketService(context.Background(), req)

	// if err != nil {
	// 	log.Fatal("Some Error Occured: ", err)
	// }

	// streamData(data)

}
