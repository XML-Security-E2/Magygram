package service

import (
	"context"
	"github.com/go-playground/validator"
	_ "net/http"
	"post-service/domain/model"
	"post-service/domain/repository"
	"post-service/domain/service-contracts"
	"post-service/service/intercomm"
)



type postService struct {
	repository.PostRepository
	intercomm.MediaClient
	intercomm.UserClient
}

func NewPostService(r repository.PostRepository, ic intercomm.MediaClient, uc intercomm.UserClient) service_contracts.PostService {
	return &postService{r , ic, uc}
}

func (p postService) CreatePost(ctx context.Context, bearer string, postRequest *model.PostRequest) (string, error) {

	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil { return "", err}

	media, err := p.MediaClient.SaveMedia(postRequest.Media)
	if err != nil { return "", err}

	post, err := model.NewPost(postRequest, *userInfo, "REGULAR", media)

	if err != nil { return "", err}

	if err := validator.New().Struct(post); err!= nil {
		return "", err
	}

	result, err := p.PostRepository.Create(ctx, post)

	if err != nil { return "", err}
	if postId, ok := result.InsertedID.(string); ok {
		if err != nil { return "", err}
		return postId, nil
	}

	return "", err
}

func (p postService) GetPostsForTimeline(ctx context.Context, bearer string) ([]*model.PostResponse, error) {
	result, err := p.PostRepository.GetAll(ctx)

	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil {
		return nil, err
	}

	retVal := mapPostsToResponsePostDTO(result, userInfo.Id)


	return retVal, nil
}

func (p postService) LikePost(ctx context.Context, bearer string, postId string) error {
	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil {
		return err
	}

	result, err := p.PostRepository.GetOne(ctx,postId)
	if err != nil {
		return err
	}
	var res model.UserInfo

	res.Id= userInfo.Id
	res.ImageURL= userInfo.ImageURL
	res.Username= userInfo.Username

	result.LikedBy = append(result.LikedBy, res)

	_, err = p.PostRepository.Update(ctx,result)
	if err != nil {
		return err
	}

	return nil
}

func (p postService) UnlikePost(ctx context.Context, bearer string, postId string) error {
	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil {
		return err
	}

	result, err := p.PostRepository.GetOne(ctx,postId)
	if err != nil {
		return err
	}

	result.LikedBy = findAndDeleteLikedBy(result, userInfo)

	_, err = p.PostRepository.Update(ctx,result)
	if err != nil {
		return err
	}

	return nil
}

func findAndDeleteLikedBy(result *model.Post, info *model.UserInfo) []model.UserInfo {
	index := 0
	for _, i := range result.LikedBy {
		if i.Id != info.Id {
			result.LikedBy[index] = i
			index++
		}
	}
	return result.LikedBy[:index]
}



func findAndDeleteLikedBy1(s [4]int, item int) []int {
	index := 0
	for _, i := range s {
		if i != item {
			s[index] = i
			index++
		}
	}
	return s[:index]
}

func mapPostsToResponsePostDTO(result []*model.Post, userId string) []*model.PostResponse {
	var retVal []*model.PostResponse
	
	for _, post := range result {
		res, err := model.NewPostResponse(post,hasUserLikedPost(post,userId),hasUserDislikedPost(post,userId))

		if err != nil { return nil}

		retVal = append(retVal, res)
	}

	return retVal
}

func hasUserLikedPost(post *model.Post, usedId string) bool {
	var retVal = false

	for _, likedUserInfo := range post.LikedBy {
		if likedUserInfo.Id == usedId{
			retVal=true
			break
		}
	}

	return retVal
}

func hasUserDislikedPost(post *model.Post, usedId string) bool {
	var retVal = false

	for _, dislikedUserInfo := range post.DislikedBy {
		if dislikedUserInfo.Id == usedId{
			retVal=true
			break
		}
	}

	return retVal
}



