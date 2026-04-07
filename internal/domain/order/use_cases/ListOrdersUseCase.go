package usecases

import (
	"context"

	"github.com/MXLange/desafio-pos-clean-architecture/internal/domain/order/dto"
	"github.com/MXLange/desafio-pos-clean-architecture/internal/domain/order/mapper"
	"github.com/MXLange/desafio-pos-clean-architecture/internal/domain/order/repository"
	e "github.com/MXLange/desafio-pos-clean-architecture/internal/errors"
	"github.com/MXLange/desafio-pos-clean-architecture/internal/logger"
)

type ListOrdersUseCase struct {
	orderRepository repository.OrderIF
	logger          logger.LoggerIF
}

func NewListOrdersUseCase(orderRepo repository.OrderIF, logger logger.LoggerIF) (*ListOrdersUseCase, error) {
	if orderRepo == nil {
		return nil, e.ErrNilRepository
	}
	if logger == nil {
		return nil, e.ErrNilLogger
	}
	return &ListOrdersUseCase{
		orderRepository: orderRepo,
		logger:          logger,
	}, nil
}

func (uc *ListOrdersUseCase) Execute(ctx context.Context) ([]dto.Order, error) {

	uc.logger.Info(ctx, "ListOrdersUseCase - Execute started")

	orders, err := uc.orderRepository.ListOrders(ctx)
	if err != nil {
		uc.logger.Errorf(ctx, "ListOrdersUseCase - failed to list orders: %v", err)
		return nil, err
	}
	uc.logger.Infof(ctx, "ListOrdersUseCase - orders listed successfully: %v", orders)
	return mapper.EntityListToOrderListResponse(orders), nil
}
