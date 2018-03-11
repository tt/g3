package main

import (
	"log"
	"net"

	"github.com/tt/g3/bank/internal/pb"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterBankServer(s, &server{})

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct{}
