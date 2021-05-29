package model

import (
	"github.com/beevik/guid"
	"time"
)

var (
	SuccessfulLogin = "successful"
	UnsuccessfulLogin = "unsuccessful"
	ActivatedAccount = "reactivated"
)

type LoginEvent struct {
	Id string `bson:"_id,omitempty"`
	UserEmail string `bson:"userEmail"`
	Outcome string `bson:"outcome"`
	Timestamp time.Time `bson:"timestamp"`
	RepetitionNumber int `bson:"repetitionNumber"`
}

func NewLoginEvent(userEmail string, outcome string, repetitionNumber int) *LoginEvent {
	repNum := 0
	if outcome == UnsuccessfulLogin {
		repNum = repetitionNumber
	}
	return &LoginEvent{ Id : guid.New().String(), UserEmail: userEmail, Outcome: outcome, RepetitionNumber: repNum, Timestamp: time.Now()}
}