package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"request-service/domain/model"
	"request-service/domain/repository"
	"request-service/domain/service-contracts"
	"request-service/service/intercomm"
	"strings"
)

type verificationService struct {
	repository.VerificationRequestsRepository
	intercomm.MediaClient
	intercomm.AuthClient
	repository.ReportRequestsRepository
	repository.CampaignRequestsRepository
	intercomm.UserClient
}

func NewVerificationServiceService(vr repository.VerificationRequestsRepository,r repository.ReportRequestsRepository, cr repository.CampaignRequestsRepository, mc intercomm.MediaClient, ac intercomm.AuthClient, uc intercomm.UserClient) service_contracts.VerificationRequestService {
	return &verificationService{vr,mc, ac,r ,cr,uc}
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
func (v verificationService) CreateCampaignRequest(ctx context.Context, bearer string,  campaignRequestDTO *model.CampaignRequestDTO)  (string, error) {

	campaignRequest, err := model.NewCampaignRequest(campaignRequestDTO)
	if err != nil {
		return "", err
	}

	result, err := v.CampaignRequestsRepository.CreateRequest(ctx, campaignRequest)
	if err != nil {
		return "", err
	}
	if requestId, ok := result.InsertedID.(string); ok {
		return requestId, nil
	}

	return "",err
}


func (v verificationService) CreateReportRequest(ctx context.Context, bearer string, reportRequestDTO *model.ReportRequestDTO) (string, error) {
	loggedId, err := v.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return "", err
	}

	reportRequest, err := model.NewReportRequest(reportRequestDTO, loggedId)
	if err != nil {
		return "", err
	}

	request, err := v.ReportRequestsRepository.GetReportByContentIdAndUserWhoReported(ctx, loggedId, reportRequest.ContentId)
	if err != nil {
		return "", err
	}

	if request != nil {
		return "", errors.New("content already reported")
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
func (v verificationService) GetCampaignRequests(ctx context.Context, requestId string) ([]*model.CampaignRequestResponseDTO, error) {
	var campaignRequest []*model.CampaignRequest

	campaignRequest, err := v.CampaignRequestsRepository.GetAllRequests(ctx, requestId)
	if err != nil {
		return []*model.CampaignRequestResponseDTO{}, err
	}

	retVal := mapCampaignRequestToCampaignRequestDTO(campaignRequest)

	return retVal, nil
}

func mapCampaignRequestToCampaignRequestDTO(requests []*model.CampaignRequest) []*model.CampaignRequestResponseDTO {
	var retVal []*model.CampaignRequestResponseDTO

	for _, request := range requests{
		var result = model.CampaignRequestResponseDTO{ Id: request.Id,
			ContentId:      request.ContentId,
			ContentType:    request.ContentType,
			Influencer:   request.Influencer,
			Status: request.Status,
			Price: request.Price,
		}
		retVal = append(retVal, &result)
	}

	return retVal
}


func (v verificationService) GetReportRequests(ctx context.Context) ([]*model.ReportRequestResponseDTO, error) {
	var reportRequest []*model.ReportRequest

	reportRequest, err := v.ReportRequestsRepository.GetAllReports(ctx)
	if err != nil {
		return []*model.ReportRequestResponseDTO{}, err
	}

	retVal := mapReportRequestToReportRequestDTO(reportRequest)

	return retVal, nil
}

func mapReportRequestToReportRequestDTO(requests []*model.ReportRequest) []*model.ReportRequestResponseDTO {
	var retVal []*model.ReportRequestResponseDTO

	for _, request := range requests{
		var result = model.ReportRequestResponseDTO{ Id: request.Id,
			ContentId:      request.ContentId,
			ContentType:   request.ContentType,
			UserWhoReportedId: request.UserWhoReportedId,
		}
		retVal = append(retVal, &result)
	}

	return retVal
}

func (v verificationService) GetVerificationRequests(ctx context.Context) ([]*model.VerificationRequestResponseDTO, error) {
	var verificationRequest []*model.VerificationRequest

	verificationRequest, err := v.VerificationRequestsRepository.GetAllPendingRequests(ctx)
	if err != nil {
		return []*model.VerificationRequestResponseDTO{}, err
	}

	retVal := mapVerificationRequestToVerificationRequestResponseDTO(verificationRequest)

	return retVal, nil
}

func mapVerificationRequestToVerificationRequestResponseDTO(requests []*model.VerificationRequest) []*model.VerificationRequestResponseDTO {
	var retVal []*model.VerificationRequestResponseDTO

	for _, request := range requests{
		var result = model.VerificationRequestResponseDTO{ Id: request.Id,
			Name:      request.Name,
			Surname:   request.Surname,
			UserId :   request.UserId,
			Document : request.Document,
			Category: string(request.Category[0]) + strings.ToLower(string(request.Category[1:])),
		}
		retVal = append(retVal, &result)
	}

	return retVal
}

func (v verificationService) ApproveVerificationRequest(ctx context.Context, requestId string) error {
	request, err := v.VerificationRequestsRepository.GetVerificationRequestById(ctx, requestId)
	if err!=nil {
		return errors.New("Request not found")
	}

	if request.Status=="REJECTED" || request.Status=="APPROVED" {
		return errors.New("The request has already been processed.")
	}

	request.Status="APPROVED"
	//TODO: pozvati userClient da verifikuje u userServicu i eventualno poslati mail useru
	verifyDTO := model.VerifyAccountDTO{
		UserId:   request.UserId,
		Category: string(request.Category),
	}

	log.Println("test")
	err = v.UserClient.VerifyAccount(verifyDTO)
	if err!=nil{
		return err
	}

	v.VerificationRequestsRepository.UpdateVerificationRequest(ctx,request)

	return nil
}

func (v verificationService) RejectVerificationRequest(ctx context.Context, requestId string) error {
	request, err := v.VerificationRequestsRepository.GetVerificationRequestById(ctx, requestId)
	if err!=nil {
		return errors.New("Request not found")
	}

	if request.Status=="REJECTED" || request.Status=="APPROVED" {
		return errors.New("The request has already been processed.")
	}

	request.Status="REJECTED"

	//TODO: pozvati userClient da verifikuje u userServicu i eventualno poslati mail useru
	v.VerificationRequestsRepository.UpdateVerificationRequest(ctx,request)

	return nil
}
func (v verificationService) DeleteCampaignRequest(ctx context.Context, requestId string) error {
	request, err := v.CampaignRequestsRepository.GetRequestById(ctx, requestId)
	fmt.Println(request)
	if err!=nil {
		return errors.New("Request not found")
	}

	request.IsDeleted=true
	fmt.Println(request)
	v.CampaignRequestsRepository.CreateRequest(ctx,request)

	return nil
}

func (v verificationService) DeleteReportRequest(ctx context.Context, requestId string) error {
	request, err := v.ReportRequestsRepository.GetReportRequestById(ctx, requestId)

	if err!=nil {
		return errors.New("Request not found")
	}

	request.IsDeleted=true

	v.ReportRequestsRepository.DeleteReportRequest(ctx,request)

	return nil
}


func (v verificationService) HasUserPendingRequest(ctx context.Context, bearer string) (bool, error) {
	loggedId, err := v.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return false, err
	}

	_, err = v.VerificationRequestsRepository.GetVerificationPendingRequestByUserId(ctx,loggedId)
	if err!=nil{
		return false, nil
	}

	return true,nil
}