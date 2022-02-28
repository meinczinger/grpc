package main

import (
	"context"
	"fmt"
	calculatorpb "grpc_basics/calculator/calculatorpb/go"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedSumServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum service was invoked: %v", req)

	num1 := req.GetNum1()
	num2 := req.GetNum2()
	sum := num1 + num2

	res := &calculatorpb.SumResponse{
		Sum: sum,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
