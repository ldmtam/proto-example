package greeting

import (
	"context"
	"errors"

	greetingv1 "github.com/ldmtam/proto-example/proto/greeting/v1"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SayHello(ctx context.Context, in *greetingv1.SayHelloRequest) (*greetingv1.SayHelloResponse, error) {
	if len(in.Name) == 0 {
		return nil, errors.New("name cannot be empty")
	}

	return &greetingv1.SayHelloResponse{Reply: "Hello " + in.Name}, nil
}
