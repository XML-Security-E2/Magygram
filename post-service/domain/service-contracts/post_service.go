package service_contracts

import (
	"context"
	"post-service/domain/model"
)

type PostService interface {
	GetPostsForTimeline(ctx context.Context, bearer string) ([]*model.PostResponse , error)
	CreatePost(ctx context.Context, bearer string,  post *model.PostRequest) (string, error)
	EditPost(ctx context.Context, bearer string,  post *model.PostEditRequest) error
	LikePost(ctx context.Context, bearer string,  postId string) error
	UnlikePost(ctx context.Context, bearer string,  postId string) error
	DislikePost(ctx context.Context, bearer string,  postId string) error
	UndislikePost(ctx context.Context, bearer string,  postId string) error
	GetPostsFirstImage(ctx context.Context, postId string) (*model.Media, error)
	AddComment(ctx context.Context,  postId string,  content string, bearer string) (*model.Comment, error)
	CheckIfUsersPostFromBearer(bearer string, postOwnerId string) (bool, error)
	SearchForPostsByHashTagByGuest(ctx context.Context,  hashTagValue string) ([]*model.HashTageSearchResponseDTO , error)
	GetPostsByHashTagForGuest(ctx context.Context, hashtag string) ([]*model.GuestTimelinePostResponse , error)
}