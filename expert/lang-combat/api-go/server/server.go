package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func CreateNewgRPCServer(isSSL bool, creds credentials.TransportCredentials) *grpc.Server {
	if isSSL {
		return grpc.NewServer(grpc.Creds(creds))
	} else {
		return grpc.NewServer()
	}
}
