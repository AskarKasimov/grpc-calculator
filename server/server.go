package main

import (
	"context"
	"log"
	"net"

	calculatorpc "github.com/AskarKasimov/grpc-calculator/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	calculatorpc.ExpressionServiceServer
}

func NewServer() *server {
	return &server{}
}

type expressionItem struct {
	Id           int64
	IncomingDate int64
	Vanilla      string
	Answer       string
	Progress     string
}

func (*server) Register(tx context.Context, req *calculatorpc.RegisterRequest) (*calculatorpc.CreateExpressionResponse, error) {

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
	log.Println("Starting TCP server...")
	lis, err := net.Listen("tcp", "0.0.0.0:4041")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Success!")

	s := grpc.NewServer()
	reflection.Register(s)

	calculatorpc.RegisterExpressionServiceServer(s, NewServer())

	log.Println("Loading...")

	if err := s.Serve(lis); err == nil {
		log.Println("Error serving grpc: ", err)
		return
	}
}
