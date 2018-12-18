package gtoken

import (
	"crypto/tls"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

func AuthAccess() []grpc.DialOption {
	perRPC := oauth.NewOauthAccess(fetchToken())
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(perRPC),
		grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}),
		),
	}
	return opts
}

func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "ACCESS_TOKEN",
	}
}
