package model

import (
	"encoding/xml"
	"time"
)

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

type CampaignStatisticResponse struct {
	Id string `json:"id"`
	ExposeOnceDate time.Time `json:"exposeOnceDate"`
	MinDisplaysForRepeatedly int `json:"minDisplaysForRepeatedly"`
	Type ContentType `json:"campaignType"`
	Frequency CampaignFrequency `json:"frequency"`
	UserViews int `json:"userViews"`
	WebsiteClicks int `json:"websiteClicks"`
	TargetGroup TargetGroup `json:"targetGroup"`
	DateFrom time.Time `json:"dateFrom"`
	DateTo time.Time `json:"dateTo"`
	DisplayTime int `json:"displayTime"`
	CampaignStatus CampaignStatisticStatus `json:"campaignStatus"`
	InfluencerUsername string `json:"influencerUsername"`
	InfluencerId string `json:"influencerId"`
	Media Media `json:"media"`
	Website string `json:"website"`
	Likes int
	Dislikes int
	Comments int
	StoryViews int
	DailyAverage float32 `json:"dailyAverage"`
	Activity CampaignStatisticActivity `json:"activity"`
}


type CampaignStatisticReport struct {
	XMLName        xml.Name       `xml:"report"`
	FileId         string         `json:"fileId" xml:"id,attr"`
	DateCreating   time.Time	  `json:"dateCreating" xml:"date-creation,attr"`
	Campaigns 	   []CampaignStatisticInfo `json:"campaigns" xml:"campaign"`
}

type CampaignStatisticInfo struct {
	ExposeOnceDate time.Time `json:"exposeOnceDate" xml:"exposeOnceDate"`
	MinDisplaysForRepeatedly int `json:"minDisplaysForRepeatedly" xml:"minDisplaysForRepeatedly"`
	Type ContentType `json:"campaignType" xml:"campaignType"`
	Frequency CampaignFrequency `json:"frequency" xml:"CampaignFrequency"`
	UserViews int `json:"userViews"  xml:"userViews"`
	WebsiteClicks int `json:"websiteClicks" xml:"websiteClicks"`
	TargetGroup TargetGroup `json:"targetGroup" xml:"targetGroup"`
	DateFrom time.Time `json:"dateFrom" xml:"dateFrom"`
	DateTo time.Time `json:"dateTo"  xml:"dateTo"`
	DisplayTime int `json:"displayTime" xml:"displayTime"`
	CampaignStatus CampaignStatisticStatus `json:"campaignStatus" xml:"CampaignStatisticStatus"`
	InfluencerUsername string `json:"influencerUsername" xml:"influencerUsername,omitempty"`
	Media Media `json:"media" xml:"media"`
	Website string `json:"website" xml:"website"`
	Likes int `xml:"likes"`
	Dislikes int `xml:"dislikes"`
	Comments int `xml:"comments"`
	StoryViews int `xml:"storyViews"`
	DailyAverage float32 `json:"dailyAverage" xml:"dailyAverage"`
	Activity CampaignStatisticActivity `json:"activity" xml:"activity"`
}

type XmlDatabaseResponse struct {
	XmlName xml.Name `xml:"exist:result"`
	Url string `xml:"exist,attr"`
	XmlDatabaseCollections XmlDatabaseCollection `xml:"collection"`
}

type XmlDatabaseCollection struct {
	Name string `xml:"name,attr"`
	Created string `xml:"created,attr"`
	Owner string `xml:"owner,attr"`
	Group string `xml:"group,attr"`
	Permissions string `xml:"permissions,attr"`
	Resources []XmlDatabaseResource `xml:"resource,omitempty"`
}

type XmlDatabaseResource struct {
	Name string `xml:"name,attr"`
	Created string `xml:"created,attr"`
	LastModified string `xml:"last-modified,attr"`
	Owner string `xml:"owner,attr"`
	Group string `xml:"group,attr"`
	Permissions string `xml:"permissions,attr"`
}

type CampaignStatisticStatus string

const (
	INFLUENCER = iota
	REGULAR
)

type CampaignStatisticActivity string

const (
	ACTIVE = iota
	UNACTIVE
)

type Media struct {
	Url string `xml:"url"`
	MediaType string `xml:"mediaType"`
}

type MediaType string

const(
	IMAGE = iota
	VIDEO
)