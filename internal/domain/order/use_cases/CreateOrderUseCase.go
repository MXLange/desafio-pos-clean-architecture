package usecases

import (
	"context"

	"github.com/MXLange/desafio-pos-clean-architecture/internal/domain/order/dto"
	"github.com/MXLange/desafio-pos-clean-architecture/internal/domain/order/mapper"
	"github.com/MXLange/desafio-pos-clean-architecture/internal/domain/order/repository"
	e "github.com/MXLange/desafio-pos-clean-architecture/internal/errors"
	"github.com/MXLange/desafio-pos-clean-architecture/internal/logger"
)

type CreateOrderUseCase struct {
	orderRepository repository.OrderIF
	logger          logger.LoggerIF
}

func NewCreateOrderUseCase(orderRepo repository.OrderIF, logger logger.LoggerIF) (*CreateOrderUseCase, error) {
	if orderRepo == nil {
		return nil, e.ErrNilRepository
	}
	if logger == nil {
		return nil, e.ErrNilLogger
	}
	return &CreateOrderUseCase{
		orderRepository: orderRepo,
		logger:          logger,
	}, nil
}

func (uc *CreateOrderUseCase) Execute(ctx context.Context, d *dto.OrderCreateRequest) (*dto.OrderCreateResponse, error) {

	uc.logger.Info(ctx, "CreateOrderUseCase - Execute started")

	order, err := uc.orderRepository.CreateOrder(ctx, d.ProductID, d.Quantity)
	if err != nil {
		uc.logger.Errorf(ctx, "CreateOrderUseCase - failed to create order: %v", err)
		return nil, err
	}
	uc.logger.Infof(ctx, "CreateOrderUseCase - order created successfully: %v", order.ID)
	return mapper.EntityToOrderCreateResponse(order), nil
}
