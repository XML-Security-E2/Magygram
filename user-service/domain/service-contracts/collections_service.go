package service_contracts

import (
	"context"
	"user-service/domain/model"
)

type CollectionsService interface {
	CreateCollection(ctx context.Context, bearer string, collectionName string) error
	AddPostToCollection(ctx context.Context, bearer string, favouritePostRequest *model.FavouritePostRequest) error
}