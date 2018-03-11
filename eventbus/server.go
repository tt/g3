package eventbus

import (
	"context"
	"strconv"
	"sync"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/tt/g3/eventbus/internal/pb"
	"google.golang.org/grpc"
)

func NewServer() *grpc.Server {
	stream := newStream()

	s := grpc.NewServer()
	pb.RegisterEventBusServer(s, &server{stream: stream})

	return s
}

type server struct {
	stream *stream
}

func (s *server) Publish(ctx context.Context, in *pb.PublishRequest) (*empty.Empty, error) {
	s.stream.Write(in.Event)

	return &empty.Empty{}, nil
}

func (s *server) Subscribe(in *pb.SubscribeRequest, stream pb.EventBus_SubscribeServer) error {
	offset := in.Offset

	for {
		var any *any.Any
		var err error
		any, offset, err = s.stream.Read(offset)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.SubscribeResponse{Event: any, Offset: offset})
		if err != nil {
			return err
		}
	}
}

type stream struct {
	mu     sync.RWMutex
	cond   *sync.Cond
	events []*any.Any
}

func newStream() *stream {
	s := &stream{}
	s.cond = sync.NewCond(s.mu.RLocker())

	return s
}

func (s *stream) Write(event *any.Any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.events = append(s.events, event)
	s.cond.Broadcast()
}

func (s *stream) Read(offset string) (*any.Any, string, error) {
	i, err := strconv.Atoi(offset)
	if err != nil {
		return nil, "", err
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	for len(s.events) <= i {
		s.cond.Wait()
	}

	return s.events[i], strconv.Itoa(i + 1), nil
}
