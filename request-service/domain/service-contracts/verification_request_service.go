package service_contracts

import (
	"context"
	"request-service/domain/model"
)

type VerificationRequestService interface {
	CreateVerificationRequest(ctx context.Context, user *model.VerificationRequestDTO)  (string, error)
	CreateReportRequest(ctx context.Context, report *model.ReportRequestDTO)  (string, error)
}