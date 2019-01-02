package main

import (
	"context"
	"log"
	"time"

	"gRPC-Remote-Procedure-Call-/expert/multipart-gRPC-chunking/proto"

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

	msg := uploadpb.NewUploadServiceClient(client)

	stream, err := msg.UploadFileService(context.Background())

	if err != nil {
		log.Fatal("Some Error Occured: ", err)
	}

	UploadFileStream(stream)

}

func UploadFileStream(stream uploadpb.UploadService_UploadFileServiceClient) {

	//convert image into small pices of chunk
	image, err := ioutil.ReadFile("expert/multipart-gRPC-chunking/sample-image/test.jpg")
	if err != nil {
		log.Fatalln("Error While Reading the file:",err)
	}

	req := 

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
