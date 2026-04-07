package repository

import "context"

type OrderIF interface {
	CreateOrder(ctx context.Context, productId uint, quantity uint) (*Order, error)
	ListOrders(ctx context.Context) ([]Order, error)
}
