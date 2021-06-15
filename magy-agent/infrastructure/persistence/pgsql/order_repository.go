package pgsql

import (
	"context"
	"gorm.io/gorm"
	"magyAgent/domain/model"
	"magyAgent/domain/repository"
)

type orderRepository struct {
	Conn *gorm.DB
}

func NewOrderRepository(Conn *gorm.DB) repository.OrderRepository {
	return &orderRepository{Conn}
}

func (p orderRepository) Create(ctx context.Context, order *model.Order) (*model.Order, error) {
	err := p.Conn.Create(order).Error
	return order, err
}
