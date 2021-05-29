package model

import (
	"github.com/beevik/guid"
	"time"
)

type ResetPassword struct {
	Id string `bson:"_id,omitempty"`
	UserId string `bson:"userId"`
	GenerationDate time.Time `bson:"generationDate"`
	ExpirationDate time.Time `bson:"expirationDate"`
	Used bool `bson:"used"`
}

func NewResetPassword(userId string) *ResetPassword {
	generationDate := time.Now()
	return &ResetPassword{Id: guid.New().String(), UserId: userId, GenerationDate: generationDate, ExpirationDate: generationDate.Add(30 * time.Minute), Used: false}
}