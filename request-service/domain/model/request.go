package model

import (
	"github.com/beevik/guid"
)

type VerificationRequest struct {
	Id string `bson:"_id,omitempty"`
	UserInfo UserInfo `bson:"user_info"`
	Document string `bson:"document"`
	Status RequestStatus `bson:"request_status"`
}

type UserInfo struct {
	Id string
	Username string
	ImageURL string
}

type RequestStatus string

const(
	PENDING = iota
	APPROVED
	REJECTED
)

type VerificationRequestDTO struct {
	Document string
}

func NewVerificationRequest(verificationRequest *VerificationRequestDTO, requestOwner UserInfo) (*VerificationRequest, error) {
	return &VerificationRequest{Id: guid.New().String(),
		UserInfo:   requestOwner,
		Document:    verificationRequest.Document,
		Status: "PENDING",
	}, nil
}