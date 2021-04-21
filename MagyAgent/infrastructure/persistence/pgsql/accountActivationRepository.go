package pgsql

import (
	"context"
	"gorm.io/gorm"
	"magyAgent/domain/model"
	"magyAgent/domain/repository"
)

type accountActivationRepository struct {
	Conn *gorm.DB
}

func NewAccountActivationRepository(Conn *gorm.DB) repository.AccountActivationRepository {
	return &accountActivationRepository{Conn}
}

func (r *accountActivationRepository) Create(ctx context.Context, a *model.AccountActivation) (*model.AccountActivation, error) {
	err := r.Conn.Create(a).Error
	return a, err
}

func (r *accountActivationRepository) GetById(ctx context.Context, id string) (*model.AccountActivation, error) {
	a := &model.AccountActivation{Id: id}
	err := r.Conn.First(a).Error
	return a, err
}

