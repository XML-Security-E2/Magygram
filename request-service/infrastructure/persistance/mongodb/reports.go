package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"request-service/domain/model"
	"request-service/domain/repository"
)

type reportRequestsRepository struct {
	Col *mongo.Collection
}

func (v *reportRequestsRepository) GetAllReports(ctx context.Context) ([]*model.ReportRequest, error) {
	cursor, err := v.Col.Find(context.TODO(), bson.M{"deleted": false})
	var results []*model.ReportRequest

	if err != nil {
		defer cursor.Close(ctx)
	} else {
		for cursor.Next(ctx) {
			var result model.ReportRequest

			err := cursor.Decode(&result)
			results = append(results, &result)

			if err != nil {
				os.Exit(1)
			}
		}
	}
	return results, nil
}
func (v reportRequestsRepository) DeleteReportRequest(ctx context.Context, request *model.ReportRequest) (*mongo.UpdateResult, error) {
	return v.Col.UpdateOne(ctx, bson.M{"_id":  request.Id},bson.D{{"$set", bson.D{
		{"deleted" , true}}}})
}

func (v reportRequestsRepository) CreateReport(ctx context.Context, request *model.ReportRequest) (*mongo.InsertOneResult, error) {
	return v.Col.InsertOne(ctx, request)
}

func (v *reportRequestsRepository) GetReportRequestById(ctx context.Context, requestId string) (*model.ReportRequest, error) {
	var post = model.ReportRequest{}
	err := v.Col.FindOne(ctx, bson.M{"_id": requestId}).Decode(&post)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &post, nil
}

func NewReportRequestsRepository(Col *mongo.Collection) repository.ReportRequestsRepository {
	return &reportRequestsRepository{Col}
}
