package main

import (
	"context"
	"fmt"
	calculatorpb "grpc_basics/calculator/calculatorpb/go"
	"io"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
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

func (*server) Factorize(req *calculatorpb.FactorizeRequest, stream calculatorpb.CalculatorService_FactorizeServer) error {
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

func (*server) Average(stream calculatorpb.CalculatorService_AverageServer) error {
	fmt.Println("Calculating average...")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.AverageResponse{
				Average: float32(sum) / float32(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		number := req.GetNumber()
		sum += number
		count += 1
	}
}

func (*server) Maximum(stream calculatorpb.CalculatorService_MaximumServer) error {
	fmt.Println("Calculating maximum...")

	var maximum int32 = math.MinInt32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Got EOF")
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		number := req.GetNumber()
		if number > maximum {
			maximum = number
			stream.Send(&calculatorpb.MaximumResponse{
				Maximum: maximum,
			})
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
