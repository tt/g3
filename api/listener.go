package api

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/tt/g3/eventbus"
)

func Listen(client *eventbus.Client) error {
	return client.Subscribe(eventbus.EarliestOffset, handle)
}

func handle(event proto.Message, offset string) error {
	switch t := event.(type) {
	default:
		return fmt.Errorf("unknown event: %T", t)
	}

	return nil
}
