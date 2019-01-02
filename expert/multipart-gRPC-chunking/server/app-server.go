package main

import (
	"gRPC-Remote-Procedure-Call-/expert/multipart-gRPC-chunking/proto"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type Server struct {
}

func (*Server) UploadFileService(stream uploadpb.UploadService_UploadFileServiceServer) error {
	var result []byte
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			//finish upload chunk file
			return stream.SendAndClose(&uploadpb.UploadResponse{
				Message: "Success in file upload",
				Code:    200,
			})
		}
		if err != nil {
			//show error
			log.Fatal("Some Error Occured &v", err)
		}

		//print client streaming data
		//log.Println("Client Streaming Data:", req.GetReqLoad().GetAdd())
		//result += "Make Result" + fmt.Sprint(req.GetReqLoad().GetAdd())
		result = append(result, data.GetFileChunk()...)
	}
	log.Println("Image byte:", result)
}

func GetEnv() string {
	env, isEnv := os.LookupEnv("environment")
	if !isEnv {
		env = "tcp"
	}
	return env
}

func main() {
	Init()
}

func Init() {
	ServerSetup()
}

func ServerSetup() {
	listner, err := net.Listen(GetEnv(), ":8080")
	if err != nil {
		log.Fatal("Error in server listen: ", err)
	}

	srv := CreateNewgRPCServer()

	uploadpb.RegisterUploadServiceServer(srv, &Server{})

	if err := srv.Serve(listner); err != nil {
		log.Fatal("Error in Serve the Server:", err)
	}
}

func CreateNewgRPCServer() *grpc.Server {
	return grpc.NewServer()
}
