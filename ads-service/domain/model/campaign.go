package model

import "time"

type Campaign struct {
	Id string `bson:"_id,omitempty"`
	ContentId string `bson:"content_id,omitempty"`
	MinDisplays int `bson:"min_displays"`
	SeenBy []string `bson:"seen_by"`
	Type CampaignType `bson:"campaign_type"`
	Frequency CampaignFrequency `bson:"frequency"`
	TargetGroup TargetGroup `bson:"target_group"`
	DateFrom time.Time `bson:"date_from"`
	DateTo time.Time `bson:"date_to"`
	InfluencersContent []InfluencerContent `bson:"influencers_content"`
}

type InfluencerContent struct {
	InfluencerId string `bson:"influencer_id"`
	ContentId int `bson:"content_id"`
	ContentType ContentType `bson:"content_type"`
}

type ContentType string

const (
	POST = iota
	STORY
)

type TargetGroup struct {
	MinAge int `bson:"min_age"`
	MaxAge int `bson:"max_age"`
	Gender GenderType `bson:"gender"`
}

type GenderType string

const (
	MALE = iota
	FEMALE
)


type CampaignType string

const(
	REGULAR = iota
	CAMPAIGN
)

type CampaignFrequency string

const(
	ONCE = iota
	REPEATEDLY
)
