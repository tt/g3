package main

import (
	"log"
	"os"

	"github.com/tt/g3/api"
	"github.com/tt/g3/eventbus"
)

func main() {
	client, err := eventbus.Dial(os.Getenv("EVENTBUS_ADDR"))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer client.Conn.Close()

	err = api.Listen(client)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
