package service

import (
	"context"
	"request-service/domain/model"
	"request-service/domain/repository"
	service_contracts "request-service/domain/service-contracts"
)

type verificationService struct {
	repository.VerificationRequestsRepository
}

func NewVerificationServiceService(v repository.VerificationRequestsRepository) service_contracts.VerificationRequestService {
	return &verificationService{v}
}

func (v verificationService) CreateVerificationRequest(ctx context.Context, verificationRequestDTO *model.VerificationRequestDTO) (string, error) {
	//TODO
	var userInfo = model.UserInfo{Id: "123",
		Username: "pera",
		ImageURL: "djura",
		}

	verificationRequest, err := model.NewVerificationRequest(verificationRequestDTO,userInfo)

	if err != nil {
		return "", err
	}

	result, err := v.VerificationRequestsRepository.Create(ctx, verificationRequest)
	if err != nil {
		return "", err
	}

	if requestId, ok := result.InsertedID.(string); ok {
		return requestId, nil
	}

	return "",err
}