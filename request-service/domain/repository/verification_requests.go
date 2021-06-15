package repository

import (
"context"
"go.mongodb.org/mongo-driver/mongo"
"request-service/domain/model"
)

type VerificationRequestsRepository interface {
	Create(ctx context.Context, user *model.VerificationRequest) (*mongo.InsertOneResult, error)
}
