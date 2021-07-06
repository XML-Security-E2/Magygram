package service

import (
	"context"
	"errors"
	"github.com/beevik/guid"
	"magyAgent/domain/model"
	"magyAgent/domain/repository"
	"magyAgent/domain/service-contracts"
	"time"
)

type orderService struct {
	repository.OrderRepository
	repository.ProductRepository
}

func NewOrderService(r repository.OrderRepository, pr repository.ProductRepository) service_contracts.OrderService {
	return &orderService{r, pr }
}

func (o orderService) CreateOrder(ctx context.Context, orderReq *model.OrderRequest) (*model.Order, error) {
	order := &model.Order{
		Id:        guid.New().String(),
		Address:   orderReq.Address,
		Timestamp: time.Now(),
		Items:     []*model.OrderItem{},
	}
	if len(orderReq.Items) == 0 {
		return nil, errors.New("at least one product must be selected")
	}

	for _, item := range orderReq.Items {
		product, err := o.ProductRepository.GetById(ctx, item.ProductId)
		if err != nil || product == nil {
			return nil, errors.New("invalid product id")
		}

		if product.Quantity < item.Count {
			return nil, errors.New("product " + product.Name + " quantity not available")
		}

		order.Items = append(order.Items, &model.OrderItem{
			OrderId:         order.Id,
			ProductId:       product.Id,
			ProductName:     product.Name,
			Count:           item.Count,
			PricePerProduct: product.Price,
		})
	}

	for _, item := range orderReq.Items {
		product, err := o.ProductRepository.GetById(ctx, item.ProductId)
		if err != nil || product == nil {
			return nil, errors.New("invalid product id")
		}

		product.Quantity = product.Quantity - item.Count
		o.ProductRepository.Update(ctx, product)
	}

	return o.OrderRepository.Create(ctx, order)
}