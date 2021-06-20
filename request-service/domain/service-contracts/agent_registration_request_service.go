package service_contracts

import (
	"context"
	"request-service/domain/model"
)

type AgentRegistrationRequestService interface {
	CreateVerificationRequest(ctx context.Context, request model.AgentRegistrationRequestDTO) (string, error)
}