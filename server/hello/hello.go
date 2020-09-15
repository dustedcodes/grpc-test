package hello

import (
	"context"
	"fmt"
)

// Service is the Hello GRPC service.
type Service struct{}

// SayHello says hello.
func (s *Service) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: fmt.Sprintf("%s", req.Name)}, nil
}
