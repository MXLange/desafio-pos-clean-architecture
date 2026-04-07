package mapper

import (
	"github.com/MXLange/desafio-pos-clean-architecture/internal/domain/order/dto"
	"github.com/MXLange/desafio-pos-clean-architecture/internal/domain/order/repository"
)

func EntityToOrderCreateResponse(order *repository.Order) *dto.OrderCreateResponse {
	return &dto.OrderCreateResponse{
		ID:        order.ID,
		ProductID: order.ProductId,
		Quantity:  order.Quantity,
	}
}

func EntityListToOrderListResponse(orders []repository.Order) []dto.Order {
	orderListResponse := make([]dto.Order, 0, len(orders))

	for _, order := range orders {
		orderListResponse = append(orderListResponse, dto.Order{
			ID:        order.ID,
			ProductID: order.ProductId,
			Quantity:  order.Quantity,
		})
	}

	return orderListResponse
}
