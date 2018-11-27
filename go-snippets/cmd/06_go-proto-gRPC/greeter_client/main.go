package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/nikhil-thomas/technik-dojo/go-snippets/cmd/06_go-proto-gRPC/helloworld"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:5000"
	defaultname = "World"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	name := defaultname

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s\n", r.Message)
}
