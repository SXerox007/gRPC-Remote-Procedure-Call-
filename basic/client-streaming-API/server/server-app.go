package main

import (
	"fmt"
	"gRPC-Remote-Procedure-Call-/basic/client-streaming-API/proto"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type Server struct {
}

func (*Server) HeavyLoadData(stream heavyload.HeavyLoadDataContainerService_HeavyLoadDataServer) error {
	result := "Result Start: "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//Finish from client streaming
			//When we done
			return stream.SendAndClose(&heavyload.HeavyLoadDataResponse{
				GetLoad: &heavyload.HeavyLoadData{
					Success:    result,
					LoadBinary: 1010,
					Code:       200,
				},
			})
		}
		if err != nil {
			//show error
			log.Fatal("Some Error Occured &v", err)
		}
		//print client streaming data
		log.Println("Client Streaming Data:", req.GetReqLoad().GetAdd())
		result += "Make Result" + fmt.Sprint(req.GetReqLoad().GetAdd())
	}

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

	heavyload.RegisterHeavyLoadDataContainerServiceServer(srv, &Server{})

	if err := srv.Serve(listner); err != nil {
		log.Fatal("Error in Serve the Server:", err)
	}
}

func CreateNewgRPCServer() *grpc.Server {
	return grpc.NewServer()
}
