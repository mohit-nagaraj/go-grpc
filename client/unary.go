package main

import (
	"context"
	"log"
	"time"

	pb "github.com/mohit-nagaraj/go-grpc/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	/*  creates a context using the context.WithTimeout function. The context package in Go provides a way to manage and propagate cancellation signals and deadlines across goroutines (concurrent functions). In this case, the context is created with a timeout of one second*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	/* After setting up the context, the code calls a function SayHello on a client object, passing in the created context and a NoParam object as arguments. The SayHello function is expected to return a response (res) and an error (err). */
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
