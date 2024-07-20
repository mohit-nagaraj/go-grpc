package main

import (
	"io"
	"log"

	pb "github.com/mohit-nagaraj/go-grpc/proto"
)

func (s *helloServer) SayHelloClientStream(stream pb.GreetService_SayHelloClientStreamServer) error {
	//receive the stream from client while it is open
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//send the response back to the client after receiving all the messages
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name : %v", req.Name)
		//append the message to the list
		messages = append(messages, "Hello "+req.Name)
	}

}
