#generate gRPC-Code
#To generate the grpc code 
protoc expert/crud-api-using-mongodb/proto/informatica.proto --go_out=plugins=grpc:.

 #server-run
go run expert/crud-api-using-mongodb/server/server-app.go
# or 
make server

#client-run
go run slack-bot/client/client-app.go
//# or 
make client


#For get slack package 
go get github.com/nlopes/slack
//#source the keys when run hit command
source slack-bot/secure-keys/keys


#We use wit.ai for the msgs
#create account and just use
