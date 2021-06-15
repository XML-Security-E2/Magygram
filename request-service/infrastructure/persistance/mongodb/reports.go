package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"request-service/domain/model"
	"request-service/domain/repository"
)

type reportRequestsRepository struct {
	Col *mongo.Collection
}

func (v reportRequestsRepository) CreateReport(ctx context.Context, request *model.ReportRequest) (*mongo.InsertOneResult, error) {
	return v.Col.InsertOne(ctx, request)
}

func NewReportRequestsRepository(Col *mongo.Collection) repository.ReportRequestsRepository {
	return &reportRequestsRepository{Col}
}
