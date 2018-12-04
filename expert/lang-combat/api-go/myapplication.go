package main

import (
	"gRPC-Remote-Procedure-Call-/expert/lang-combat/api-go/environment"
	"gRPC-Remote-Procedure-Call-/expert/lang-combat/api-go/server"
	"gRPC-Remote-Procedure-Call-/expert/lang-combat/api-go/ssl_f"
	"log"
	"net"

	"google.golang.org/grpc/credentials"
)

func main() {
	Init()
}

func Init() {
	ServerSetup()
}

func ServerSetup() {
	listner, err := net.Listen(environment.GetEnv(), ":8080")
	if err != nil {
		log.Fatal("Error in server listen: ", err)
	}

	creds, sslErr := credentials.NewServerTLSFromFile(ssl_f.GetServerCertFile(), ssl_f.GetServerKeyFile())
	if sslErr != nil {
		log.Fatalln("Some Error Occured in cred file: ", sslErr)
		return
	}

	srv := server.CreateNewgRPCServer(true, creds)

	if err := srv.Serve(listner); err != nil {
		log.Fatal("Error in Serve the Server:", err)
	}

}
