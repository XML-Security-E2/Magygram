package service_contracts

import (
	"context"
	"request-service/domain/model"
)

type AgentRegistrationRequestService interface {
	CreateVerificationRequest(ctx context.Context, request model.AgentRegistrationRequestDTO) (string, error)
	GetAgentRegistrationRequests(ctx context.Context) ([]*model.AgentRegistrationRequestResponseDTO, error)
	ApproveAgentRegistrationRequest(ctx context.Context, requestId string) error
	RejectAgentRegistrationRequest(ctx context.Context, requestId string) error

}