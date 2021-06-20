package handler

import (
	"context"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"request-service/domain/model"
	"request-service/domain/service-contracts"
)

type AgentRegistrationRequestHandler interface {
	CreateAgentRegistrationRequest(c echo.Context) error
}

type agentRegistrationRequestHandler struct {
	AgentRegistrationRequestService service_contracts.AgentRegistrationRequestService
}


func NewAgentRegistrationRequestHandler(a service_contracts.AgentRegistrationRequestService) AgentRegistrationRequestHandler {
	return &agentRegistrationRequestHandler{a}
}

func (a agentRegistrationRequestHandler) CreateAgentRegistrationRequest(c echo.Context) error {
	agentRegistrationRequest := &model.AgentRegistrationRequestDTO{}
	if err := c.Bind(agentRegistrationRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	log.Println("test")

	request, err := a.AgentRegistrationRequestService.CreateVerificationRequest(ctx, *agentRegistrationRequest)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, request)}
