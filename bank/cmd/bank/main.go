package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/tt/g3/bank/internal/pb"
	"github.com/tt/g3/eventbus"
	"github.com/tt/g3/events"
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

func (*server) OpenAccount(ctx context.Context, in *pb.OpenAccountRequest) (*pb.OpenAccountResponse, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	event := &events.AccountOpened{Id: id.String()}

	err = publish(event)
	if err != nil {
		return nil, err
	}

	return &pb.OpenAccountResponse{Id: id.String()}, nil
}

func publish(event proto.Message) error {
	client, err := eventbus.Dial(os.Getenv("EVENTBUS_ADDR"))
	if err != nil {
		return err
	}

	defer client.Conn.Close()

	return client.Publish(event)
}
