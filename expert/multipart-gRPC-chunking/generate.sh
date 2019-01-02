// #generate gRPC-Code
//To generate the grpc code 
protoc expert/multipart-gRPC-chunking/proto/upload.proto --go_out=plugins=grpc:.

// #server-run
go run expert/multipart-gRPC-chunking/server/app-server.go
or 
make server

// #client-run
go run expert/multipart-gRPC-chunking/client/app-client.go
or 
make client

// #Note:-
Note:- 
For changing the proto buffer file name use:
option go_package = "example_newpb";
package example;



// #gRPCvsREST || HTTP1.1 vs HTTP2
https://husobee.github.io/golang/rest/grpc/2016/05/28/golang-rest-v-grpc.html
https://imagekit.io/demo/http2-vs-http1