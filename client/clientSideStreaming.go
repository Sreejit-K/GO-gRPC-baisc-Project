package main

import (
	"context"
	"log"
	"time"

	"github.com/Sreejit-K/GO-gRPC-basic-Project/pb"
)

func callHelloClientSideSteaming(client pb.GreetingServiceClient, req *pb.ReqList) {
	log.Println("Client steaming has started!!")

	stream, err := client.SayHelloClientStreamin(context.Background())

	if err != nil {
		log.Fatalf("couldnt send the names: %v", req)
	}

	for _, name := range req.Requests {
		reqName := &pb.GreetingRequest{
			Message: name,
		}

		if err := stream.Send(reqName); err != nil {
			log.Fatalf("error while streaming: %s", err)
		}

		log.Printf("sent the requset with name : %v", name)
		time.Sleep(2 * time.Second)

	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming finished!!")
	if err != nil {
		log.Fatalf("error while receving %v", err)
	}
	log.Printf("This is the message form the server at once : %v", res.Responses)
}
