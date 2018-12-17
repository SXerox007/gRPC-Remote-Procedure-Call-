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

### Advance (gRPC)

Create Update Read data form Monogo db using gRPC
```
Curd API Mongodb
```
Create Update Read data form Postgres using gRPC
```
Curd API Postgres
```
Deadline if the response form server not coming in a particular time how to call deadline
```
Deadline
```
Error Handling in gRPC its just a basic how to do
for More :
https://grpc.io/docs/guides/error.html
http://avi.im/grpc-errors/
```
Error handling
```
Server in go and Client in java. Multiple Language combat using Protos same proto use server(Go) and client (Java)
```
Lang Combat
```

Reflection Its a great a tool See our gRPC code test our code. Best to use
```
Reflection
```
ssl security in gRPC
```
ssl security
```



