package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"request-service/domain/model"
)

type ReportRequestsRepository interface {
	CreateReport(ctx context.Context, user *model.ReportRequest) (*mongo.InsertOneResult, error)
	GetAllReports(ctx context.Context) ([]*model.ReportRequest, error)
	DeleteReportRequest(ctx context.Context, request *model.ReportRequest) (*mongo.UpdateResult, error)
	GetReportRequestById(ctx context.Context, requestId string) (*model.ReportRequest, error)
}


