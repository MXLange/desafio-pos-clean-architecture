package rest

import (
	"context"
	"net/http"

	e "github.com/MXLange/desafio-pos-clean-architecture/internal/errors"

	"github.com/MXLange/desafio-pos-clean-architecture/internal/logger"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	port     string
	logger   logger.LoggerIF
	handlers *Handlers
	server   *http.Server
}

func NewServer(port string, logger logger.LoggerIF, handlers *Handlers) (*Server, error) {
	if logger == nil {
		return nil, e.ErrNilLogger
	}
	if handlers == nil {
		return nil, e.ErrNilHandler
	}
	return &Server{
		port:     port,
		logger:   logger,
		handlers: handlers,
	}, nil
}

func (s *Server) Start(ctx context.Context) error {

	s.logger.Infof(ctx, "Starting REST server on port %s", s.port)

	r := chi.NewRouter()

	r.Get("/orders", s.handlers.ListOrders)
	r.Post("/orders", s.handlers.CreateOrder)

	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: r,
	}

	go func() {
		s.logger.Infof(ctx, "REST server is listening on port %s", s.port)
		if err := s.server.ListenAndServe(); err != nil {
			ctx.Done()
		}
	}()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Infof(ctx, "Stopping REST server on port %s", s.port)
	if s.server == nil {
		return nil
	}
	return s.server.Shutdown(ctx)
}
