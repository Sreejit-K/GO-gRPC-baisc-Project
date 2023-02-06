package main

import (
	"context"

	"github.com/Sreejit-K/GO-gRPC-basic-Project/pb"
)

func (s *helloServer) GreetAPerson(ctx context.Context, req *pb.NoParams) (*pb.GreetingResponse, error) {
	return &pb.GreetingResponse{
		Msg: "Hello from the server, how can I help you.",
	}, nil
}
