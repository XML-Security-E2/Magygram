package service

import (
	"context"
	"errors"
	"request-service/domain/model"
	"request-service/domain/repository"
	"request-service/domain/service-contracts"
	"request-service/service/intercomm"
)

type agentRegistrationRequestService struct {
	repository.AgentRegistrationRequests
	intercomm.AuthClient
	intercomm.UserClient
}

func NewAgentRegistrationRequestService(vr repository.AgentRegistrationRequests,  ac intercomm.AuthClient, uc intercomm.UserClient) service_contracts.AgentRegistrationRequestService {
	return &agentRegistrationRequestService{vr,ac,uc}
}

func (a agentRegistrationRequestService) CreateVerificationRequest(ctx context.Context, request model.AgentRegistrationRequestDTO) (string, error) {
	user, _ := a.AgentRegistrationRequests.GetByUsernamePendingRequest(ctx, request.Username)
	if user != nil {
		return "", errors.New("Request for user with this username exist")
	}

	user, _ = a.AgentRegistrationRequests.GetByEmailPendingRequest(ctx, request.Email)
	if user != nil {
		return "", errors.New("Request for user with this email exist")
	}

	agentRegistrationRequest, err:= model.NewAgentRegistrationRequest(&request,"PENDING")
	if err != nil {
		return "", err
	}

	result, err := a.AgentRegistrationRequests.Create(ctx, agentRegistrationRequest)
	if err != nil {
		return "", err
	}

	if requestId, ok := result.InsertedID.(string); ok {
		return requestId, nil
	}

	return "",err
}

