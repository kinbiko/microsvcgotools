package main

import (
	"context"
	"log"
	"net"

	"net/http"
	_ "net/http/pprof"

	"github.com/kinbiko/microsvcgotools/fizzbuzz"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	fizzbuzz.RegisterFizzBuzzServer(grpcServer, &gRPCServer{})
	log.Println("starting http profiling endpoint")
	go http.ListenAndServe(":4321", nil)
	log.Println("starting gRPC server")
	grpcServer.Serve(lis)
}

type gRPCServer struct {
}

func (s *gRPCServer) GetFizzBuzzSequence(ctx context.Context, req *fizzbuzz.FizzBuzzRequest) (*fizzbuzz.FizzBuzzResponse, error) {
	res := []string{}
	for i := req.Start; i <= req.End; i++ {
		res = append(res, fizzbuzz.Fizzbuzz(i))
	}
	log.Printf("Serving sequence of length %d", len(res))
	return &fizzbuzz.FizzBuzzResponse{Success: true, Sequence: res}, nil
}
