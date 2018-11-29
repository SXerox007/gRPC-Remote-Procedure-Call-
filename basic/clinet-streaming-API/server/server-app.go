package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type Server struct {
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

	//messages.RegisterDataPacketMessageContainerServiceServer(srv, &Server{})

	if err := srv.Serve(listner); err != nil {
		log.Fatal("Error in Serve the Server:", err)
	}
}

func CreateNewgRPCServer() *grpc.Server {
	return grpc.NewServer()
}
