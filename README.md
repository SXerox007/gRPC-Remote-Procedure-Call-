# gRPC-
Introduction to Remote Procedure calls. It contains basic to advance level Microservices API's.

## How to write Protos
```
syntax = "proto3";

message Request {
    int64 phone_number =1;
}

message Response {
    int64 phone_number = 1;
    string message = 2;
    int32 status_code =3;
}

service CreateService{
    rpc Service (Request) returns (Response);
}
```


### Gettting Started (Basic)
Unary API basic API user Get request and Server Send Response
```
Unary API
```
Server Streaming API server continous send data to user in simple User request in one time and server send multiple response
```
Server Streaming API
```
Client Streaming API is opposite to server streaming
```
Client Streaming API
```
Bi-Directional API its like a chat Client and Server send data at same time
```
BI-Directional API
```






