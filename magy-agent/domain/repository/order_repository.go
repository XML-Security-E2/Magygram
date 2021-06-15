package repository

import (
	"context"
	"magyAgent/domain/model"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order) (*model.Order, error)
}
