package model

import (
	"github.com/beevik/guid"
)

type Post struct {
	Id string `bson:"_id,omitempty"`
	Description string `bson:"description"`
	Location string `bson:"location"`
}

type PostRequest struct {
	Description string `bson:"description"`
	Location string `bson:"location"`
}

func NewPost(postRequest *PostRequest) (*Post, error) {
	return &Post{Id: guid.New().String(),
		Description:   postRequest.Description,
		Location:    postRequest.Location}, nil
}