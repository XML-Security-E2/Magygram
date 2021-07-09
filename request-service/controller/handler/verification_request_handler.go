package handler

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"request-service/domain/model"
	"request-service/tracer"

	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"
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
	tracer                     opentracing.Tracer
	closer                     io.Closer
}

func NewVerificationRequestHandler(u service_contracts.VerificationRequestService) VerificationRequestHandler {
	tracer, closer := tracer.Init("request-service")
	opentracing.SetGlobalTracer(tracer)
	return &verificationRequestHandler{u, tracer, closer}
}

func (v verificationRequestHandler) CreateVerificationRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("RequestHandlerCreateVerificationRequest", v.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling create verification request at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	mpf, _ := c.MultipartForm()
	var headers []*multipart.FileHeader
	for _, v := range mpf.File {
		headers = append(headers, v[0])
	}

	var formValues = mpf.Value

	var verificationRequestDTO = model.VerificationRequestDTO{
		Name:     formValues["name"][0],
		Surname:  formValues["surname"][0],
		Category: formValues["category"][0],
	}

	bearer := c.Request().Header.Get("Authorization")

	request, err := v.VerificationRequestService.CreateVerificationRequest(ctx, verificationRequestDTO, bearer, headers)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, request)
}

func (v verificationRequestHandler) CreateCampaignRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("RequestHandlerCreateCampaignRequest", v.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling create campaign request at %s\n", c.Path())),
	)

	campaignRequest := &model.CampaignRequestDTO{}
	if err := c.Bind(campaignRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	request, err := v.VerificationRequestService.CreateCampaignRequest(ctx, bearer, campaignRequest)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, request)
}

func (v verificationRequestHandler) CreateReportRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("RequestHandlerCreateReportRequest", v.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling create report request at %s\n", c.Path())),
	)

	reportRequest := &model.ReportRequestDTO{}
	if err := c.Bind(reportRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)
	bearer := c.Request().Header.Get("Authorization")

	request, err := v.VerificationRequestService.CreateReportRequest(ctx, bearer, reportRequest)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, request)
}

func (v verificationRequestHandler) GetVerificationRequests(c echo.Context) error {
	span := tracer.StartSpanFromRequest("RequestHandlerGetVerificationRequests", v.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get verification requests at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	retVal, err := v.VerificationRequestService.GetVerificationRequests(ctx)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, retVal)
}

func (v verificationRequestHandler) GetCampaignRequests(c echo.Context) error {
	span := tracer.StartSpanFromRequest("RequestHandlerGetCampaignRequests", v.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get campaign requests at %s\n", c.Path())),
	)

	postId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	retVal, err := v.VerificationRequestService.GetCampaignRequests(ctx, postId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, retVal)
}

func (v verificationRequestHandler) GetReportRequests(c echo.Context) error {
	span := tracer.StartSpanFromRequest("RequestHandlerGetReportRequests", v.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get report requests at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	retVal, err := v.VerificationRequestService.GetReportRequests(ctx)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, retVal)
}

func (v verificationRequestHandler) DeleteCampaignRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("RequestHandlerDeleteCampaignRequest", v.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling delete campaign requests at %s\n", c.Path())),
	)

	postId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	err := v.VerificationRequestService.DeleteCampaignRequest(ctx, postId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (v verificationRequestHandler) DeleteReportRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("RequestHandlerDeleteReportRequest", v.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling delete report requests at %s\n", c.Path())),
	)

	postId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	err := v.VerificationRequestService.DeleteReportRequest(ctx, postId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (v verificationRequestHandler) ApproveVerificationRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("RequestHandlerApproveVerificationRequest", v.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling approve verification request at %s\n", c.Path())),
	)

	postId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	err := v.VerificationRequestService.ApproveVerificationRequest(ctx, postId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (v verificationRequestHandler) RejectVerificationRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("RequestHandlerRejectVerificationRequest", v.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling reject verification request at %s\n", c.Path())),
	)

	postId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	err := v.VerificationRequestService.RejectVerificationRequest(ctx, postId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (v verificationRequestHandler) HasUserPendingRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("RequestHandlerHasUserPendingRequest", v.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling has user pending request at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	bearer := c.Request().Header.Get("Authorization")

	result, err := v.VerificationRequestService.HasUserPendingRequest(ctx, bearer)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
