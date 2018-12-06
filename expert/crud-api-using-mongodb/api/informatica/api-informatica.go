package informatica

import (
	"context"
	"gRPC-Remote-Procedure-Call-/expert/crud-api-using-mongodb/collection/informatica"
	"net/http"
)

//TODO:- Pending
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
