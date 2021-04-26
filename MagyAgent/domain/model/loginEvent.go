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
	Id string `gorm:"primaryKey"`
	UserEmail string
	Outcome string
	Timestamp time.Time
	RepetitionNumber int
}

func NewLoginEvent(userEmail string, outcome string, repetitionNumber int) *LoginEvent {
	repNum := 0
	if outcome == UnsuccessfulLogin {
		repNum = repetitionNumber
	}
	return &LoginEvent{ Id : guid.New().String(), UserEmail: userEmail, Outcome: outcome, RepetitionNumber: repNum, Timestamp: time.Now()}
}
