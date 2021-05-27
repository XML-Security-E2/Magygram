package pgsql

import (
	"context"
	"gorm.io/gorm"
	"magyAgent/domain/model"
	"magyAgent/domain/repository"
)

type loginEventRepository struct {
	Conn *gorm.DB
}

func NewLoginEventRepository(Conn *gorm.DB) repository.LoginEventRepository {
	return &loginEventRepository{Conn}
}

func (r *loginEventRepository) Create(ctx context.Context, event *model.LoginEvent) (*model.LoginEvent, error) {
	err := r.Conn.Create(event).Error
	return event, err
}

func (r *loginEventRepository) GetLastByUserEmail(ctx context.Context, email string) (*model.LoginEvent, error) {
	var events []model.LoginEvent
	err := r.Conn.Where("user_email = ?", email).Order("timestamp desc").Find(&events).Error

	if len(events) == 0 {
		return nil, err
	}
	return &events[0], err
}