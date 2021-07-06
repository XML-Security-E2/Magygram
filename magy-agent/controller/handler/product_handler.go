package handler

import (
	"context"
	"github.com/labstack/echo"
	"magyAgent/domain/model"
	"magyAgent/domain/service-contracts"
	"net/http"
	"strconv"
)

type ProductHandler interface {
	CreateProduct(c echo.Context) error
	UpdateProduct(c echo.Context) error
	UpdateProductImage(c echo.Context) error
	GetProductById(c echo.Context) error
	DeleteProductById(c echo.Context) error
	GetAllProducts(c echo.Context) error
	CreateProductCampaign(c echo.Context) error
}

func NewProductHandler(a service_contracts.ProductService) ProductHandler {
	return &productHandler{a}
}

type productHandler struct {
	ProductService service_contracts.ProductService
}

func (p productHandler) CreateProductCampaign(c echo.Context) error {
	campaignReq := &model.CampaignRequest{}
	if err := c.Bind(campaignReq); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return c.JSON(http.StatusUnauthorized, "")
	}

	err := p.ProductService.CreateProductCampaign(ctx, campaignReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusCreated, "")
}

func (p productHandler) CreateProduct(c echo.Context) error {
	name := c.FormValue("name")
	price := c.FormValue("price")
	image, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pr, err := strconv.ParseFloat(price, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	productRequest := &model.ProductRequest{
		Name:  name,
		Price: pr,
		Image: image,
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return c.JSON(http.StatusUnauthorized, "")
	}

	product, err := p.ProductService.CreateProduct(ctx, productRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, product)
}

func (p productHandler) UpdateProduct(c echo.Context) error {
	productId := c.Param("productId")
	updateReq := &model.ProductUpdateRequest{}
	if err := c.Bind(updateReq); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return c.JSON(http.StatusUnauthorized, "")
	}

	product, err := p.ProductService.UpdateProduct(ctx, productId, updateReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, product)
}

func (p productHandler) UpdateProductImage(c echo.Context) error {
	productId := c.Param("productId")
	image, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return c.JSON(http.StatusUnauthorized, "")
	}

	product, err := p.ProductService.UpdateProductImage(ctx, productId, image)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product.ImageURL)
}

func (p productHandler) GetProductById(c echo.Context) error {
	productId := c.Param("productId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	product, err := p.ProductService.GetProductById(ctx, productId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

func (p productHandler) DeleteProductById(c echo.Context) error {
	productId := c.Param("productId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := p.ProductService.DeleteProductById(ctx, productId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, productId)
}

func (p productHandler) GetAllProducts(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	products, err := p.ProductService.GetAllProducts(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, products)
}

