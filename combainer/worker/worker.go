package worker

import (
	"context"
	"fmt"

	"github.com/cocaine/cocaine-framework-go/cocaine"
)

var (
	ErrAppCall = fmt.Errorf("Application call error")
)

// Future represents asynchronous resutl of a work
type Future interface {
	Wait(ctx context.Context, result interface{}) error
}

type futureV11 struct {
	ch chan cocaine.ServiceResult
}

func (f futureV11) Wait(ctx context.Context, result interface{}) error {
	select {
	case res := <-f.ch:
		if res == nil {
			return ErrAppCall
		}
		if res.Err() != nil {
			return res.Err()
		}
		return res.Extract(result)
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Worker is an interface on top of Cocaine Workers
type Worker interface {
	Do(_ context.Context, name string, args ...interface{}) Future
	Footprint() string
	Close()
}

type workerV11 struct {
	*cocaine.Service
}

func NewSlave(app *cocaine.Service) Worker {
	return &workerV11{
		Service: app,
	}
}

func (s *workerV11) Close() {
	s.Service.Close()
}

func (s *workerV11) Do(todo context.Context, name string, args ...interface{}) Future {
	return futureV11{
		ch: s.Service.Call(name, args...),
	}
}

func (s *workerV11) Footprint() string {
	return s.Service.Endpoint.AsString()
}
