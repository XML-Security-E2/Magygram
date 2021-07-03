package service_contracts

import (
	"auth-service/domain/model"
	"context"
)

type AuthService interface {
	AuthenticateUser(ctx context.Context, loginRequest *model.LoginRequest) (*model.User, error)
	HasUserPermission(ctx context.Context, desiredPermission string, userId string) (bool, error)
	HandleLoginEventAndAccountActivation(ctx context.Context, userEmail string, successful bool, eventType string)
	AuthenticateTwoFactoryUser(ctx context.Context, loginRequest *model.LoginTwoFactoryRequest) (*model.User, error)
	UpdateAgentCampaignJWTToken(ctx context.Context, bearer string, token string) error
	DeleteCampaignJWTToken(ctx context.Context, bearer string) error
}