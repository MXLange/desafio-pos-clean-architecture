package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/MXLange/desafio-pos-clean-architecture/env"
	"github.com/MXLange/desafio-pos-clean-architecture/internal/domain/order/repository"
	usecases "github.com/MXLange/desafio-pos-clean-architecture/internal/domain/order/use_cases"
	"github.com/MXLange/desafio-pos-clean-architecture/internal/infra/db"
	"github.com/MXLange/desafio-pos-clean-architecture/internal/logger"
	"github.com/MXLange/desafio-pos-clean-architecture/internal/servers/rest"
)

func main() {
	ctx := context.Background()
	l := logger.NewLogger()

	l.Info(ctx, "Starting server...")

	e, err := env.New(ctx, l)
	if err != nil {
		l.Errorf(ctx, "Failed to load environment variables: %v", err)
		return
	}

	oDB := "./orders.db"

	d, err := db.NewConnection(ctx, "sqlite3", oDB)
	if err != nil {
		l.Panicf(ctx, "Failed to connect to the database: %v", err)
		return
	}

	err = db.Migrate("internal/infra/db/migrations", oDB)
	if err != nil {
		l.Panicf(ctx, "Failed to migrate database: %v", err)
		return
	}

	orderRepo, err := repository.NewOrderRepo(d)
	if err != nil {
		l.Panicf(ctx, "Failed to create order repository: %v", err)
		return
	}

	createOrderUseCase, err := usecases.NewCreateOrderUseCase(orderRepo, l)
	if err != nil {
		l.Panicf(ctx, "Failed to create create order use case: %v", err)
		return
	}
	listOrdersUseCase, err := usecases.NewListOrdersUseCase(orderRepo, l)
	if err != nil {
		l.Panicf(ctx, "Failed to create list orders use case: %v", err)
		return
	}

	handlers, err := rest.NewHandler(createOrderUseCase, listOrdersUseCase)
	if err != nil {
		l.Panicf(ctx, "Failed to create handlers: %v", err)
		return
	}
	restServer, err := rest.NewServer(e.RestPort, l, handlers)
	if err != nil {
		l.Panicf(ctx, "Failed to create REST server: %v", err)
		return
	}
	if err := restServer.Start(ctx); err != nil {
		l.Panicf(ctx, "Failed to start REST server: %v", err)
		return
	}

	shutdownCtx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-shutdownCtx.Done()

	l.Info(ctx, "Shutdown signal received")

	stopCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := restServer.Stop(stopCtx); err != nil {
		l.Errorf(ctx, "Failed to stop REST server: %v", err)
	}

	l.Info(ctx, "Server stopped")
}
