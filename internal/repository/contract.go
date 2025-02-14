package repository

import (
	"context"
	"go-boilerplate/internal/repository/order"
)

type OrderRepositoryInterface interface {
	Save(ctx context.Context, order order.Order) (string, error)
}
