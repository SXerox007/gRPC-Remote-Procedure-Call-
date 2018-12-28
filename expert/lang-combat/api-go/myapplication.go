package main

import (
	"context"
	"gRPC-Remote-Procedure-Call-/expert/lang-combat/protos/login"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
}

func (*Server) Login(ctx context.Context, req *login.LoginRequest) (*login.LoginResponse, error) {
	log.Println("Request Comming ", req.GetLoginRequestData())
	return &login.LoginResponse{
		LoginResponseData: &login.LoginResponseData{
			Success: "true",
			Status:  200,
		},
	}, nil
}

func main() {
	Init()
}

func Init() {
	ServerSetup()
}

func ServerSetup() {
	listner, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Error in server listen: ", err)
		return
	}
	// creds, sslErr := ssl_f.GetCreds()
	// if sslErr != nil {
	// 	log.Fatalln("Some Error Occured in cred file: ", sslErr)
	// 	return
	// }
	srv := CreateNewgRPCServer()
	login.RegisterLoginServiceServer(srv, &Server{})

	log.Println("Server Start:", "localhost:50051")
	if err := srv.Serve(listner); err != nil {
		log.Fatal("Error in Serve the Server:", err)
		return
	}

}

func CreateNewgRPCServer() *grpc.Server {
	return grpc.NewServer()
}
