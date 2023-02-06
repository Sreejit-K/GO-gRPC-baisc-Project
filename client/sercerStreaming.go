package main

import (
	"context"
	"io"
	"log"

	"github.com/Sreejit-K/GO-gRPC-basic-Project/pb"
)

func callHelloServerStreamServer(client pb.GreetingServiceClient, names *pb.ReqList) {
	log.Printf("Steaming has started")

	// setup the response
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalln("clouldnt connet: ", err)
	}
	// since the response is a stream we will loop till we recieve the EOF error
	// which is basically a end of the file error
	// there is nothing left to recive from the error
	// this error is used to check weather the steam has ended or not

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming: %v", err)
		}
		log.Println(message)
	}
	log.Println("Streaming has finished!! cheers!!")
}
