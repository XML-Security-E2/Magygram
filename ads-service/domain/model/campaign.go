package model

import (
	"errors"
	"github.com/beevik/guid"
	"time"
)

type InfluencerCampaign struct {
	Id string `bson:"_id,omitempty"`
	UserId string `bson:"user_id,omitempty"`
	ContentId string `bson:"content_id,omitempty"`
	ParentCampaignId string `bson:"parent_campaign_id,omitempty"`
	SeenBy []string `bson:"seen_by"`
	Type ContentType `bson:"campaign_type"`
}

type InfluencerCampaignRequest struct {
	ContentId string `json:"contentId"`
	ParentCampaignId string `json:"parentCampaignId"`
	Type ContentType `json:"campaignType"`
}

type Campaign struct {
	Id string `bson:"_id,omitempty"`
	OwnerId string `bson:"owner_id,omitempty"`
	ContentId string `bson:"content_id,omitempty"`
	MinDisplaysForRepeatedly int `bson:"min_displays_for_repeatedly"`
	SeenBy []string `bson:"seen_by"`
	Type ContentType `bson:"campaign_type"`
	Frequency CampaignFrequency `bson:"frequency"`
	TargetGroup TargetGroup `bson:"target_group"`
	DateFrom time.Time `bson:"date_from"`
	DateTo time.Time `bson:"date_to"`
	ExposeOnceDate time.Time `bson:"expose_once_date"`
	Deleted bool `bson:"deleted"`
	DisplayTime int `bson:"display_time" validate:"required,numeric,min=0,max=24"`
}

type CampaignUpdateRequest struct {
	Id string `bson:"_id,omitempty"`
	CampaignId string `bson:"campaign_id,omitempty"`
	MinDisplaysForRepeatedly int `bson:"min_displays_for_repeatedly"`
	TargetGroup TargetGroup `bson:"target_group"`
	DateFrom time.Time `bson:"date_from"`
	DateTo time.Time `bson:"date_to"`
	RequestedDate time.Time `bson:"requested_date"`
	CampaignUpdateStatus CampaignUpdateStatus `bson:"campaign_update_status"`
}

type CampaignUpdateRequestDTO struct {
	CampaignId string `bson:"campaign_id,omitempty" json:"campaignId"`
	MinDisplaysForRepeatedly int `bson:"min_displays_for_repeatedly" json:"minDisplaysForRepeatedly"`
	TargetGroup TargetGroup `bson:"target_group" json:"targetGroup"`
	DateFrom time.Time `bson:"date_from" json:"dateFrom"`
	DateTo time.Time `bson:"date_to" json:"dateTo"`
}

type CampaignUpdateRequestTimeDTO struct {
	CampaignId string `bson:"campaign_id,omitempty" json:"campaignId"`
	MinDisplaysForRepeatedly int `bson:"min_displays_for_repeatedly" json:"minDisplaysForRepeatedly"`
	TargetGroup TargetGroup `bson:"target_group" json:"targetGroup"`
	DateFrom int64 `bson:"date_from" json:"dateFrom"`
	DateTo int64 `bson:"date_to" json:"dateTo"`
}

type CampaignRequest struct {
	ContentId string `json:"contentId"`
	ExposeOnceDate time.Time `json:"exposeOnceDate"`
	MinDisplaysForRepeatedly int `json:"minDisplaysForRepeatedly"`
	Type ContentType `json:"campaignType"`
	Frequency CampaignFrequency `json:"frequency"`
	TargetGroup TargetGroup `json:"targetGroup"`
	DateFrom time.Time `json:"dateFrom"`
	DateTo time.Time `json:"dateTo"`
	DisplayTime int `json:"displayTime"`
}

type CampaignRetreiveRequest struct {
	Id string `json:"id"`
	ExposeOnceDate time.Time `json:"exposeOnceDate"`
	MinDisplaysForRepeatedly int `json:"minDisplaysForRepeatedly"`
	Type ContentType `json:"campaignType"`
	Frequency CampaignFrequency `json:"frequency"`
	TargetGroup TargetGroup `json:"targetGroup"`
	DateFrom time.Time `json:"dateFrom"`
	DateTo time.Time `json:"dateTo"`
	DisplayTime int `json:"displayTime"`
}

type InfluencerContent struct {
	InfluencerId string `bson:"influencer_id"`
	ContentId int `bson:"content_id"`
	ContentType ContentType `bson:"content_type"`
}

type CampaignUpdateStatus string

const (
	PENDING = iota
	DONE
)

type ContentType string

const (
	POST = iota
	STORY
)

type TargetGroup struct {
	MinAge int `bson:"min_age" json:"minAge" validate:"required,numeric,min=12,max=70"`
	MaxAge int `bson:"max_age" json:"maxAge" validate:"required,numeric,min=12,max=120"`
	Gender GenderType `bson:"gender" json:"gender"`
}

type UserTargetGroup struct {
	Id string `json:"id"`
	Age int `json:"age"`
	Gender Gender `json:"gender"`
}

type Gender string

type GenderType string

const (
	MALE = iota
	FEMALE
	ANY
)

type CampaignFrequency string

const(
	ONCE = iota
	REPEATEDLY
)

func NewCampaignUpdateRequest(campaignRequest *CampaignUpdateRequestDTO) (*CampaignUpdateRequest, error) {
	if err := validateGenderTypeEnums(campaignRequest.TargetGroup.Gender); err != nil {
		return nil, err
	}

	if campaignRequest.DateFrom.After(campaignRequest.DateTo) {
		return nil, errors.New("dates out of range")
	}
	if campaignRequest.TargetGroup.MinAge > campaignRequest.TargetGroup.MaxAge || campaignRequest.TargetGroup.MinAge < 12 || campaignRequest.TargetGroup.MaxAge > 120 {
		return nil, errors.New("age out of range")
	}

	yf,mf,df := campaignRequest.DateFrom.Date()
	timeef := time.Date(yf,mf,df,2,0,1,0, time.Local)

	yt,mt,dt := campaignRequest.DateTo.Date()
	timeet := time.Date(yt,mt,dt,2,0,1,0, time.Local)

	return &CampaignUpdateRequest{
		Id:                       guid.New().String(),
		CampaignId:               campaignRequest.CampaignId,
		MinDisplaysForRepeatedly: campaignRequest.MinDisplaysForRepeatedly,
		TargetGroup:              campaignRequest.TargetGroup,
		DateFrom:                 timeef,
		DateTo:                   timeet,
		RequestedDate:            time.Now(),
		CampaignUpdateStatus:     "PENDING",
	}, nil
}

func NewCampaign(campaignRequest *CampaignRequest, ownerId string) (*Campaign, error) {
	if err := validateContentTypeEnums(campaignRequest.Type); err != nil {
		return nil, err
	}
	if err := validateCampaignFrequencyEnums(campaignRequest.Frequency); err != nil {
		return nil, err
	}
	if err := validateGenderTypeEnums(campaignRequest.TargetGroup.Gender); err != nil {
		return nil, err
	}

	if campaignRequest.DateFrom.After(campaignRequest.DateTo) {
		return nil, errors.New("dates out of range")
	}
	if campaignRequest.TargetGroup.MinAge > campaignRequest.TargetGroup.MaxAge || campaignRequest.TargetGroup.MinAge < 12 || campaignRequest.TargetGroup.MaxAge > 120 {
		return nil, errors.New("age out of range")
	}


	yf,mf,df := campaignRequest.DateFrom.Date()
	timeef := time.Date(yf,mf,df,2,0,1,0, time.Local)

	yt,mt,dt := campaignRequest.DateTo.Date()
	timeet := time.Date(yt,mt,dt,2,0,1,0, time.Local)

	yo,mo,do := campaignRequest.ExposeOnceDate.Date()
	timeeo := time.Date(yo,mo,do,2,0,1,0, time.Local)

	return &Campaign{
		Id:                       guid.New().String(),
		ContentId:                campaignRequest.ContentId,
		MinDisplaysForRepeatedly: campaignRequest.MinDisplaysForRepeatedly,
		SeenBy:                   []string{},
		Type:                     campaignRequest.Type,
		Frequency:                campaignRequest.Frequency,
		TargetGroup:              campaignRequest.TargetGroup,
		DateFrom:                 timeef,
		DateTo:                   timeet,
		OwnerId: 				  ownerId,
		DisplayTime:              campaignRequest.DisplayTime,
		ExposeOnceDate:           timeeo,
		Deleted:                  false,
	}, nil
}

func NewInfluencerCampaign(campaignRequest *InfluencerCampaignRequest, userId string) (*InfluencerCampaign, error) {
	if err := validateContentTypeEnums(campaignRequest.Type); err != nil {
		return nil, err
	}

	return &InfluencerCampaign{
		Id:               guid.New().String(),
		UserId:           userId,
		ContentId:        campaignRequest.ContentId,
		ParentCampaignId: campaignRequest.ParentCampaignId,
		SeenBy:           []string{},
		Type:             campaignRequest.Type,
	}, nil
}

func validateGenderTypeEnums(pt GenderType) error {
	switch pt {
	case "MALE", "FEMALE", "ANY":
		return nil
	}
	return errors.New("invalid gender type")
}

func validateCampaignFrequencyEnums(pt CampaignFrequency) error {
	switch pt {
	case "ONCE", "REPEATEDLY":
		return nil
	}
	return errors.New("invalid campaign frequency")
}

func validateContentTypeEnums(pt ContentType) error {
	switch pt {
	case "POST", "STORY":
		return nil
	}
	return errors.New("invalid campaign content type")
}