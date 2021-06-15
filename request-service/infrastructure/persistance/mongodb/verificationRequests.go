package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"request-service/domain/model"
	"request-service/domain/repository"
)

type verificatioRequestsRepository struct {
	Col *mongo.Collection
}


func NewVerificatioRequestsRepository(Col *mongo.Collection) repository.VerificationRequestsRepository {
	return &verificatioRequestsRepository{Col}
}

func (v verificatioRequestsRepository) Create(ctx context.Context, request *model.VerificationRequest) (*mongo.InsertOneResult, error) {
	return v.Col.InsertOne(ctx, request)
}