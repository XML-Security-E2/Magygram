package model

import (
	"github.com/beevik/guid"
	"time"
)

type AccountActivation struct {
	Id string `gorm:"primaryKey"`
	UserId string
	GenerationDate time.Time
	ExpirationDate time.Time
}

func NewAccountActivation(userId string) *AccountActivation {
	generationDate := time.Now()
	return &AccountActivation{Id: guid.New().String(), UserId: userId, GenerationDate: generationDate, ExpirationDate: generationDate.Add(30 * time.Minute)}
}