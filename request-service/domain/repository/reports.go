package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"request-service/domain/model"
)

type ReportRequestsRepository interface {
	CreateReport(ctx context.Context, user *model.ReportRequest) (*mongo.InsertOneResult, error)
}
