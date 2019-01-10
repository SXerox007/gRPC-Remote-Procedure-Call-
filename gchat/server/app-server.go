package main

import (
	"errors"
	"flag"
	"fmt"
	"gRPC-Remote-Procedure-Call-/expert/crud-api-using-mongodb/logs"
	"gRPC-Remote-Procedure-Call-/gchat/log"
	"gRPC-Remote-Procedure-Call-/gchat/proto"
	"gRPC-Remote-Procedure-Call-/slack-bot/mongodb"

	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
}

func GetEnv() string {
	env, isEnv := os.LookupEnv("environment")
	if !isEnv {
		env = "tcp"
	}
	return env
}

func main() {
	Init()
}

func Init() {
	logsSetup()
	MongodbSetup()
	ServerSetup()
}

func initLogs() {
	logLevel := flag.Int("loglevel", 1, "an integer value (0-4)")
	flag.Parse()
	// save logs in log.txt
	log.SetLogLevel(log.Level(*logLevel), "server_logs_chat.txt")
	err := errors.New("Start trace Important Error")
	//to trace the error
	log.Info.Println(err)
}

func logsSetup() {
	initLogs()
	logs.InitLogs()

}

func MongodbSetup() {
	if err := mongodb.InitDB(); err != nil {
		//for saving
		log.Error.Println("Can't connect to mongoDB: ", err)
		return
	}
}

func ServerSetup() {
	listner, err := net.Listen(GetEnv(), ":8080")
	if err != nil {
		log.Error.Println("Error in server listen: ", err)
		return
	}

	srv := CreateNewgRPCServer()
	//Register reflection on gRPC server
	reflection.Register(srv)

	gChat.RegisterChatRoomServer(srv, &Server{})

	go func() {
		fmt.Println("Starting Server: localhost:8080")
		if err := srv.Serve(listner); err != nil {
			log.Error.Println("Error in Serve the Server:", err)
			return
		}
	}()
	//make a channnel that will wait for server to close or interrupt by control^c
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// this will block the code while server
	<-ch
	fmt.Println("Exit the server with ", os.Interrupt)
}

func CreateNewgRPCServer() *grpc.Server {
	return grpc.NewServer()
}
