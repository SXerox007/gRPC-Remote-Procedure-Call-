package main

import (
	"context"
	"log"
	"time"

	"gRPC-Remote-Procedure-Call-/basic/client-streaming-API/proto"

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

	msg := heavyload.NewHeavyLoadDataContainerServiceClient(client)

	stream, err := msg.HeavyLoadData(context.Background())

	if err != nil {
		log.Fatal("Some Error Occured: ", err)
	}

	ClientStream(stream)

}

func ClientStream(stream heavyload.HeavyLoadDataContainerService_HeavyLoadDataClient) {
	req := []*heavyload.HeavyLoadDataRequest{
		&heavyload.HeavyLoadDataRequest{
			ReqLoad: &heavyload.HeavyLoadInsert{
				Add: 1,
			},
		},
		&heavyload.HeavyLoadDataRequest{
			ReqLoad: &heavyload.HeavyLoadInsert{
				Add: 2,
			},
		},
		&heavyload.HeavyLoadDataRequest{
			ReqLoad: &heavyload.HeavyLoadInsert{
				Add: 3,
			},
		},
		&heavyload.HeavyLoadDataRequest{
			ReqLoad: &heavyload.HeavyLoadInsert{
				Add: 4,
			},
		},
		&heavyload.HeavyLoadDataRequest{
			ReqLoad: &heavyload.HeavyLoadInsert{
				Add: 5,
			},
		},
	}

	for _, item := range req {
		log.Println("Stream Data")
		stream.Send(item)
		time.Sleep(2000 * time.Millisecond)
	}

	//close sending data and getting response
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalln("Error in getting response: %v", err)
	}

	log.Println("Success Data: ", res)
}
