package main

import (
	"context"
	"grpc_learning/protos"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051" // target port
)

type CurrencyServer struct { // embedding the service
	protos.UnimplementedCurrencyServer
}

// rpc functions must match, with context comes greater responsibilty
func (p *CurrencyServer) GetRate(ctx context.Context, request *protos.RateRequest) (*protos.RateResponse, error) {
	log.Printf("Received %s", (*request).String())
	return &protos.RateResponse{Rate: 34}, nil
}

func main() {

	// listening to port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listed")
	}
	// new server object
	serv := grpc.NewServer()

	// regstering service
	protos.RegisterCurrencyServer(serv, &CurrencyServer{})

	log.Printf("server listening to %v", lis.Addr())

	// initiating server
	if err := serv.Serve(lis); err != nil {
		log.Fatal("Failed to serve")
	}

}
