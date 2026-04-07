package grpcserver

import (
	"context"
	"net"

	e "github.com/MXLange/desafio-pos-clean-architecture/internal/errors"
	"github.com/MXLange/desafio-pos-clean-architecture/internal/logger"
	orderpb "github.com/MXLange/desafio-pos-clean-architecture/proto/order"
	"google.golang.org/grpc"
)

type Server struct {
	port    string
	logger  logger.LoggerIF
	service *OrderService
	server  *grpc.Server
	lis     net.Listener
}

func NewServer(port string, logger logger.LoggerIF, service *OrderService) (*Server, error) {
	if logger == nil {
		return nil, e.ErrNilLogger
	}
	if service == nil {
		return nil, e.ErrNilService
	}

	return &Server{
		port:    port,
		logger:  logger,
		service: service,
	}, nil
}

func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return err
	}

	s.lis = lis
	s.server = grpc.NewServer()
	orderpb.RegisterOrderServiceServer(s.server, s.service)

	s.logger.Infof(ctx, "Starting gRPC server on port %s", s.port)

	go func() {
		s.logger.Infof(ctx, "gRPC server is listening on port %s", s.port)
		if err := s.server.Serve(s.lis); err != nil {
			s.logger.Errorf(ctx, "Failed to start gRPC server: %v", err)
		}
	}()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Infof(ctx, "Stopping gRPC server on port %s", s.port)
	if s.server == nil {
		return nil
	}

	done := make(chan struct{})
	go func() {
		s.server.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		s.server.Stop()
		return ctx.Err()
	}
}
