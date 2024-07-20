package main

import (
	"context"
	"io"
	"log"

	pb "github.com/mohit-nagaraj/go-grpc/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	//receive the names and send a msg to server and get a stream back
	stream, err := client.SayHelloStream(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for {
		//receive the stream from server while it is open
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming %v", err)
		}
		log.Println(message)
	}

	log.Printf("Streaming finished")
}
