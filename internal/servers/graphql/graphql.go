package graphql

import (
	"context"
	"net/http"

	e "github.com/MXLange/desafio-pos-clean-architecture/internal/errors"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MXLange/desafio-pos-clean-architecture/graph"
	"github.com/MXLange/desafio-pos-clean-architecture/internal/logger"
	"github.com/vektah/gqlparser/v2/ast"
)

type Server struct {
	port     string
	logger   logger.LoggerIF
	resolver *graph.Resolver
	server   *handler.Server
}

func NewServer(port string, logger logger.LoggerIF, resolver *graph.Resolver) (*Server, error) {
	if logger == nil {
		return nil, e.ErrNilLogger
	}
	if resolver == nil {
		return nil, e.ErrNilService
	}
	return &Server{
		port:     port,
		logger:   logger,
		resolver: resolver,
	}, nil
}

func (s *Server) Start(ctx context.Context) {

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: s.resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	s.logger.Infof(ctx, "connect to http://localhost:%s/ for GraphQL playground", s.port)

	go func() {
		if err := http.ListenAndServe(":"+s.port, nil); err != nil {
			s.logger.Errorf(ctx, "Failed to start GraphQL server: %v", err)
			ctx.Done()
		}
	}()
}
