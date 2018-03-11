package eventbus

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/tt/g3/eventbus/internal/pb"
	"google.golang.org/grpc"
)

const EarliestOffset string = "0"

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

func (c *Client) Publish(event proto.Message) error {
	client := pb.NewEventBusClient(c.Conn)

	any, err := ptypes.MarshalAny(event)
	if err != nil {
		return err
	}

	request := &pb.PublishRequest{Event: any}

	_, err = client.Publish(context.Background(), request)
	return err
}

func (c *Client) Subscribe(offset string, handler func(proto.Message, string) error) error {
	client := pb.NewEventBusClient(c.Conn)

	request := &pb.SubscribeRequest{Offset: offset}

	stream, err := client.Subscribe(context.Background(), request)
	if err != nil {
		return err
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			return err
		}

		var any ptypes.DynamicAny
		err = ptypes.UnmarshalAny(reply.Event, &any)
		if err != nil {
			return err
		}

		err = handler(any.Message, reply.Offset)
		if err != nil {
			return err
		}
	}
}
