package main

import (
	"context"
	"gRPC-Remote-Procedure-Call-/basic/server-streaming-API/proto"
	"io"
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

	msg := messages.NewDataPacketMessageContainerServiceClient(client)

	req := &messages.DataPacketRequest{
		ReqDataPacket: true,
	}
	data, err := msg.DataPacketService(context.Background(), req)

	if err != nil {
		log.Fatal("Some Error Occured: ", err)
	}

	streamData(data)

}

func streamData(data messages.DataPacketMessageContainerService_DataPacketServiceClient) {
	for {
		dataStream, err := data.Recv()
		log.Println("Response Send by Server:")
		if err == io.EOF {
			//if End of file then server stream complete
			break
		}
		if err != nil {
			log.Println("Error while streaming: %v", err)
		}
		log.Println(dataStream)
	}
}
