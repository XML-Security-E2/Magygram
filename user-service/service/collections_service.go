package service

import (
	"context"
	"errors"
	"fmt"
	"user-service/domain/model"
	"user-service/domain/repository"
	"user-service/domain/service-contracts"
	"user-service/service/intercomm"
)

type collectionsService struct {
	repository.UserRepository
	intercomm.AuthClient
	intercomm.PostClient
}

func NewCollectionsService(r repository.UserRepository, ic 	intercomm.AuthClient, pc intercomm.PostClient) service_contracts.CollectionsService {
	return &collectionsService{r, ic, pc}
}

func (c collectionsService) CreateCollection(ctx context.Context, bearer string, collectionName string) error {

	userId, err := c.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	user, err := c.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return errors.New("invalid user id")
	}

	if _, ok := user.FavouritePosts[collectionName]; ok {
		return errors.New(fmt.Sprintf("collection with name %s already exist", collectionName))
	}

	user.FavouritePosts[collectionName] = []model.Media{{}}
	_, err = c.UserRepository.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (c collectionsService) AddPostToCollection(ctx context.Context, bearer string, favouritePostRequest *model.FavouritePostRequest) error {

	userId, err := c.AuthClient.GetLoggedUserId(bearer)
	if err != nil {
		return err
	}

	user, err := c.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return errors.New("invalid user id")
	}

	if _, ok := user.FavouritePosts[favouritePostRequest.CollectionName]; !ok {
		return errors.New(fmt.Sprintf("invalid %s collection", favouritePostRequest.CollectionName))
	}

	postImage, err := c.PostClient.GetPostsFirstImage(favouritePostRequest.PostId)
	if err != nil {
		return err
	}

	user.FavouritePosts[favouritePostRequest.CollectionName] = append(user.FavouritePosts[favouritePostRequest.CollectionName], *postImage)
	_, err = c.UserRepository.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}