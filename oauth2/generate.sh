#generate gRPC-Code
#To generate the grpc code 
protoc oauth2/proto/oauth.proto --go_out=plugins=grpc:.

 #brain server-run
go run oauth2/server/brain/server-app.go
# or 
make server

#client-run
go run oauth2/client/client-app.go
# or 
make client

#Rest server run
go run oauth2/server/rest/server.go


#source the keys when run hit command
source oauth2/secure-keys/keys


go get golang.org/x/oauth2
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway


# Generate proto
protoc -I/usr/local/include -I.  -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=google/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. oauth2/proto/oauth.proto

#Reverse Proxy For REST
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. oauth2/proto/oauth.proto


protoc -I /usr/local/include -I. -I $GOPATH/src -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. oauth2/proto/oauth.proto



#curl API Expose in Port 5051
curl -X POST "http://localhost:5051/v1/oauth/10001/true"
