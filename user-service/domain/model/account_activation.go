package model

import (
	"github.com/beevik/guid"
	"time"
)

type AccountActivation struct {
	Id string `bson:"_id,omitempty"`
	UserId string `bson:"userId"`
	GenerationDate time.Time `bson:"generationDate"`
	ExpirationDate time.Time `bson:"expirationDate"`
	Used bool `bson:"used"`
}

func NewAccountActivation(userId string) *AccountActivation {
	generationDate := time.Now()
	return &AccountActivation{Id: guid.New().String(), UserId: userId, GenerationDate: generationDate, ExpirationDate: generationDate.Add(30 * time.Minute), Used: false}
}
