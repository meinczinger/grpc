package main

import (
	"context"
	"fmt"
	calculatorpb "grpc_basics/calculator/calculatorpb/go"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	// Execute at the end
	defer cc.Close()

	s := calculatorpb.NewCalculatorServiceClient(cc)

	// doUnary(s)

	// doServerStreaming(s)

	doClientStreaming(s)
}

func doUnary(s calculatorpb.CalculatorServiceClient) {
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

func doServerStreaming(s calculatorpb.CalculatorServiceClient) {
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

func doClientStreaming(s calculatorpb.CalculatorServiceClient) {
	fmt.Println("Calculating average")

	numbers := []int32{2, 5, 91, 12, 43}

	stream, err := s.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Average: %v", err)
	}

	for _, number := range numbers {
		fmt.Printf("Sending req: %v\n", number)
		stream.Send(&calculatorpb.AverageRequest{
			Number: number,
		})
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while recieving response from Average: %v", err)
	}

	fmt.Printf("Average: %v \n", res.GetAverage())

}
