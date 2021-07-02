package model

import (
	"errors"
	"github.com/beevik/guid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
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
	UserWhoReportedId string
}

type CampaignRequestResponseDTO struct {
	Id string
	ContentId string
	ContentType ContentType
	Influencer string
	Status RequestStatus
	Price string
}

type ReportRequest struct {
	Id string `bson:"_id,omitempty"`
	ReportReasons []string `json:"report_reasons"`
	UserWhoReportedId string `bson:"user_who_reported_id"`
	ContentId string `bson:"content_id"`
	ContentType ContentType `bson:"content_type"`
	IsDeleted bool `bson:"deleted"`
}

type CampaignRequest struct {
	Id string `bson:"_id,omitempty"`
	Influencer string `bson:"username"`
	ContentId string `bson:"content_id"`
	ContentType ContentType `bson:"content_type"`
	Price string `bson:"price"`
	Status RequestStatus `bson:"request_status"`
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
	ReportReasons []string `json:"reportReasons"`
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

func NewReportRequest(reportRequest *ReportRequestDTO, whoReported string) (*ReportRequest, error) {
	return &ReportRequest{Id: guid.New().String(),
		ContentId:   reportRequest.ContentId,
		ContentType:    reportRequest.ContentType,
		UserWhoReportedId: whoReported,
		ReportReasons: reportRequest.ReportReasons,
		IsDeleted: false,
	}, nil
}


type CampaignRequestDTO struct {
	ContentId string `json:"contentId"`
	Influencer string `json:"username"`
	Price string `json:"price"`
	ContentType ContentType `json:"contentType"`
	Status RequestStatus `json:"status"`

}

func NewCampaignRequest(campaignRequest *CampaignRequestDTO) (*CampaignRequest, error) {
	return &CampaignRequest{Id: guid.New().String(),
		ContentId:   campaignRequest.ContentId,
		ContentType:    campaignRequest.ContentType,
		Influencer:    campaignRequest.Influencer,
		Price: campaignRequest.Price,
		Status: campaignRequest.Status,
		IsDeleted: false,
	}, nil
}

type VerifyAccountDTO struct{
	UserId string
	Category string
}


type AgentRegistrationRequest struct {
	Id	string	`bson:"_id,omitempty"`
	Username	string	`bson:"username" validate:"required,min=1"`
	Name	string 	`bson:"name" validate:"required,min=2"`
	Email	 string	`bson:"email" validate:"required,email"`
	Surname	 string	`bson:"surname" validate:"required,min=2"`
	Website	string	`bson:"website" `
	Status RequestStatus `bson:"request_status"`
	Password string `bson:"password"`
}

type AgentRegistrationRequestDTO struct {
	Username string
	Name string
	Email string
	Surname string
	Website string
	Password string
	RepeatedPassword string
}

type AgentRegistrationDTO struct {
	Username 	string
	Name	string
	Email	 string
	Surname	 string
	Website	string
	Password string
}

func NewAgentRegistrationRequest(registrationRequest *AgentRegistrationRequestDTO, requestStatus RequestStatus) (*AgentRegistrationRequest, error) {
	hashAndSalt, err := HashAndSaltPasswordIfStrongAndMatching(registrationRequest.Password, registrationRequest.RepeatedPassword)
	if err != nil {
		return nil, err
	}

	err = validateRequestStatusTypeEnums(requestStatus)
	if err != nil {
		return nil, err
	}


	return &AgentRegistrationRequest{Id: guid.New().String(),
		Name:   registrationRequest.Name,
		Surname:    registrationRequest.Surname,
		Email: registrationRequest.Email,
		Username: registrationRequest.Username,
		Website: registrationRequest.Website,
		Status: requestStatus,
		Password: hashAndSalt,
	}, nil
}

func HashAndSaltPasswordIfStrongAndMatching(password string, repeatedPassword string) (string, error) {
	isMatching := password == repeatedPassword
	if !isMatching {
		return "", errors.New("passwords are not matching")
	}
	isWeak, _ := regexp.MatchString("^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[^!@#$%^&*(),.?\":{}|<>~'_+=]*)$", password)

	if isWeak {
		return "", errors.New("password must contain minimum eight characters, at least one capital letter, one number and one special character")
	}
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash), err
}

type AgentRegistrationRequestResponseDTO struct {
	Id string
	Name string
	Surname string
	Username string
	Email string
	WebSite string
}