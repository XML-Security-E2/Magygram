package service_contracts

import (
	"context"
	"user-service/domain/model"
)

type CollectionsService interface {
	CreateCollection(ctx context.Context, bearer string, collectionName string) error
	AddPostToCollection(ctx context.Context, bearer string, favouritePostRequest *model.FavouritePostRequest) error
	DeletePostFromCollections(ctx context.Context, bearer string, postId string) error
	GetUsersCollections(ctx context.Context, bearer string, except string) (map[string][]model.IdWithMedia,error)
	CheckIfPostsInFavourites(ctx context.Context, bearer string,postIds *[]string) ([]*model.PostIdFavouritesFlag,error)
}