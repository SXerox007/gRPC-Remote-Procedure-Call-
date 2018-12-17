#generate gRPC-Code
#To generate the grpc code 
protoc blockchain/protos/blockchain.proto --go_out=plugins=grpc:.

 #server-run
go run blockchain/server/server-app.go
# or 
make server

#client-run
go run blockchain/client/client-app.go
//# or 
make client

