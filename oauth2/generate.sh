#generate gRPC-Code
#To generate the grpc code 
protoc oauth2/proto/oauth.proto --go_out=plugins=grpc:.

 #server-run
go run oauth2/server/server-app.go
# or 
make server

#client-run
go run oauth2/client/client-app.go
# or 
make client


#source the keys when run hit command
source oauth2/secure-keys/keys


go get golang.org/x/oauth2
