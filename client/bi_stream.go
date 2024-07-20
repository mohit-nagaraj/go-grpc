package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/mohit-nagaraj/go-grpc/proto"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional Streaming started")
	stream, err := client.SayHelloBiDiStream(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}
	//waitc is a channel to wait for the stream to finish
	waitc := make(chan struct{})
	//receive the stream from server while it is open
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break //if the stream is closed by the server
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			//print the message received from the server
			log.Println(message)
		}
		close(waitc) //close the channel when the stream is closed
	}()

	for _, name := range names.Names {
		log.Printf("Sending request with name: %s", name)
		req := &pb.HelloRequest{
			Name: name, //send the names one by one
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(2 * time.Second) // 2 second delay to simulate a long running process
	}
	//close the stream after sending all the names
	stream.CloseSend()
	//wait for the stream to finish
	<-waitc
	log.Printf("Bidirectional Streaming finished")
}
