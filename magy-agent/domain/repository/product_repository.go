package repository

import (
	"context"
	"magyAgent/domain/model"
)

type ProductRepository interface {
	Create(ctx context.Context, product *model.Product) (*model.Product, error)
	Update(ctx context.Context, product *model.Product) (*model.Product, error)
	GetById(ctx context.Context, id string) (*model.Product, error)
	DeleteById(ctx context.Context, id string) error
	GetAll(ctx context.Context) (*[]model.Product, error)
}