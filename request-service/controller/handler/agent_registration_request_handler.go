package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"request-service/domain/model"
	"request-service/tracer"

	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"
)

type AgentRegistrationRequestHandler interface {
	CreateAgentRegistrationRequest(c echo.Context) error
	GetAgentRegistrationRequests(c echo.Context) error
	ApproveAgentRegistrationRequest(c echo.Context) error
	RejectAgentRegistrationRequest(c echo.Context) error
}

type agentRegistrationRequestHandler struct {
	AgentRegistrationRequestService service_contracts.AgentRegistrationRequestService
	tracer                          opentracing.Tracer
	closer                          io.Closer
}

func NewAgentRegistrationRequestHandler(a service_contracts.AgentRegistrationRequestService) AgentRegistrationRequestHandler {
	tracer, closer := tracer.Init("request-service")
	opentracing.SetGlobalTracer(tracer)
	return &agentRegistrationRequestHandler{a, tracer, closer}
}

func (a agentRegistrationRequestHandler) CreateAgentRegistrationRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("RequestHandlerCreateAgentRegistrationRequest", a.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling create agent registration request at %s\n", c.Path())),
	)

	agentRegistrationRequest := &model.AgentRegistrationRequestDTO{}
	if err := c.Bind(agentRegistrationRequest); err != nil {
		tracer.LogError(span, err)
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	request, err := a.AgentRegistrationRequestService.CreateVerificationRequest(ctx, *agentRegistrationRequest)

	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, request)
}

func (a agentRegistrationRequestHandler) GetAgentRegistrationRequests(c echo.Context) error {
	span := tracer.StartSpanFromRequest("RequestHandlerGetAgentRegistrationRequests", a.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling get agent registration request at %s\n", c.Path())),
	)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	retVal, err := a.AgentRegistrationRequestService.GetAgentRegistrationRequests(ctx)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, retVal)
}

func (a agentRegistrationRequestHandler) ApproveAgentRegistrationRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("RequestHandlerApproveAgentRegistrationRequest", a.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling approve agent registration request at %s\n", c.Path())),
	)

	requestId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	err := a.AgentRegistrationRequestService.ApproveAgentRegistrationRequest(ctx, requestId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (a agentRegistrationRequestHandler) RejectAgentRegistrationRequest(c echo.Context) error {
	span := tracer.StartSpanFromRequest("RequestHandlerRejectAgentRegistrationRequest", a.tracer, c.Request())
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling reject agent registration request at %s\n", c.Path())),
	)

	requestId := c.Param("requestId")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = tracer.ContextWithSpan(ctx, span)

	err := a.AgentRegistrationRequestService.RejectAgentRegistrationRequest(ctx, requestId)
	if err != nil {
		tracer.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}
