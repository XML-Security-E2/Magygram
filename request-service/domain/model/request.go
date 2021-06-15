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

type ReportRequest struct {
	Id string `bson:"_id,omitempty"`
	ContentId string `bson:"content_id"`
	ContentType ContentType `bson:"content_type"`
}

type ContentType string

const(
	USER = iota
	POST
	STORY
)

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

type ReportRequestDTO struct {
	ContentId string `json:"contentId"`
	ContentType ContentType `json:"contentType"`
}

func NewVerificationRequest(verificationRequest *VerificationRequestDTO, requestOwner UserInfo) (*VerificationRequest, error) {
	return &VerificationRequest{Id: guid.New().String(),
		UserInfo:   requestOwner,
		Document:    verificationRequest.Document,
		Status: "PENDING",
	}, nil
}

func NewReportRequest(reportRequest *ReportRequestDTO) (*ReportRequest, error) {
	return &ReportRequest{Id: guid.New().String(),
		ContentId:   reportRequest.ContentId,
		ContentType:    reportRequest.ContentType,
	}, nil
}