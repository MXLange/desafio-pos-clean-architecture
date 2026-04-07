package errors

import (
	"errors"
)

var (
	ErrNilLogger             = errors.New("logger is nil")
	ErrNilHandler            = errors.New("handler is nil")
	ErrNilMux                = errors.New("mux is nil")
	ErrNilService            = errors.New("service is nil")
	ErrNilRepository         = errors.New("repository is nil")
	ErrNilDB                 = errors.New("db is nil")
	ErrNilCreateOrderUseCase = errors.New("create order use case is nil")
	ErrNilListOrdersUseCase  = errors.New("list orders use case is nil")
)
