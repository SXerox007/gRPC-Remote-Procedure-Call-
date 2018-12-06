// #generate gRPC-Code
//To generate the grpc code 
protoc expert/crud-api-using-mongodb/proto/informatica.proto --go_out=plugins=grpc:.

// #server-run
go run expert/crud-api-using-mongodb/server/server-app.go
//# or 
make server

// #client-run
go run expert/crud-api-using-mongodb/client/client-app.go
//# or 
make client