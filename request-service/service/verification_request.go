package service

import (
	"context"
	"errors"
	"mime/multipart"
	"request-service/domain/model"
	"request-service/domain/repository"
	"request-service/domain/service-contracts"
	"request-service/service/intercomm"
)

type verificationService struct {
	repository.VerificationRequestsRepository
	intercomm.MediaClient
	intercomm.AuthClient
	repository.ReportRequestsRepository

}

func NewVerificationServiceService(v repository.VerificationRequestsRepository,r repository.ReportRequestsRepository, mc intercomm.MediaClient, ac intercomm.AuthClient) service_contracts.VerificationRequestService {
	return &verificationService{v,mc, ac,r}
}

func (v verificationService) CreateVerificationRequest(ctx context.Context, verificationRequsetDTO model.VerificationRequestDTO, bearer string, documentImage []*multipart.FileHeader) (string, error) {
	loggedId, err := v.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return "", err
	}

	media, err := v.MediaClient.SaveMedia(documentImage)
	if err != nil { return "", err}

	if len(media) == 0 {
		return "", errors.New("error while saving image")
	}

	verificationRequest, err:= model.NewVerificationRequest(&verificationRequsetDTO,"PENDING", model.Category(verificationRequsetDTO.Category), loggedId, media[0].Url)
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