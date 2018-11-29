package main

import (
	"gRPC-Remote-Procedure-Call-/basic/server-streaming-API/proto"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
}

//server streaming in gRPC
func (*Server) DataPacketService(req *messages.DataPacketRequest, stream messages.DataPacketMessageContainerService_DataPacketServiceServer) error {
	if req.ReqDataPacket {
		for i := 1; i <= 5; i++ {
			res := &messages.DataPacketResponse{
				ResDataPacket: &messages.DataPacketContainer{
					Message:     "Success DataPacket" + strconv.Itoa(i),
					Description: "Server streaming",
					Code:        strconv.Itoa(i),
				},
			}
			stream.Send(res)
			time.Sleep(2000 * time.Millisecond)
		}
	}
	return nil
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
	ServerSetup()
}

func ServerSetup() {
	listner, err := net.Listen(GetEnv(), ":8080")
	if err != nil {
		log.Fatal("Error in server listen: ", err)
	}

	srv := CreateNewgRPCServer()

	messages.RegisterDataPacketMessageContainerServiceServer(srv, &Server{})

	if err := srv.Serve(listner); err != nil {
		log.Fatal("Error in Serve the Server:", err)
	}
}

func CreateNewgRPCServer() *grpc.Server {
	return grpc.NewServer()
}
