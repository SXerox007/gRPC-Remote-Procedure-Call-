package main

import (
	"gRPC-Remote-Procedure-Call-/basic/bi-directional-streaming-API/proto"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
}

func (*Server) ChatService(stream chat.ChatServiceContainer_ChatServiceServer) error {
	for {
		//Recive stream message
		req, err := stream.Recv()
		if err == io.EOF {
			return err
		}
		if err != nil {
			log.Fatalf("Some Error occured: %v", err)
			return err
		}
		log.Println("Message Client:", req.GetMessage())

		err = stream.Send(&chat.ChatServiceResponse{
			Message: "Hi there: " + req.GetMessage(),
		})
		if err != nil {
			log.Fatalf("Some Error occured when send data to client : %v", err)
			return err
		}

		//time.Sleep(2000 * time.Millisecond)
	}
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
	//Register reflection on gRPC server
	reflection.Register(srv)

	chat.RegisterChatServiceContainerServer(srv, &Server{})

	if err := srv.Serve(listner); err != nil {
		log.Fatal("Error in Serve the Server:", err)
	}
}

func CreateNewgRPCServer() *grpc.Server {
	return grpc.NewServer()
}
