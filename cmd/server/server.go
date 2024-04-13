package main

import (
	"context"
	"log"
	"net"

	db "github.com/AskarKasimov/grpc-calculator/internal/server_db"
	calculatorpc "github.com/AskarKasimov/grpc-calculator/pkg/proto"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpc.UnimplementedExpressionServiceServer
}

func NewServer() *server {
	return &server{}
}

func (*server) Register(tx context.Context, req *calculatorpc.RegisterRequest) (*calculatorpc.RegisterResponse, error) {
	id, _ := db.DB().GetWorkerIdByName(req.GetName())

	if id != "" {
		return &calculatorpc.RegisterResponse{
			Id: id,
		}, nil
	}

	createdId, err := db.DB().NewWorker(req.GetName())
	if err != nil {
		return nil, err
	}
	return &calculatorpc.RegisterResponse{
		Id: createdId,
	}, nil
}

func (*server) CreateExpression(tx context.Context, req *calculatorpc.CreateExpressionRequest) (*calculatorpc.CreateExpressionResponse, error) {
	vanilla := req.GetVanilla()
	return &calculatorpc.CreateExpressionResponse{
		Expression: &calculatorpc.Expression{
			Id:           "1",
			IncomingDate: "2.0",
			Vanilla:      vanilla,
			Progress:     "waiting",
			Answer:       "",
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calculatorpc.RegisterExpressionServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
