package model

import (
	"github.com/beevik/guid"
	"time"
)

type AccountResetPassword struct {
	Id string `gorm:"primaryKey"`
	UserId string `gorm:"not null"`
	GenerationDate time.Time `gorm:"not null"`
	ExpirationDate time.Time `gorm:"not null"`
	Used bool `gorm:"not null"`
}

func NewAccountResetPassword(userId string) *AccountResetPassword {
	generationDate := time.Now()
	return &AccountResetPassword{Id: guid.New().String(), UserId: userId, GenerationDate: generationDate, ExpirationDate: generationDate.Add(30 * time.Minute), Used: false}
}