package main

import (
	"context"

	pb "github.com/mohit-nagaraj/go-grpc/proto"
)

// take in the context and the request and return the response
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
