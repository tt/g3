.PHONY: build generate

build:
	go build -v ./...

generate:
	go generate -v ./...
