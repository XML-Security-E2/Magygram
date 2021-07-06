package model

type CampaignRequest struct {
	ProductId string `json:"productId"`
	ExposeOnceDate int64 `json:"exposeOnceDate"`
	MinDisplaysForRepeatedly int `json:"minDisplays"`
	Type ContentType `json:"campaignType"`
	Frequency CampaignFrequency `json:"frequency"`
	TargetGroup TargetGroup `json:"targetGroup"`
	DateFrom int64 `json:"startDate"`
	DateTo int64 `json:"endDate"`
	DisplayTime int `json:"displayTime"`
}

type ContentType string
type CampaignFrequency string

type TargetGroup struct {
	MinAge int `bson:"min_age" json:"minAge" validate:"required,numeric,min=12,max=70"`
	MaxAge int `bson:"max_age" json:"maxAge" validate:"required,numeric,min=12,max=120"`
	Gender GenderType `bson:"gender" json:"gender"`
}

type GenderType string

type CampaignApiRequest struct {
	MinDisplaysForRepeatedly int `json:"minDisplaysForRepeatedly"`
	Frequency CampaignFrequency `json:"frequency"`
	TargetGroup TargetGroup `json:"targetGroup"`
	DisplayTime int `json:"displayTime"`
	DateFrom int64 `json:"dateFrom"`
	DateTo int64 `json:"dateTo"`
	ExposeOnceDate int64 `json:"exposeOnceDate"`
	Type ContentType `json:"campaignType"`
	FilePath string `json:"media"`
}

