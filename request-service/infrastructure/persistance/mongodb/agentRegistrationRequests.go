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
	panic("implement me")
}

func (a agentRegistrationRequest) UpdateVerificationRequest(ctx context.Context, request *model.AgentRegistrationRequest) (*mongo.UpdateResult, error) {
	panic("implement me")
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

	return &agentRegistrationRequest, nil}