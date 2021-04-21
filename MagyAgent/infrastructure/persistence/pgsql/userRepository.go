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

func (r *userRepository) Update(ctx context.Context, u *model.User) (*model.User, error) {
	err := r.Conn.Model(u).Updates(u).Error
	return u, err
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	u := &model.User{Id: id}
	err := r.Conn.First(u).Error
	return u, err
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	u := &model.User{Email: email}
	err := r.Conn.Where("email = ?", email).First(&u).Error
	return u, err
}


