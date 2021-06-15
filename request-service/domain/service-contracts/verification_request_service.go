package service_contracts

import (
	"context"
	"request-service/domain/model"
)

type VerificationRequestService interface {
	CreateVerificationRequest(ctx context.Context, user *model.VerificationRequestDTO)  (string, error)
}