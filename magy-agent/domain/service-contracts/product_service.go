package service_contracts

import (
	"context"
	"magyAgent/domain/model"
	"mime/multipart"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product *model.ProductRequest) (*model.Product, error)
	UpdateProduct(ctx context.Context, productId string, product *model.ProductUpdateRequest) (*model.Product, error)
	UpdateProductImage(ctx context.Context, productId string, image *multipart.FileHeader) (*model.Product, error)
	GetProductById(ctx context.Context, id string) (*model.Product, error)
	DeleteProductById(ctx context.Context, id string) error
	GetAllProducts(ctx context.Context) (*[]model.Product, error)
	CreateProductCampaign(ctx context.Context, campaignReq *model.CampaignRequest) error
	GetProductCampaignStatistics(ctx context.Context) (*model.CampaignStatisticReport, error)

	GetAllProductCampaignStatisticsReports(ctx context.Context) ([]*model.CampaignStatisticReport, error)
	GetDocumentByIdInPdf(ctx context.Context, filename string) ([]byte, error)
}