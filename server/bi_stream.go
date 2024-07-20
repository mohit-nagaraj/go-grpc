package main

import (
	"io"
	"log"
	"time"

	pb "github.com/mohit-nagaraj/go-grpc/proto"
)

func (s *helloServer) SayHelloBiDiStream(stream pb.GreetService_SayHelloBiDiStreamServer) error {
	for {
		req, err := stream.Recv()
		//if the stream is closed by the client
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
		log.Printf("Got request with name : %v", req.Name)
		//send the response back to the client in the stream
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}

	}
}
