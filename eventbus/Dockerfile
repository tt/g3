FROM golang:1.11
WORKDIR /go/src/github.com/tt/g3
COPY . .
RUN go install -v ./...
EXPOSE 50051
CMD ["/go/bin/eventbus"]
