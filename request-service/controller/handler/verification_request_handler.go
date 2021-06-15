package handler

import (
	"context"
	"github.com/labstack/echo"
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
	verificationRequest := &model.VerificationRequestDTO{}
	if err := c.Bind(verificationRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	request, err := v.VerificationRequestService.CreateVerificationRequest(ctx, verificationRequest)

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