package service

import (
	"context"
	"github.com/beevik/guid"
	"magyAgent/domain/model"
	"magyAgent/domain/repository"
	"magyAgent/domain/service-contracts"
)

type authService struct {
	repository.UserRepository
}

func NewAuthService(r repository.UserRepository) service_contracts.AuthService {
	return &authService{r}
}

func (u *authService) RegisterUser(ctx context.Context, user *model.User) (*model.User, error) {
	user.Id = guid.New().String()
	return u.UserRepository.Create(ctx, user)
}