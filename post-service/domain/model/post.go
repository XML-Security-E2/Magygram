package model

import (
	"errors"
	"github.com/beevik/guid"
	"mime/multipart"
	"strings"
	"time"
)

/* Za postmana
{
    "Description" : "Description",
    "Location" : "Kraljevo",
    "PostType" : "REGULAR",
    "Media": [
        { "Url":"images/image1", "MediaType": "IMAGE" },
        { "Url":"images/image2", "MediaType": "IMAGE" },
        { "Url":"videos/video1", "MediaType": "VIDEO" }
    ],
    "UserInfo" : {
        "Id" : "2a5a1dc2-f5ad-4eb8-ae69-f4ea160422ff",
        "Username" : "testniusername",
        "ImageURL" : "images/image253"
    }
}
 */

type Post struct {
	Id string `bson:"_id,omitempty"`
	Description string `bson:"description"`
	Location string `bson:"location"`
	ContentType ContentType `bson:"post_type"`
	Tags []Tag `bson:"tags"`
	HashTags []string `bson:"hashTags"`
	Media []Media `bson:"media"`
	UserInfo UserInfo `bson:"user_info"`
	LikedBy []UserInfo `bson:"liked_by"`
	DislikedBy []UserInfo `bson:"disliked_by"`
	Comments []Comment `bson:"comments"`
	IsDeleted bool `bson:"deleted"`
}

type ContentType string

const(
	REGULAR = iota
	CAMPAIGN
)

type PostEditRequest struct {
	Id string `json:"id"`
	Description string `json:"description"`
	Location string `json:"location"`
	Tags []Tag `json:"tags"`
}

type PostRequest struct {
	Description string `json:"description"`
	Location string `json:"location"`
	Media []*multipart.FileHeader `json:"media"`
	Tags []Tag `json:"tags"`
}

type CampaignRequest struct {
	ContentId string `json:"contentId"`
	MinDisplaysForRepeatedly int `json:"minDisplaysForRepeatedly"`
	Frequency CampaignFrequency `json:"frequency"`
	TargetGroup TargetGroup `json:"targetGroup"`
	DisplayTime int `json:"displayTime"`
	DateFrom time.Time `json:"dateFrom"`
	DateTo time.Time `json:"dateTo"`
	ExposeOnceDate time.Time `json:"exposeOnceDate"`
	Type string `json:"campaignType"`
}

type TargetGroup struct {
	MinAge int `json:"minAge"`
	MaxAge int `json:"maxAge"`
	Gender GenderType `json:"gender"`
}

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

func NewPost(postRequest *PostRequest, postOwner UserInfo, postType ContentType, media []Media) (*Post, error) {
	err := validatePostTypeEnums(postType)
	if err != nil {
		return nil, err
	}

	err = validateMediaTypeEnums(media)
	if err != nil {
		return nil, err
	}

	return &Post{Id: guid.New().String(),
		Description:   postRequest.Description,
		Location:    postRequest.Location,
		HashTags: GetHashTagsFromDescription(postRequest.Description),
		UserInfo: postOwner,
		Media: media,
		Tags: postRequest.Tags,
		ContentType : postType,
		LikedBy: []UserInfo{},
		DislikedBy: []UserInfo{},
		Comments: []Comment{},
		IsDeleted: false,
	}, nil
}

func GetHashTagsFromDescription(description string) []string {
	var hashTags []string
	words := strings.Fields(description)
	for _, w := range words {
		if strings.HasPrefix(w, "#") {
			hashTags = append(hashTags, strings.TrimPrefix(w, "#"))
		}
	}
	if hashTags == nil {
		return []string{}
	}
	return hashTags
}

func validatePostTypeEnums(pt ContentType) error {
	switch pt {
	case "REGULAR", "CAMPAIGN":
		return nil
	}
	return errors.New("Invalid post type")
}

func validateMediaTypeEnums(md []Media) error {
	if len(md)==0 {
		return nil
	}

	for _, media := range md {
		if media.MediaType != "IMAGE" && media.MediaType !="VIDEO" {
			return errors.New("Invalid media type")
		}
	}

	return nil
}

type Media struct {
	Url string
	MediaType string
}

type MediaType string

const(
	IMAGE = iota
	VIDEO
)

type UserInfo struct {
	Id string
	Username string
	ImageURL string
}

type UserInfoEdit struct {
	Id string
	Username string
	ImageURL string
	PostIds []string
}

type Comment struct {
	Id string
	CreatedBy UserInfo
	Content string
	TimeCreated time.Time
	Tags []Tag
}

type PostResponse struct {
	Id string
	Description string
	Location string
	ContentType ContentType
	Tags []Tag
	HashTags []string
	Media []Media
	UserInfo UserInfo
	LikedBy []UserInfo
	DislikedBy []UserInfo
	Comments []Comment
	Favourites bool
	Liked bool
	Disliked bool
}

type PostProfileResponse struct {
	Id string `json:"id"`
	Media Media `json:"media"`
}

type PostIdFavouritesFlag struct {
	Id string `json:"id"`
	Favourites bool `json:"favourites"`
}

func NewPostResponse(post *Post, liked bool, disliked bool, favourites bool) (*PostResponse, error) {
	return &PostResponse{Id: post.Id,
		Description:   post.Description,
		Location:    post.Location,
		HashTags: post.HashTags,
		Media: post.Media,
		UserInfo: post.UserInfo,
		LikedBy: post.LikedBy,
		DislikedBy: post.DislikedBy,
		Comments: post.Comments,
		Liked: liked,
		Tags: post.Tags,
		Disliked: disliked,
		Favourites: favourites,
		ContentType: post.ContentType,
	}, nil
}

type PostId struct {
	Id string
}

type CommentRequest struct {
	PostId string
	Content string
	Tags []Tag
}

type Location struct {
	Id string `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"username"`
}

type Tag struct {
	Id string `bson:"_id,omitempty"`
	Username string `bson:"username"`
}

type FollowedUsersResponse struct {
	Users []string
}

type HashTageSearchResponseDTO struct {
	Hashtag string
	NumberOfPosts int
}

type LocationSearchResponseDTO struct {
	Hashtag string
	NumberOfPosts int
}

type GuestTimelinePostResponse struct {
	Id string
	Description string
	Location string
	Media []Media
	UserInfo UserInfo
}

func NewGuestTimelinePostResponse(post *Post) (*GuestTimelinePostResponse, error) {
	return &GuestTimelinePostResponse{Id: post.Id,
		Description:   post.Description,
		Location:    post.Location,
		Media: post.Media,
		UserInfo: post.UserInfo,
	}, nil
}