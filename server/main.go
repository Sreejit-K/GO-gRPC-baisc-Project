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
	pb.GreetingServiceServer // interface of the server you have designed in the proto files
}

func main() {

	// Listener or activate the port
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

// step 1 : decide a port
// step 2 : setup the port access using protos calls
// step 3 : register the services that you have defined in the proto files
// step 4 : use the Register functional from the genrater protobuffs files
// step 5 : define the server type to pass as reference to the register funtion
// step 6 : And you are go to go and make the logical changes in the functions that you have declared in the proto files
