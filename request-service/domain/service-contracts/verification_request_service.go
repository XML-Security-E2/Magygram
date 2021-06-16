package service_contracts

import (
	"context"
	"mime/multipart"
	"request-service/domain/model"
)

type VerificationRequestService interface {
	CreateVerificationRequest(ctx context.Context, verificationRequsetDTO model.VerificationRequestDTO, bearer string, documentImage []*multipart.FileHeader)  (string, error)
	CreateReportRequest(ctx context.Context, report *model.ReportRequestDTO)  (string, error)
}