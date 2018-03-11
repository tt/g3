package bank

import (
	"context"

	"github.com/tt/g3/bank/internal/pb"
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

func (c *Client) OpenAccount() (string, error) {
	client := pb.NewBankClient(c.Conn)

	request := &pb.OpenAccountRequest{}

	response, err := client.OpenAccount(context.Background(), request)
	if err != nil {
		return "", err
	}

	return response.Id, nil
}
