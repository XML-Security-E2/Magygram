package service_contracts

import (
	"context"
	"mime/multipart"
	"request-service/domain/model"
)

type VerificationRequestService interface {
	CreateVerificationRequest(ctx context.Context, verificationRequsetDTO model.VerificationRequestDTO, bearer string, documentImage []*multipart.FileHeader)  (string, error)
	CreateReportRequest(ctx context.Context, report *model.ReportRequestDTO)  (string, error)
	GetVerificationRequests(ctx context.Context) ([]*model.VerificationRequestResponseDTO, error)
	GetReportRequests(ctx context.Context) ([]*model.ReportRequestResponseDTO, error)
	ApproveVerificationRequest(ctx context.Context, requestId string) error
	DeleteReportRequest(ctx context.Context, requestId string) error
	RejectVerificationRequest(ctx context.Context, requestId string) error
	HasUserPendingRequest(ctx context.Context, bearer string) (bool, error)
}