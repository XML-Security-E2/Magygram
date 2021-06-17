package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
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

func (v verificatioRequestsRepository) GetAllPendingRequests(ctx context.Context) ([]*model.VerificationRequest, error) {
	cursor, err := v.Col.Find(ctx, bson.M{"request_status": "PENDING"})
	var results []*model.VerificationRequest

	if err != nil {
		defer cursor.Close(ctx)
		return nil, err
	} else {
		for cursor.Next(ctx) {
			var result model.VerificationRequest

			_ = cursor.Decode(&result)
			results = append(results, &result)

		}
	}
	return results, nil
}

func (v verificatioRequestsRepository) GetVerificationRequestById(ctx context.Context, requestId string) (*model.VerificationRequest, error) {
	var post = model.VerificationRequest{}
	err := v.Col.FindOne(ctx, bson.M{"_id": requestId}).Decode(&post)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &post, nil
}

func (v verificatioRequestsRepository) UpdateVerificationRequest(ctx context.Context, request *model.VerificationRequest) (*mongo.UpdateResult, error) {
	return v.Col.UpdateOne(ctx, bson.M{"_id":  request.Id},bson.D{{"$set", bson.D{
		{"user_id" , request.UserId},
		{"user_name" , request.Name},
		{"user_surname" , request.Surname},
		{"document" , request.Document},
		{"request_status" , request.Status},
		{"category" , request.Category}}}})
}

func (v verificatioRequestsRepository) GetVerificationPendingRequestByUserId(ctx context.Context, userId string) (*model.VerificationRequest, error) {
	var verificationRequest = model.VerificationRequest{}
	err := v.Col.FindOne(ctx, bson.M{"user_id": userId,"request_status": "PENDING"}).Decode(&verificationRequest)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &verificationRequest, nil
}
