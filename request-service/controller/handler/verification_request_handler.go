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
	GetVerificationRequests(c echo.Context) error
	ApproveVerificationRequest(c echo.Context) error
	RejectVerificationRequest(c echo.Context) error
	HasUserPendingRequest(c echo.Context) error
	GetReportRequests(c echo.Context) error
	DeleteReportRequest(c echo.Context) error
	GetCampaignRequests(c echo.Context) error
	DeleteCampaignRequest(c echo.Context) error
	CreateCampaignRequest(c echo.Context) error
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

func (v verificationRequestHandler) CreateCampaignRequest(c echo.Context) error {
	campaignRequest := &model.CampaignRequestDTO{}
	if err := c.Bind(campaignRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	bearer := c.Request().Header.Get("Authorization")

	request, err := v.VerificationRequestService.CreateCampaignRequest(ctx, bearer, campaignRequest)

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
	bearer := c.Request().Header.Get("Authorization")

	request, err := v.VerificationRequestService.CreateReportRequest(ctx, bearer, reportRequest)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, request)
}

func (v verificationRequestHandler) GetVerificationRequests(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	retVal, err := v.VerificationRequestService.GetVerificationRequests(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, retVal)
}

func (v verificationRequestHandler) GetCampaignRequests(c echo.Context) error {
	postId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	retVal, err := v.VerificationRequestService.GetCampaignRequests(ctx, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, retVal)
}

func (v verificationRequestHandler) GetReportRequests(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	retVal, err := v.VerificationRequestService.GetReportRequests(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, retVal)
}

func (v verificationRequestHandler) DeleteCampaignRequest(c echo.Context) error {
	postId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := v.VerificationRequestService.DeleteCampaignRequest(ctx, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (v verificationRequestHandler) DeleteReportRequest(c echo.Context) error {
	postId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := v.VerificationRequestService.DeleteReportRequest(ctx, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (v verificationRequestHandler) ApproveVerificationRequest(c echo.Context) error {
	postId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := v.VerificationRequestService.ApproveVerificationRequest(ctx, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (v verificationRequestHandler) RejectVerificationRequest(c echo.Context) error {
	postId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := v.VerificationRequestService.RejectVerificationRequest(ctx, postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (v verificationRequestHandler) HasUserPendingRequest(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	bearer := c.Request().Header.Get("Authorization")

	result,err := v.VerificationRequestService.HasUserPendingRequest(ctx, bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}