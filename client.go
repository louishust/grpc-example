package main

import (
	"log"

	pb "github.com/louishust/grpc-example/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	h := pb.NewHeartClient(conn)
	q, err := h.HeartBeat(context.Background(), &pb.HeartRequest{})
	if err != nil {
		log.Fatalf("could not heartbeat: %v", err)
	}
	log.Printf("HeartBeat: %s", q.Message)

}
