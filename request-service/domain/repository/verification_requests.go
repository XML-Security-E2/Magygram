package repository

import (
"context"
"go.mongodb.org/mongo-driver/mongo"
"request-service/domain/model"
)

type VerificationRequestsRepository interface {
	Create(ctx context.Context, user *model.VerificationRequest) (*mongo.InsertOneResult, error)
	GetAllPendingRequests(ctx context.Context) ([]*model.VerificationRequest, error)
	GetVerificationRequestById(ctx context.Context, requestId string) (*model.VerificationRequest, error)
	UpdateVerificationRequest(ctx context.Context, request *model.VerificationRequest) (*mongo.UpdateResult, error)
	GetVerificationPendingRequestByUserId(ctx context.Context, userId string) (*model.VerificationRequest, error)
}
