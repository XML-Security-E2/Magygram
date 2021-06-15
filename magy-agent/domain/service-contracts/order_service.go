package service_contracts

import (
	"context"
	"magyAgent/domain/model"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *model.OrderRequest) (*model.Order, error)
}
