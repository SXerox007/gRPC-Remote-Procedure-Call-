// #generate gRPC-Code
//To generate the grpc code 
protoc gchat/proto/chat.proto --go_out=plugins=grpc:.

// #server-run
go run gchat/server/app-server.go
or 
make server

// #client-run
go run gchat/client/app-client.go
or 
make client
