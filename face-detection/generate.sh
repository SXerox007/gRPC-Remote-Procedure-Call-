// #generate gRPC-Code
//To generate the grpc code 
protoc acedetection/proto/upload.proto --go_out=plugins=grpc:.

// #server-run
go run face-detection/server/app-server.go
or 
make server

// #client-run
go run face-detection/client/app-client.go
or 
make client
