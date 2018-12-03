package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gRPC-Remote-Procedure-Call-/expert/error-handling/proto"

	"google.golang.org/grpc"
)

type Server struct {
}

func (*Server) ErrorHandlingService(ctx context.Context, req *errorpb.EvenNumberRequest) (*errorpb.SuccessResponse, error) {
	//its just a sample to check the number is even or not and throw the error message if not a even number
	if req.GetNumber()%2 == 0 {
		//even number
		res := &errorpb.SuccessResponse{
			SuccessInfo: []*errorpb.SuccessResponse_SuccessInfo{
				&errorpb.SuccessResponse_SuccessInfo{
					Message: "Success",
					Status:  http.StatusOK,
					Data: &errorpb.Data{
						Name: &errorpb.Name{
							FirstName:  "Sumit",
							LastName:   "Thakur",
							MiddleName: "",
						},
						Color: errorpb.Color_BLUE,
					},
				},
				&errorpb.SuccessResponse_SuccessInfo{
					Message: "Success",
					Status:  http.StatusOK,
					Data: &errorpb.Data{
						Name: &errorpb.Name{
							FirstName:  "Hello",
							LastName:   "World",
							MiddleName: "Zenith",
						},
						Color: errorpb.Color_RED,
					},
				},
			},
		}
		return res, nil

	} else {
		//odd number
		// throw error msg
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Not a even number", req.GetNumber()),
		)
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

	errorpb.RegisterErrorHandlingServiceServer(srv, &Server{})

	if err := srv.Serve(listner); err != nil {
		log.Fatal("Error in Serve the Server:", err)
	}

}

func CreateNewgRPCServer() *grpc.Server {
	return grpc.NewServer()
}
