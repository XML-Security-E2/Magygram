package model

import (
	"github.com/beevik/guid"
	"time"
)

type AccountResetPassword struct {
	Id string `gorm:"primaryKey"`
	UserId string
	GenerationDate time.Time
	ExpirationDate time.Time
	Used bool
}

func NewAccountResetPassword(userId string) *AccountResetPassword {
	generationDate := time.Now()
	return &AccountResetPassword{Id: guid.New().String(), UserId: userId, GenerationDate: generationDate, ExpirationDate: generationDate.Add(30 * time.Minute), Used: false}
}