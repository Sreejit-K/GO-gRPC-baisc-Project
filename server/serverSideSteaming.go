package main

import (
	"log"
	"time"

	"github.com/Sreejit-K/GO-gRPC-basic-Project/pb"
)

func (*helloServer) SayHelloServerStreaming(req *pb.ReqList, stream pb.GreetingService_SayHelloServerStreamingServer) error {
	log.Println("have recived this req from the clent it consists of list of names : %v", req.Requests)

	for _, name := range req.Requests {
		res := &pb.GreetingResponse{
			Msg: "Hello" + name,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
