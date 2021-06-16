package pgsql

import (
	"context"
	"gorm.io/gorm"
	"magyAgent/domain/model"
	"magyAgent/domain/repository"
)

type productRepository struct {
	Conn *gorm.DB
}

func NewProductRepository(Conn *gorm.DB) repository.ProductRepository {
	return &productRepository{Conn}
}

func (p productRepository) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	err := p.Conn.Create(product).Error
	return product, err
}

func (p productRepository) Update(ctx context.Context, product *model.Product) (*model.Product, error) {
	err := p.Conn.Model(product).Updates(map[string]interface{}{
		"Name" : product.Name,
		"Price" : product.Price,
		"ImageURL" : product.ImageURL,
		"Quantity" : product.Quantity,
	}).Error
	return product, err
}

func (p productRepository) GetById(ctx context.Context, id string) (*model.Product, error) {
	u := &model.Product{Id: id}
	err := p.Conn.First(u).Error
	return u, err
}

func (p productRepository) GetAll(ctx context.Context) (*[]model.Product, error) {
	products := &[]model.Product{}
	err := p.Conn.Find(products).Error
	return products, err
}

func (p productRepository) DeleteById(ctx context.Context, id string) error {
	return p.Conn.Delete(&model.Product{Id: id}).Error
}