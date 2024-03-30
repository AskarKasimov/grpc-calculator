package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	calculatorpc "github.com/AskarKasimov/grpc-calculator/pkg/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var HOST string
var PORT string
var GOROUTINES int
var NAME string = uuid.New().String()
var ID int64
var MULTIPLICATION int64
var DIVISION int64
var ADDITION int64
var SUBTRACTION int64

func init() {
	host := os.Getenv("HOST")
	if host == "" {
		log.Fatalln("No HOST env found")
	}
	HOST = host

	port := os.Getenv("PORT")
	if host == "" {
		log.Fatalln("No PORT env found")
	}
	PORT = port

	goroutines, err := strconv.Atoi(os.Getenv("GOROUTINES"))
	if err != nil {
		log.Fatalln(err)
	}
	GOROUTINES = goroutines

	multiplication, err := strconv.ParseInt(os.Getenv("MULTIPLICATION"), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	MULTIPLICATION = multiplication

	division, err := strconv.ParseInt(os.Getenv("DIVISION"), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	DIVISION = division

	addition, err := strconv.ParseInt(os.Getenv("ADDITION"), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	ADDITION = addition

	subtraction, err := strconv.ParseInt(os.Getenv("SUBTRACTION"), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	SUBTRACTION = subtraction

	log.Printf("Host: %s", HOST)
	log.Printf("Port: %s", PORT)
	log.Printf("Name (uuid): %s", NAME)
	log.Printf("Number of goroutines: %d", GOROUTINES)
	log.Printf("Time (seconds) for multiplication: %d", MULTIPLICATION)
	log.Printf("Time (seconds) for division: %d", DIVISION)
	log.Printf("Time (seconds) for addition: %d", ADDITION)
	log.Printf("Time (seconds) for subtraction: %d", SUBTRACTION)
}

func main() {
	log.Println("Client started!")

	cc, err := grpc.Dial(fmt.Sprintf("%s:%s", HOST, PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpc.NewExpressionServiceClient(cc)

	log.Println("Trying to get ID")

	id, err := c.Register(context.Background(), &calculatorpc.RegisterRequest{
		Name: NAME,
	})

	if err != nil {
		log.Fatalf("Could not connect to get ID: %v", err)
	}

	log.Printf("Got ID: %s", id)

	createExpression, err := c.CreateExpression(context.Background(),
		&calculatorpc.CreateExpressionRequest{
			Vanilla: "123 + 12312313",
		})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	log.Println(createExpression.GetExpression())
}
