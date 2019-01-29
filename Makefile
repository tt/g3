.PHONY: build generate

build:
	go build -v ./...

protoc-gen-go:
	go install ./vendor/github.com/golang/protobuf/protoc-gen-go

generate: protoc-gen-go
	go generate -v ./...
