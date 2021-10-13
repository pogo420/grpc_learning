package main

import (
	"context"
	"grpc_learning/protos"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	// opening connection
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("not able to connection")
	}
	defer conn.Close()

	// getting client
	client := protos.NewCurrencyClient(conn)

	// prepairing context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// RPC call
	res, err := client.GetRate(ctx, &protos.RateRequest{Base: "INR", Destination: "USD"})

	// checking of network issues
	if err != nil {
		log.Fatal("Error seen", err)
	}

	// returning response
	log.Printf("details %s", res.String())

}
