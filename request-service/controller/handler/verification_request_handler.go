package handler

import (
	"context"
	"github.com/labstack/echo"
	"mime/multipart"
	"net/http"
	"request-service/domain/model"
	"request-service/domain/service-contracts"
)

type VerificationRequestHandler interface {
	CreateVerificationRequest(c echo.Context) error
	CreateReportRequest(c echo.Context) error
}

type verificationRequestHandler struct {
	VerificationRequestService service_contracts.VerificationRequestService
}

func NewVerificationRequestHandler(u service_contracts.VerificationRequestService) VerificationRequestHandler {
	return &verificationRequestHandler{u}
}

func (v verificationRequestHandler) CreateVerificationRequest(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	mpf, _ := c.MultipartForm()
	var headers []*multipart.FileHeader
	for _, v := range mpf.File {
		headers = append(headers, v[0])
	}

	var formValues = mpf.Value

	var verificationRequestDTO = model.VerificationRequestDTO{
		Name: formValues["name"][0],
		Surname: formValues["surname"][0],
		Category: formValues["category"][0],
	}

	bearer := c.Request().Header.Get("Authorization")

	request, err := v.VerificationRequestService.CreateVerificationRequest(ctx, verificationRequestDTO, bearer ,headers)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, request)
}

func (v verificationRequestHandler) CreateReportRequest(c echo.Context) error {
	reportRequest := &model.ReportRequestDTO{}
	if err := c.Bind(reportRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	request, err := v.VerificationRequestService.CreateReportRequest(ctx, reportRequest)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, request)
}