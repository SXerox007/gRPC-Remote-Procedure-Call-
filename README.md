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
Add Image upload using gRPC (Chunk the file and send to backend)
### why we send file to small chunks?
The answer is gRPC encode and decode the data when send and recive it take less time if we send data in the form of small packet 
### proto file
```
syntax = "proto3";

package uploadpb;

message UploadChunkRequest{
    bytes file_chunk = 1;
}

message UploadResponse{
    string message = 1;
    int32 code = 2;
}

service UploadService{
    rpc UploadFileService (stream UploadChunkRequest) returns (UploadResponse){};
}
```
### Logic 
```
C_SIZE = 64*64
req := &uploadpb.UploadChunkRequest{}
	for start := 0; start < len(image); start += C_SIZE {
		if start+C_SIZE > len(image) {
			req.FileChunk = image[start:len(image)]
		} else {
			req.FileChunk = image[start : start+C_SIZE]
		}
		if err := stream.Send(req); err != nil {
			log.Fatalln("Error in send:", err)

		}
```

# Slack Bot
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

# BlockChain (gRPC) Basic Example
## Reference :- Alex Pliutau
### what is Blockchain ?
The answer is very simple it's a DS Linkedlist type not linkedlist its just same as linkedlist it just always store the previous hash/id. 
### Who will use ?
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

# OAuth2 and API Expose in gRPC
Google Authentication Sample and Expose the Micoservice API to others

## Go Get First
```
go get golang.org/x/oauth2
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```


### Services 
```
service OAuthService {
    rpc GetOAuthService (OAuthRequest) returns (OAuthResponse);
    rpc GetCodeState(OAuthCodeRequest) returns (OAuthCodeResponse) {
		option (google.api.http) = {
            post: "/v1/oauth/{code}/{state}"
            body: "*"
		};
	}

}
}
```
### To Generate the Proto

```
protoc -I/usr/local/include -I.  -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=google/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. oauth2/proto/oauth.proto
```


### Reverse Proxy
```
protoc -I /usr/local/include -I. -I $GOPATH/src -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. oauth2/proto/oauth.proto

```

## Note: 
* There are 3 servers in this OAuth. 
* Go through the generate.sh file First


### Servers:
* Ist is for gRPC server. (Brain)
* 2nd is the API Expose server (Rest)
* 3rd is for the web page render/ UI run for an example For oAuth (gWeb)


### Expose API Check
```
curl -X POST "http://localhost:5051/v1/oauth/10001/true"
```

## Contact if any problem:
Email: sumitthakur769@gmail.com





