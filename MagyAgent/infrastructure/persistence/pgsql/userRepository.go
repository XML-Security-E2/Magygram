package pgsql

import (
	"context"
	"gorm.io/gorm"
	"magyAgent/domain/model"
	"magyAgent/domain/repository"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) repository.UserRepository {
	return &userRepository{Conn}
}

func (r *userRepository) Create(ctx context.Context, u *model.User) (*model.User, error) {
	err := r.Conn.Create(u).Error
	return u, err
}
