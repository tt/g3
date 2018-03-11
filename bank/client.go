package bank

import (
	"google.golang.org/grpc"
)

type Client struct {
	Conn *grpc.ClientConn
}

func Dial(address string) (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{Conn: conn}, nil
}
