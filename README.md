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
Crud API Mongodb
```
Create Update Read data form Postgres using gRPC
```
Crud API Postgres
```
Deadline if the response form server not coming in a particular time how to call deadline
```
Deadline
```
Error Handling in gRPC its just a basic how to do
for More :
* [gRPC Error Official](https://grpc.io/docs/guides/error.html)
* [Another Example](http://avi.im/grpc-errors/)

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

### Slack Bot
Slack bot user read and send message (Basic) using gRPC. DB used postgres and Mongodb. It will dump the data in db of user message and bot response.

```
syntax = "proto3";

message CommonResponse {
    int32 code = 1;
    string message = 2;
}

message SlackDumpRequest {
    string question_from_user = 1;
    string answer_from_ai =2;
    bool mongodb_enable = 3;
    bool postgres_enable = 4;
}

message SlackDumpResponse {
    CommonResponse CommonResponse = 1;
}

service SlackBotService {
    rpc SlackDumpingGround(SlackDumpRequest) returns (SlackDumpResponse);
}
```

### BlockChain (gRPC) Basic Example
## Refrence :- Alex Pliutau
# what is Blockchain ?
The answer is very simple it's a DS Linkedlist type not linkedlist its just same as linkedlist it just always store the previous hash/id. 
# Use
* IBM
* Bitcoin
* Samsung
* etc Many more

```
type Block struct {
	Data          string
	PrevBlockHash string
	Hash          string
}
```


