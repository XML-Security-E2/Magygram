package service_contracts

import (
	"context"
	"post-service/domain/model"
)

type PostService interface {
	GetPostsForTimeline(ctx context.Context, bearer string) ([]*model.PostResponse , error)
	GetPostById(ctx context.Context, bearer string, postId string) (*model.PostResponse , error)
	GetPostForMessagesById(ctx context.Context, bearer string, postId string) (*model.PostResponse, *model.UserInfo, error)
	CreatePost(ctx context.Context, bearer string,  post *model.PostRequest) (string, error)
	CreatePostInfluencer(ctx context.Context, bearer string, request *model.InfluencerRequest) (string, error)
	CreatePostCampaign(ctx context.Context, bearer string,  post *model.PostRequest, campaignReq *model.CampaignRequest) (string, error)
	CreatePostCampaignFromApi(ctx context.Context, bearer string,  post *model.PostRequest) (string, error)
	GetUserPostCampaigns(ctx context.Context, bearer string) ([]*model.PostProfileResponse, error)

	EditPost(ctx context.Context, bearer string,  post *model.PostEditRequest) error
	LikePost(ctx context.Context, bearer string,  postId string) error
	UnlikePost(ctx context.Context, bearer string,  postId string) error
	DislikePost(ctx context.Context, bearer string,  postId string) error
	UndislikePost(ctx context.Context, bearer string,  postId string) error
	GetPostsFirstImage(ctx context.Context, postId string) (*model.Media, error)
	AddComment(ctx context.Context,  postId string,  content string, bearer string, tags []model.Tag) (*model.Comment, error)
	CheckIfUsersPostFromBearer(ctx context.Context, bearer string, postOwnerId string) (bool, error)
	GetUsersPosts(ctx context.Context, bearer string, postOwnerId string) ([]*model.PostProfileResponse, error)
	GetUsersPostsCount(ctx context.Context, postOwnerId string) (int, error)
	SearchForPostsByHashTagByGuest(ctx context.Context,  hashTagValue string) ([]*model.HashTageSearchResponseDTO , error)
	GetPostsByHashTagForGuest(ctx context.Context, hashtag string) ([]*model.GuestTimelinePostResponse , error)
	GetPostForUserTimelineByHashTag(ctx context.Context, hashtag string,bearer string) ([]*model.PostResponse , error)
	SearchPostsByLocation(ctx context.Context,  locationValue string) ([]*model.LocationSearchResponseDTO , error)
	GetPostForGuestTimelineByLocation(ctx context.Context, location string) ([]*model.GuestTimelinePostResponse , error)
	GetPostForUserTimelineByLocation(ctx context.Context, location string, bearer string) ([]*model.PostResponse , error)
	GetPostByIdForGuest(ctx context.Context, postId string) (*model.GuestTimelinePostResponse , error)
	GetUserLikedPosts(ctx context.Context, bearer string) ([]*model.PostProfileResponse, error)
	GetUserDislikedPosts(ctx context.Context, bearer string) ([]*model.PostProfileResponse, error)
	DeletePost(ctx context.Context, bearer string, requestId string) error
	EditPostOwnerInfo(ctx context.Context, bearer string, userInfo *model.UserInfo) error
	EditLikedByInfo(ctx context.Context, bearer string, userInfoEdit *model.UserInfoEdit) error
	EditDislikedByInfo(ctx context.Context, bearer string, userInfoEdit *model.UserInfoEdit) error
	EditCommentedByInfo(ctx context.Context, bearer string, userInfoEdit *model.UserInfoEdit) error

	GetPostsMediaAndWebsiteByIds(ctx context.Context, ids *model.FollowedUsersResponse) ([]*model.IdMediaWebsiteResponse, error)
}