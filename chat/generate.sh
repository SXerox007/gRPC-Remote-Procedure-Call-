#generate gRPC-Code
#To generate the grpc code 
protoc chat/proto/chat.proto --go_out=plugins=grpc:.

 #server-run
go run chat/server/server-app.go
# or 
make server

#client-run
go run chat/client/client-app.go chat/client/menu.go
//# or 
make client
