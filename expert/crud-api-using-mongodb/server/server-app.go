package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"gRPC-Remote-Procedure-Call-/expert/crud-api-using-mongodb/collection/informatica"
	"gRPC-Remote-Procedure-Call-/expert/crud-api-using-mongodb/logs"
	"gRPC-Remote-Procedure-Call-/expert/crud-api-using-mongodb/mongodb"
	"gRPC-Remote-Procedure-Call-/expert/crud-api-using-mongodb/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
}

// Create informatica
func (*Server) CreateInformatica(ctx context.Context, req *informaticapb.InformaticaRequest) (*informaticapb.InformaticaResponse, error) {
	//insert data into the mongodb
	err := informatica.InsertDataInInformatica(req.GetInformatica())
	if err == nil {
		return &informaticapb.InformaticaResponse{
			CommonResponse: &informaticapb.CommonResponse{
				Status:  http.StatusOK,
				Message: "Success",
			},
		}, nil
	}
	return nil, err
}

// Get informatica
func (*Server) GetInformatica(ctx context.Context, req *informaticapb.CommonRequest) (*informaticapb.GetInformaticaResponse, error) {
	data, err := informatica.GetDataFromInformatica(req.GetAccessToken(), req.GetEmail())
	if err == nil {
		informaticaData := []*informaticapb.Informatica{}

		for index, element := range data {
			informaticaData = append(informaticaData[:index], &informaticapb.Informatica{
				Sequence: element.Sequence,
				Title:    element.Title,
				Info:     element.Info,
				HostName: element.HostName,
			},
			)
		}

		return &informaticapb.GetInformaticaResponse{
			CommonResponse: &informaticapb.CommonResponse{
				Status:  http.StatusOK,
				Message: "Success",
			},
			Data: informaticaData,
		}, nil
	}
	return nil, err
}

//update the informatica
func (*Server) UpdateInformatica(ctx context.Context, req *informaticapb.UpdateInformaticaRequest) (*informaticapb.InformaticaResponse, error) {
	err := informatica.UpdateDataInInformatica(req)
	if err == nil {
		return &informaticapb.InformaticaResponse{
			CommonResponse: &informaticapb.CommonResponse{
				Status:  http.StatusOK,
				Message: "Success",
			},
		}, nil
	}
	return nil, err

}

//delete the informatica data
func (*Server) DeleteInformatica(ctx context.Context, req *informaticapb.UpdateInformaticaRequest) (*informaticapb.InformaticaResponse, error) {
	err := informatica.DeleteInformaticaRow(req)
	if err == nil {
		return &informaticapb.InformaticaResponse{
			CommonResponse: &informaticapb.CommonResponse{
				Status:  http.StatusOK,
				Message: "Success",
			},
		}, nil
	}
	return nil, err
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
	dbSetup()
	ServerSetup()
}

func logsSetup() {
	logs.InitLogs()
}

func dbSetup() {
	if err := mongodb.InitDB(); err != nil {
		log.Fatal("Can't connect to mongoDB: ", err)
		return
	}

}

func ServerSetup() {
	listner, err := net.Listen(GetEnv(), ":8080")
	if err != nil {
		log.Fatal("Error in server listen: ", err)
		return
	}

	srv := CreateNewgRPCServer()
	//Register reflection on gRPC server
	reflection.Register(srv)

	informaticapb.RegisterInformaticaServiceServer(srv, &Server{})

	go func() {
		log.Println("Starting Server: localhost:8080")
		if err := srv.Serve(listner); err != nil {
			log.Fatal("Error in Serve the Server:", err)
			return
		}
	}()
	//make a channnel that will wait for server to close or interrupt by control^c
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// this will block the code while server
	<-ch
	log.Println("Exit the server with ", os.Interrupt)
}

func CreateNewgRPCServer() *grpc.Server {
	return grpc.NewServer()
}
