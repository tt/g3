FROM golang:1.11
WORKDIR /go/src/github.com/tt/g3
COPY . .
RUN go install -v ./...
EXPOSE 8080
CMD ["/go/bin/api"]
