package service

import (
	"context"
	"request-service/domain/model"
	"request-service/domain/repository"
	service_contracts "request-service/domain/service-contracts"
)

type verificationService struct {
	repository.VerificationRequestsRepository
	repository.ReportRequestsRepository
}

func NewVerificationServiceService(v repository.VerificationRequestsRepository,r repository.ReportRequestsRepository) service_contracts.VerificationRequestService {
	return &verificationService{v, r}
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

func (v verificationService) CreateReportRequest(ctx context.Context, reportRequestDTO *model.ReportRequestDTO) (string, error) {

	reportRequest, err := model.NewReportRequest(reportRequestDTO)

	if err != nil {
		return "", err
	}

	result, err := v.ReportRequestsRepository.CreateReport(ctx, reportRequest)
	if err != nil {
		return "", err
	}

	if requestId, ok := result.InsertedID.(string); ok {
		return requestId, nil
	}

	return "",err
}