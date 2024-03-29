package main

import (
	"context"
	"fmt"
	"log"

	calculatorpc "github.com/AskarKasimov/grpc-calculator/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Pokemon Client")

	cc, err := grpc.Dial(":4041", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close() // Может, в отдельной функции ошибка будет обработана?

	c := calculatorpc.NewExpressionServiceClient(cc)

	createExpression, err := c.CreateExpression(context.Background(), &calculatorpc.CreateExpressionRequest{Vanilla: "123 + 12312313"})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Println(createExpression.GetExpression())
}
