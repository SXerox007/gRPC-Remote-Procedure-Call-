// #generate gRPC-Code
//To generate the grpc code 
protoc proto/chat.proto --go_out=plugins=grpc:.

// #server-run
go run basic/bi-directional-streaming-API/server/server-app.go
or 
make server

// #client-run
go run basic/bi-directional-streaming-API/client/client-app.go
or 
make client

// #Note:-
Note:- 
For changing the proto buffer file name use:
option go_package = "example_newpb";
package example;

// #Some Error which you might face 
1. Make file error:
cat -e -t -v  makefile
https://stackoverflow.com/questions/3931741/why-does-make-think-the-target-is-up-to-date
https://stackoverflow.com/questions/834748/gcc-makefile-error-no-rule-to-make-target


// #Refrence:- 
https://godoc.org/google.golang.org/grpc
https://www.golang-book.com/books/intro/13