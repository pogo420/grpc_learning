package main

import (
	"encoding/json"
	"fmt"
	"grpc_learning/protos"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	request := &protos.RateRequest{Base: "INR", Destination: "USD"}

	rawData, err := proto.Marshal(request) // data that will be transferred via wire.

	if err != nil {
		log.Fatal("Issue seen", err)
	}
	fmt.Println("PROTO", rawData)

	rawData_, err_ := json.Marshal(request) // data that will be transferred via wire.

	if err_ != nil {
		log.Fatal("Issue seen", err)
	}
	fmt.Println("JSON:", rawData_)

	// proto raw to object
	data := &protos.RateRequest{}
	err_1 := proto.Unmarshal(rawData, data) // initialize the struct

	if err_1 != nil {
		log.Fatal("issue seen", err_1)
	}

	fmt.Println((*data).String())
}
