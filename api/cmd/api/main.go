package main

import (
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	"github.com/tt/g3/api"
	"github.com/tt/g3/eventbus"
)

func main() {
	errs := make(chan error)

	client, err := eventbus.Dial(os.Getenv("EVENTBUS_ADDR"))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	go func() {
		defer client.Conn.Close()

		err = api.Listen(client)
		errs <- err
	}()

	h := handler.New(&handler.Config{
		GraphiQL: true,
		Pretty:   true,
		Schema:   api.Schema,
	})

	go func() {
		errs <- http.ListenAndServe(":8080", h)
	}()

	err = <-errs
	if err != nil {
		log.Fatalf("failed to listen and/or serve: %v", err)
	}
}
