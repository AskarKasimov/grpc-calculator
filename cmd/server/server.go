package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/AskarKasimov/grpc-calculator/pkg/db"
	calculatorpc "github.com/AskarKasimov/grpc-calculator/pkg/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func runRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := calculatorpc.RegisterExpressionServiceHandlerFromEndpoint(ctx, mux, "localhost:12201", opts)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

func runGrpc() {
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

func main() {
	go runRest()
	runGrpc()
}
