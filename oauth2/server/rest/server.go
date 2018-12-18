package main

import (
	"flag"
	"log"
	"net/http"

	"gRPC-Remote-Procedure-Call-/oauth2/proto"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	authpoint = flag.String("auth_end_points", "localhost:8080", "expose end point of oAuth")
)

func ExposePoint(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	err := oauthpb.RegisterOAuthServiceHandlerFromEndpoint(ctx, mux, *authpoint, dialOpts)
	if err != nil {
		return err
	}
	log.Println("Starting Endpoint Exposed Server: localhost:5051")

	http.ListenAndServe(address, mux)
	return nil
}

func main() {
	Init()
}

func Init() {
	if err := ExposePoint(":5051"); err != nil {
		log.Fatal("Error", err)
	}
}
