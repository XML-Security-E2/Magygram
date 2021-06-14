package model

import (
	"github.com/beevik/guid"
	"mime/multipart"
)

type Product struct {
	Id string `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Price float64 `json:"price"`
	ImageURL string `json:"imageUrl"`
	Quantity int `json:"quantity"`
}

type ProductRequest struct {
	Name  string `json:"name"`
	Price float64 `json:"price"`
	Image *multipart.FileHeader `json:"image"`
}

type ProductUpdateRequest struct {
	Name  string `json:"name"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
}

type ProductImageUpdateRequest struct {
	Image *multipart.FileHeader `json:"image"`
}

func NewProduct(productRequest *ProductRequest, imageUrl string) *Product {
	return &Product{
		Id:       guid.New().String(),
		Name:     productRequest.Name,
		Price:    productRequest.Price,
		ImageURL: imageUrl,
		Quantity: 0,
	}
}
