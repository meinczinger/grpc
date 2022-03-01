package main

import (
	"context"
	"fmt"
	calculatorpb "grpc_basics/calculator/calculatorpb/go"
	"io"
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

	// doUnary(s)

	doServerStreaming(s)
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

func doServerStreaming(s calculatorpb.SumServiceClient) {
	var n int32 = 21012
	fmt.Printf("Factorizing the number %v ", n)
	req := &calculatorpb.FactorizeRequest{
		Number: n,
	}
	resStream, err := s.Factorize(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		fmt.Printf("Factor received: %v\n", msg.GetFactor())
	}
}
