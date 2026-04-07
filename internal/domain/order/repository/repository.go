package repository

import (
	"context"
	"database/sql"

	e "github.com/MXLange/desafio-pos-clean-architecture/internal/errors"
)

type OderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) (*OderRepo, error) {
	if db == nil {
		return nil, e.ErrNilDB
	}

	return &OderRepo{db: db}, nil
}

func (r *OderRepo) CreateOrder(ctx context.Context, productId uint, quantity uint) (*Order, error) {
	var order Order

	err := r.db.QueryRowContext(ctx,
		"INSERT INTO orders (product_id, quantity) VALUES ($1, $2) RETURNING id, product_id, quantity",
		productId,
		quantity,
	).Scan(&order.ID, &order.ProductId, &order.Quantity)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OderRepo) ListOrders(ctx context.Context) ([]Order, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, product_id, quantity FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order

	for rows.Next() {
		var order Order
		err := rows.Scan(&order.ID, &order.ProductId, &order.Quantity)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
