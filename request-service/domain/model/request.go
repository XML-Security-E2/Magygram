package model

import (
	"errors"
	"github.com/beevik/guid"
)

type VerificationRequest struct {
	Id string `bson:"_id,omitempty"`
	UserId string `bson:"user_id,omitempty"`
	Name string `bson:"user_name"`
	Surname string `bson:"user_surname"`
	Document string `bson:"document"`
	Status RequestStatus `bson:"request_status"`
	Category Category `bson:"category"`
}

type VerificationRequestResponseDTO struct {
	Id string
	UserId string
	Name string
	Surname string
	Document string
	Category string
}

type ReportRequestResponseDTO struct {
	Id string
	ContentId string
	ContentType ContentType
}

type ReportRequest struct {
	Id string `bson:"_id,omitempty"`
	ContentId string `bson:"content_id"`
	ContentType ContentType `bson:"content_type"`
	IsDeleted bool `bson:"deleted"`
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

type Category string

const(
	INFLUENCER = iota
	SPORTS
	NEWS//MEDIA
	BUSINESS
	BRAND
	ORGANIZATION
	MUSIC
	ACTOR
)

type VerificationRequestDTO struct {
	Name string
	Surname string
	Category string
}


type ReportRequestDTO struct {
	ContentId string `json:"contentId"`
	ContentType ContentType `json:"contentType"`
}

func NewVerificationRequest(verificationRequest *VerificationRequestDTO, requestStatus RequestStatus, category Category, userId string, imageUrl string) (*VerificationRequest, error) {
	err := validateRequestStatusTypeEnums(requestStatus)
	if err != nil {
			return nil, err
		}

		err = validateCategoryTypeEnums(category)
		if err != nil {
			return nil, err
		}

	return &VerificationRequest{Id: guid.New().String(),
		Name:   verificationRequest.Name,
		Surname:    verificationRequest.Surname,
		UserId: userId,
		Document: imageUrl,
		Status: requestStatus,
		Category: category,
	}, nil
}


func validateCategoryTypeEnums(category Category) error {
	switch category {
	case "INFLUENCER", "SPORTS", "NEWS/MEDIA", "BUSINESS", "BRAND", "ORGANIZATION", "MUSIC", "ACTOR":
		return nil
	}
	return errors.New("Invalid post type")
}

func validateRequestStatusTypeEnums(status RequestStatus) error{
	switch status {
	case "PENDING", "APPROVED", "REJECTED":
		return nil
	}
	return errors.New("Invalid post type")
}

type Media struct {
	Url       string `json:"url"`
	MediaType string `json:"mediaType"`
}

func NewReportRequest(reportRequest *ReportRequestDTO) (*ReportRequest, error) {
	return &ReportRequest{Id: guid.New().String(),
		ContentId:   reportRequest.ContentId,
		ContentType:    reportRequest.ContentType,
		IsDeleted: false,
	}, nil
}

type VerifyAccountDTO struct{
	UserId string
	Category string
}

