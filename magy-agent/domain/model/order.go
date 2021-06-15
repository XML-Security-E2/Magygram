package model

import (
	"time"
)

type Order struct {
	Id string `gorm:"primaryKey" json:"id"`
	Address  string `gorm:"not null" json:"address"`
	Timestamp time.Time `gorm:"not null" json:"time"`
	Items []*OrderItem `gorm:"foreignKey:OrderId;references:Id" json:"items"`
}

type OrderItem struct {
	OrderId string `gorm:"primaryKey" json:"orderId"`
	ProductId string `gorm:"primaryKey" json:"productId"`
	ProductName string `json:"productName"`
	Count int `json:"count"`
	PricePerProduct float64 `json:"pricePerProduct"`
}

type OrderRequest struct {
	Address  string `json:"address"`
	Items []OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	ProductId string `json:"productId"`
	Count int `json:"count"`
}