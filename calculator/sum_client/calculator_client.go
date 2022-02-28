package main

import (
	"context"
	calculatorpb "grpc_basics/calculator/calculatorpb/go"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	// Execute at the end
	defer cc.Close()

	s := calculatorpb.NewSumServiceClient(cc)

	doUnary(s)
}

func doUnary(s calculatorpb.SumServiceClient) {
	req := &calculatorpb.SumRequest{
		Num1: 3,
		Num2: 15,
	}

	res, err := s.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Sum)
}
