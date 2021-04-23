package pgsql

import (
	"context"
	"gorm.io/gorm"
	"magyAgent/domain/model"
	"magyAgent/domain/repository"
)

type accountResetPasswordRepository struct {
	Conn *gorm.DB
}

func NewAccountResetPasswordRepository(Conn *gorm.DB) repository.AccountResetPasswordRepository {
	return &accountResetPasswordRepository{Conn}
}

func (r *accountResetPasswordRepository) Create(ctx context.Context, a *model.AccountResetPassword) (*model.AccountResetPassword, error) {
	err := r.Conn.Create(a).Error
	return a, err
}

func (r *accountResetPasswordRepository) GetById(ctx context.Context, id string) (*model.AccountResetPassword, error) {
	a := &model.AccountResetPassword{Id: id}
	err := r.Conn.First(a).Error
	return a, err
}

