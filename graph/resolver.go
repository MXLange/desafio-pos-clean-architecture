package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

import (
	usecases "github.com/MXLange/desafio-pos-clean-architecture/internal/domain/order/use_cases"
	e "github.com/MXLange/desafio-pos-clean-architecture/internal/errors"
)

type Resolver struct {
	CreateOrderUseCase *usecases.CreateOrderUseCase
	ListOrdersUseCase  *usecases.ListOrdersUseCase
}

func NewResolver(createOrderUseCase *usecases.CreateOrderUseCase, listOrdersUseCase *usecases.ListOrdersUseCase) (*Resolver, error) {
	if createOrderUseCase == nil {
		return nil, e.ErrNilCreateOrderUseCase
	}
	if listOrdersUseCase == nil {
		return nil, e.ErrNilListOrdersUseCase
	}

	return &Resolver{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}, nil
}
