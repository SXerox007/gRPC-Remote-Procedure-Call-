package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"gRPC-Remote-Procedure-Call-/face-detection/proto"

	"google.golang.org/grpc"
)

const (
	C_SIZE = 1024 * 1024 // 1MB
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
	image, err := ioutil.ReadFile("face-detection/sample-image/sumit.jpg")
	if err != nil {
		log.Fatalln("Error While Reading the file:", err)
	}
	log.Println("Whole Image bytes:", image)
	req := &uploadpb.UploadChunkRequest{}
	for start := 0; start < len(image); start += C_SIZE {
		if start+C_SIZE > len(image) {
			req.FileChunk = image[start:len(image)]
		} else {
			req.FileChunk = image[start : start+C_SIZE]
		}
		if err := stream.Send(req); err != nil {
			log.Fatalln("Error in send:", err)

		}
		//remove this line this will only for demo purpose
		time.Sleep(2000 * time.Millisecond)
	}

	//close sending data and getting response
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalln("Error in getting response: %v", err)
	}

	log.Println("Success Data: ", res)
}
