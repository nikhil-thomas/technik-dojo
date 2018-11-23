package main

import (
	"context"
	"log"
	"net"

	"github.com/nikhil-thomas/technik-dojo/go-snippets/cmd/06_go-proto-gRPC/helloworld"

	"google.golang.org/grpc"
)

const address = ":5000"

// server is used to implement helloworld.GreeterServer
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello, " + in.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
