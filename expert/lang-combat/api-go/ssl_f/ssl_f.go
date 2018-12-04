package ssl_f

import (
	"google.golang.org/grpc/credentials"
)

func getServerCertFile() string {
	return "expert/lang-combat/ssl/server.crt"
}

func getServerKeyFile() string {
	return "expert/lang-combat/ssl/server.pem"
}

func GetCreds() (credentials.TransportCredentials, error) {
	return credentials.NewServerTLSFromFile(getServerCertFile(), getServerKeyFile())
}
