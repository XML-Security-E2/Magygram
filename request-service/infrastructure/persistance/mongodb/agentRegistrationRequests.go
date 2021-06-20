package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"request-service/domain/model"
	"request-service/domain/repository"
)

type agentRegistrationRequest struct {
	Col *mongo.Collection
}

func NewAgentRegistrationRequestsRepository(Col *mongo.Collection) repository.AgentRegistrationRequests {
	return &agentRegistrationRequest{Col}
}

func (a agentRegistrationRequest) Create(ctx context.Context, request *model.AgentRegistrationRequest) (*mongo.InsertOneResult, error) {
	return a.Col.InsertOne(ctx, request)
}

func (a agentRegistrationRequest) GetAllPendingRequests(ctx context.Context) ([]*model.AgentRegistrationRequest, error) {
	cursor, err := a.Col.Find(ctx, bson.M{"request_status": "PENDING"})
	var results []*model.AgentRegistrationRequest

	if err != nil {
		defer cursor.Close(ctx)
		return nil, err
	} else {
		for cursor.Next(ctx) {
			var result model.AgentRegistrationRequest

			_ = cursor.Decode(&result)
			results = append(results, &result)

		}
	}
	return results, nil
}

func (a agentRegistrationRequest) Update(ctx context.Context, request *model.AgentRegistrationRequest) (*mongo.UpdateResult, error) {
	return a.Col.UpdateOne(ctx, bson.M{"_id":  request.Id},bson.D{{"$set", bson.D{
		{"username" , request.Username},
		{"name" , request.Name},
		{"email" , request.Email},
		{"surname" , request.Surname},
		{"website" , request.Website},
		{"request_status" , request.Status}}}})
}

func (a agentRegistrationRequest) GetByUsernamePendingRequest(ctx context.Context, username string) (*model.AgentRegistrationRequest, error) {
	var agentRegistrationRequest = model.AgentRegistrationRequest{}
	err := a.Col.FindOne(ctx, bson.M{"username": username,"request_status": "PENDING"}).Decode(&agentRegistrationRequest)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &agentRegistrationRequest, nil
}

func (a agentRegistrationRequest) GetByEmailPendingRequest(ctx context.Context, email string) (*model.AgentRegistrationRequest, error) {
	var agentRegistrationRequest = model.AgentRegistrationRequest{}
	err := a.Col.FindOne(ctx, bson.M{"email": email,"request_status": "PENDING"}).Decode(&agentRegistrationRequest)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &agentRegistrationRequest, nil
}

func (a agentRegistrationRequest) GetById(ctx context.Context, requestId string) (*model.AgentRegistrationRequest, error) {
	var agentRegistrationRequest = model.AgentRegistrationRequest{}
	err := a.Col.FindOne(ctx, bson.M{"_id": requestId}).Decode(&agentRegistrationRequest)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &agentRegistrationRequest, nil
}