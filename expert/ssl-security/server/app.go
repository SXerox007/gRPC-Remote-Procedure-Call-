package main

import (
	"context"
	"log"
	"net"
	"os"

	"gRPC-Remote-Procedure-Call-/basic/unary-API/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
}

//thats all about the API response
func (*Server) SearchContact(ctx context.Context, req *search.SearchRequest) (*search.SearchResponse, error) {
	//Phone number is requested by the client
	// phone := req.GetPhoneNumber()
	response := &search.SearchResponse{
		PhoneNumber: req.GetPhoneNumber(),
		Message:     "Success",
	}

	return response, nil
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

	//with ssl certificate
	certFile := "expert/ssl-security/ssl/server.crt"
	keyFile := "expert/ssl-security/ssl/server.pem"
	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if sslErr != nil {
		log.Fatalln("Some Error Occured in cred file: ", sslErr)
		return
	}

	srv := CreateNewgRPCServer(true, creds)

	search.RegisterCreateServiceServer(srv, &Server{})

	if err := srv.Serve(listner); err != nil {
		log.Fatal("Error in Serve the Server:", err)
	}

}

func CreateNewgRPCServer(isSSL bool, creds credentials.TransportCredentials) *grpc.Server {
	if isSSL {
		return grpc.NewServer(grpc.Creds(creds))
	} else {
		return grpc.NewServer()
	}
}
