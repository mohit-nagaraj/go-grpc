package main

import (
	"log"
	"time"

	pb "github.com/mohit-nagaraj/go-grpc/proto"
)

// SayHelloStream shud match in client side calling
// GreetService_SayHelloStreamServer is used since its at server
func (s *helloServer) SayHelloStream(req *pb.NamesList, stream pb.GreetService_SayHelloStreamServer) error {
	log.Printf("Got request with names : %v", req.Names)
	for _, name := range req.Names {
		//simply send a response to each of the names
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		//send the response in the stream
		if err := stream.Send(res); err != nil {
			return err
		}
		// 2 second delay to simulate a long running process
		time.Sleep(2 * time.Second)
	}
	return nil
}
