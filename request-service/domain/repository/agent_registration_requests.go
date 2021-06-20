package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"request-service/domain/model"
)

type AgentRegistrationRequests interface {
	Create(ctx context.Context, request *model.AgentRegistrationRequest) (*mongo.InsertOneResult, error)
	GetAllPendingRequests(ctx context.Context) ([]*model.AgentRegistrationRequest, error)
	Update(ctx context.Context, request *model.AgentRegistrationRequest) (*mongo.UpdateResult, error)
	GetByUsernamePendingRequest(ctx context.Context, username string) (*model.AgentRegistrationRequest, error)
	GetByEmailPendingRequest(ctx context.Context, username string) (*model.AgentRegistrationRequest, error)
	GetById(ctx context.Context, requestId string) (*model.AgentRegistrationRequest, error)
}
