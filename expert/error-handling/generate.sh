// #generate gRPC-Code
//To generate the grpc code 
protoc expert/error-handling/proto/error.proto --go_out=plugins=grpc:.

// #server-run
go run expert/error-handling/server/server-app.go
or 
make server

// #client-run
go run expert/error-handling/client/client-app.go
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
https://flaviocopes.com/golang-environment-variables/


// #Error handling
https://grpc.io/docs/guides/error.html
http://avi.im/grpc-errors/


// #gRPCvsREST || HTTP1.1 vs HTTP2
https://husobee.github.io/golang/rest/grpc/2016/05/28/golang-rest-v-grpc.html
https://imagekit.io/demo/http2-vs-http1