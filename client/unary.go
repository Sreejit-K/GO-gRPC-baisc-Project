package main

import (
	"context"
	"log"
	"time"

	"github.com/Sreejit-K/GO-gRPC-basic-Project/pb"
)

func callSayHello(client pb.GreetingServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.GreetAPerson(ctx, &pb.NoParams{})
	if err != nil {
		log.Fatalln("Couldnt greet: %v", err)
	}

	log.Println(res.Msg)
}
