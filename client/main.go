package main

import (
	"log"

	"github.com/Sreejit-K/GO-gRPC-basic-Project/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":9090" // port for the client to listen on
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("didnt connect to server %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetingServiceClient(conn)

	// names := &pb.ReqList{
	// 	Names: []string{"Sreejith", "Alice", "Durgesh"},
	// }

	callSayHello(client)
}
