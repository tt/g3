package api

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/tt/g3/eventbus"
	"github.com/tt/g3/events"
)

func Listen(client *eventbus.Client) error {
	return client.Subscribe(eventbus.EarliestOffset, handle)
}

func handle(event proto.Message, offset string) error {
	switch t := event.(type) {
	case *events.AccountOpened:
		accountTable.Insert(Account{
			ID: t.Id,
		})
	default:
		return fmt.Errorf("unknown event: %T", t)
	}

	return nil
}
