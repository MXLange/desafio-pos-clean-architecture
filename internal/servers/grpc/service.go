package grpcserver

import (
	"context"

	"github.com/MXLange/desafio-pos-clean-architecture/internal/domain/order/dto"
	usecases "github.com/MXLange/desafio-pos-clean-architecture/internal/domain/order/use_cases"
	e "github.com/MXLange/desafio-pos-clean-architecture/internal/errors"
	orderpb "github.com/MXLange/desafio-pos-clean-architecture/proto/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderService struct {
	orderpb.UnimplementedOrderServiceServer
	createOrderUseCase *usecases.CreateOrderUseCase
	listOrdersUseCase  *usecases.ListOrdersUseCase
}

func NewOrderService(createOrderUseCase *usecases.CreateOrderUseCase, listOrdersUseCase *usecases.ListOrdersUseCase) (*OrderService, error) {
	if createOrderUseCase == nil {
		return nil, e.ErrNilCreateOrderUseCase
	}
	if listOrdersUseCase == nil {
		return nil, e.ErrNilListOrdersUseCase
	}

	return &OrderService{
		createOrderUseCase: createOrderUseCase,
		listOrdersUseCase:  listOrdersUseCase,
	}, nil
}

func (s *OrderService) CreateOrder(ctx context.Context, req *orderpb.NewOrder) (*orderpb.Order, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is required")
	}

	res, err := s.createOrderUseCase.Execute(ctx, &dto.OrderCreateRequest{
		ProductID: uint(req.GetProductId()),
		Quantity:  uint(req.GetQuantity()),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create order: %v", err)
	}

	return &orderpb.Order{
		Id:        uint64(res.ID),
		ProductId: uint64(res.ProductID),
		Quantity:  uint64(res.Quantity),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, _ *emptypb.Empty) (*orderpb.OrderList, error) {
	res, err := s.listOrdersUseCase.Execute(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list orders: %v", err)
	}

	orders := make([]*orderpb.Order, 0, len(res))
	for _, order := range res {
		orders = append(orders, &orderpb.Order{
			Id:        uint64(order.ID),
			ProductId: uint64(order.ProductID),
			Quantity:  uint64(order.Quantity),
		})
	}

	return &orderpb.OrderList{Orders: orders}, nil
}
