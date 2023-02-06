package main

import (
	"log"
	"net"

	"github.com/Sreejit-K/GO-gRPC-basic-Project/pb"
	"google.golang.org/grpc"
)

const (
	port = ":9090" // declare the port that you want to operate on / server to run on
)

type helloServer struct {
	pb.GreetingServiceServer
}

func main() {

	// Listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	// Now we have the port we need to create the GRPC Server next

	grpcServer := grpc.NewServer()

	// we need to register the service that we have created in the proto file
	pb.RegisterGreetingServiceServer(grpcServer, &helloServer{})
	log.Println("Server has started on the port %v", lis.Addr().String())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start %v", err)
	}

}
