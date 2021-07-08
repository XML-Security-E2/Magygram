package service

import (
	"context"
	"errors"
	"request-service/domain/model"
	"request-service/domain/repository"
	"request-service/domain/service-contracts"
	"request-service/service/intercomm"
	"request-service/tracer"
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
	span := tracer.StartSpanFromContext(ctx, "RequestServiceCreateVerificationRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

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

func (a agentRegistrationRequestService) GetAgentRegistrationRequests(ctx context.Context) ([]*model.AgentRegistrationRequestResponseDTO, error) {
	span := tracer.StartSpanFromContext(ctx, "RequestServiceGetAgentRegistrationRequests")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	var agentRegistrationRequest []*model.AgentRegistrationRequest

	agentRegistrationRequest, err := a.AgentRegistrationRequests.GetAllPendingRequests(ctx)
	if err != nil {
		return []*model.AgentRegistrationRequestResponseDTO{}, err
	}

	retVal := mapAgentRegistrationRequestToAgentRegistrationRequestResponseDTO(agentRegistrationRequest)

	return retVal, nil
}

func mapAgentRegistrationRequestToAgentRegistrationRequestResponseDTO(requests []*model.AgentRegistrationRequest) []*model.AgentRegistrationRequestResponseDTO {
	var retVal []*model.AgentRegistrationRequestResponseDTO

	for _, request := range requests{
		var result = model.AgentRegistrationRequestResponseDTO{ Id: request.Id,
			Name:      request.Name,
			Surname:   request.Surname,
			Username :   request.Username,
			Email : request.Email,
			WebSite: request.Website,
		}
		retVal = append(retVal, &result)
	}

	return retVal
}

func (a agentRegistrationRequestService) ApproveAgentRegistrationRequest(ctx context.Context, requestId string) error {
	span := tracer.StartSpanFromContext(ctx, "RequestServiceApproveAgentRegistrationRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	result, err := a.AgentRegistrationRequests.GetById(ctx, requestId)
	if err!=nil {
		return errors.New("Request not found")
	}

	if result.Status=="REJECTED" || result.Status=="APPROVED" {
		return errors.New("The request has already been processed.")
	}

	result.Status="APPROVED"

	//TODO: pozvati userClient da verifikuje u userServicu i eventualno poslati mail useru
	agentRegistrationDTO := model.AgentRegistrationDTO{
		Username:   result.Username,
		Password: result.Password,
		Email: result.Email,
		Name: result.Name,
		Surname : result.Surname,
		Website: result.Website,
	}

	err = a.UserClient.RegisterAgent(ctx, agentRegistrationDTO)
	if err!=nil{
		return err
	}

	a.AgentRegistrationRequests.Update(ctx,result)

	return nil
}

func (a agentRegistrationRequestService) RejectAgentRegistrationRequest(ctx context.Context, requestId string) error {
	span := tracer.StartSpanFromContext(ctx, "RequestServiceRejectAgentRegistrationRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	result, err := a.AgentRegistrationRequests.GetById(ctx, requestId)
	if err!=nil {
		return errors.New("Request not found")
	}

	if result.Status=="REJECTED" || result.Status=="APPROVED" {
		return errors.New("The request has already been processed.")
	}

	result.Status="REJECTED"

	//TODO: pozvati userClient da verifikuje u userServicu i eventualno poslati mail useru
	a.AgentRegistrationRequests.Update(ctx,result)

	return nil
}