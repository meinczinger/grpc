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

func (*server) Factorize(req *calculatorpb.FactorizeRequest, stream calculatorpb.SumService_FactorizeServer) error {
	fmt.Printf("Function Factorize has been called with %v", req)

	number := req.GetNumber()

	// var n, k int32
	var n, k int32
	n = number
	k = 2
	for n > 1 {
		if (n % k) == 0 {
			fmt.Printf("A factor has been found: %v", k)
			n = n / k
			res := &calculatorpb.FactorResponse{
				Factor: k,
			}
			stream.Send(res)
		} else {
			k++
		}
	}
	return nil
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
