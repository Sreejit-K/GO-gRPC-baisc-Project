package main

import (
	"io"
	"log"

	"github.com/Sreejit-K/GO-gRPC-basic-Project/pb"
)

func (*helloServer) SayHelloClientStreamin(stream pb.GreetingService_SayHelloClientStreaminServer) error {

	var messages []string

	for {
		req, err := stream.Recv()

		if err == io.EOF {

			return stream.SendAndClose(&pb.ResList{Responses: messages})
		}
		if err != nil {
			log.Fatalf("There is some kind of an error in this stream: %v", err)
		}
		log.Printf("This is the requset that has been received: %v", req.Message)
		messages = append(messages, "Hello ", req.Message)
	}
}
