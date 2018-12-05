package main

import (
	"context"
	"gRPC-Remote-Procedure-Call-/expert/lang-combat/api-go/environment"
	"gRPC-Remote-Procedure-Call-/expert/lang-combat/api-go/server"
	"gRPC-Remote-Procedure-Call-/expert/lang-combat/api-go/ssl_f"
	"gRPC-Remote-Procedure-Call-/expert/lang-combat/protos/login"
	"log"
	"net"
)

type Server struct {
}

func (*Server) Login(ctx context.Context, req *login.LoginRequest) (*login.LoginResponse, error) {

	return nil, nil
}

func main() {
	Init()
}

func Init() {
	ServerSetup()
}

func ServerSetup() {
	listner, err := net.Listen(environment.GetEnv(), environment.GetPort())
	if err != nil {
		log.Fatal("Error in server listen: ", err)
		return
	}
	creds, sslErr := ssl_f.GetCreds()
	if sslErr != nil {
		log.Fatalln("Some Error Occured in cred file: ", sslErr)
		return
	}
	srv := server.CreateNewgRPCServer(true, creds)
	login.RegisterLoginServiceServer(srv, &Server{})
	if err := srv.Serve(listner); err != nil {
		log.Fatal("Error in Serve the Server:", err)
		return
	}

}
