package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/beevik/guid"
	"github.com/go-playground/validator"
	"log"
	_ "net/http"
	"post-service/domain/model"
	"post-service/domain/repository"
	"post-service/domain/service-contracts"
	"post-service/domain/service-contracts/exceptions"
	"post-service/service/intercomm"
	"strings"
	"time"
)



type postService struct {
	repository.PostRepository
	intercomm.MediaClient
	intercomm.UserClient
	intercomm.RelationshipClient
}

func NewPostService(r repository.PostRepository, ic intercomm.MediaClient, uc intercomm.UserClient, ir intercomm.RelationshipClient) service_contracts.PostService {
	return &postService{r , ic, uc, ir}
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

	var posts []*model.Post
	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil {
		return nil, err
	}

	var followedUsers model.FollowedUsersResponse
	followedUsers, err = p.RelationshipClient.GetFollowedUsers(userInfo.Id)
	if err != nil {
		return nil, err
	}

	for _, userId := range followedUsers.Users {
		var newPosts []*model.Post
		newPosts, _ = p.PostRepository.GetPostsForUser(ctx,userId)
		posts= append(posts, newPosts...)
	}

	retVal := p.mapPostsToResponsePostDTO(bearer, posts, userInfo.Id)

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

func (p postService) DislikePost(ctx context.Context, bearer string, postId string) error {
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

	result.DislikedBy = append(result.DislikedBy, res)

	_, err = p.PostRepository.Update(ctx,result)
	if err != nil {
		return err
	}

	return nil
}

func (p postService) UndislikePost(ctx context.Context, bearer string, postId string) error {

	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil {
		return err
	}
	fmt.Println(postId)

	result, err := p.PostRepository.GetOne(ctx,postId)
	if err != nil {
		return err
	}

	result.DislikedBy = findAndDeleteDislikedBy(result, userInfo)

	_, err = p.PostRepository.Update(ctx,result)
	if err != nil {
		return err
	}

	return nil
}

func (p postService) AddComment(ctx context.Context, postId string, content string, bearer string) (*model.Comment, error) {
	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil {
		return nil, err
	}

	result, err := p.PostRepository.GetOne(ctx,postId)
	if err != nil {
		return nil,err
	}

	var res model.Comment

	res.Id= guid.New().String()
	res.Content= content
	res.CreatedBy= *userInfo
	res.TimeCreated = time.Now()

	result.Comments = append(result.Comments, res)

	_, err = p.PostRepository.Update(ctx,result)
	if err != nil {
		return nil, err
	}

	return &res, nil
}


func findAndDeleteDislikedBy(result *model.Post, info *model.UserInfo) []model.UserInfo {
	index := 0
	for _, i := range result.DislikedBy {
		if i.Id != info.Id {
			result.DislikedBy[index] = i
			index++
		}
	}
	return result.DislikedBy[:index]
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

func (p postService) mapPostsToResponsePostDTO(bearer string, result []*model.Post, userId string) []*model.PostResponse {
	var retVal []*model.PostResponse

	postIdFavourites, err := p.UserClient.MapPostsToFavourites(bearer, getIdsFromPosts(result))
	if err != nil { return nil}

	for _, post := range result {
		res, err := model.NewPostResponse(post,hasUserLikedPost(post,userId),hasUserDislikedPost(post,userId), isInFavourites(post, postIdFavourites))

		if err != nil { return nil}

		retVal = append(retVal, res)
	}

	return retVal
}

func isInFavourites(post *model.Post, favourites []*model.PostIdFavouritesFlag) bool {
	for _, postFav := range favourites {
		if post.Id == postFav.Id {
			return postFav.Favourites
		}
	}
	return false
}

func getIdsFromPosts(posts []*model.Post) []string {
	var ids []string
	for _, post := range posts {
		ids = append(ids, post.Id)
	}
	return ids
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


func (p postService) GetPostsFirstImage(ctx context.Context, postId string) (*model.Media, error) {

	post, err := p.PostRepository.GetByID(ctx, postId)

	if err != nil {
		return nil, errors.New("invalid post id")
	}
	if len(post.Media) > 0 {
		return &post.Media[0], nil
	}
	return nil, nil
}

func (p postService) EditPost(ctx context.Context, bearer string, postRequest *model.PostEditRequest) error {
	post, err := p.PostRepository.GetByID(ctx, postRequest.Id)
	if err != nil {
		return errors.New("invalid post id")
	}

	isOwner, err := p.CheckIfUsersPostFromBearer(bearer, post.UserInfo.Id)
	if err != nil {
		return err
	}

	if !isOwner {
		return &exceptions.UnauthorizedAccessError{Msg: "User not authorized"}
	}

	post.Tags = postRequest.Tags
	post.Description = postRequest.Description
	post.Location = postRequest.Location
	post.HashTags = model.GetHashTagsFromDescription(postRequest.Description)

	_, err = p.PostRepository.Update(ctx, post)
	if err != nil {
		return err
	}

	return nil
}

func (p postService) CheckIfUsersPostFromBearer(bearer string, postOwnerId string) (bool, error) {
	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil {
		return false, err
	}

	if postOwnerId != userInfo.Id {
		return false, nil
	}
	return true, nil
}

func (p postService) SearchForPostsByHashTagByGuest(ctx context.Context, hashTagValue string) ([]*model.HashTageSearchResponseDTO, error) {
	posts, err := p.PostRepository.GetPostsThatContainHashTag(ctx, hashTagValue)
	if err != nil {
		return nil, errors.New("Couldn't find any posts")
	}

	elementMap := makeHashTagMap(posts, hashTagValue)

	retVal := mapHashTagMapToHashTagResponseDTO(elementMap)
	
	return retVal, err
}

func mapHashTagMapToHashTagResponseDTO(hashTagMap map[string]int) []*model.HashTageSearchResponseDTO {
	var retVal []*model.HashTageSearchResponseDTO

	for key, element := range hashTagMap {
		res := model.HashTageSearchResponseDTO{Hashtag: key, NumberOfPosts: element}
		retVal = append(retVal, &res)
	}

	return retVal
}

func makeHashTagMap(posts []*model.Post, hashTagValue string) map[string]int {
	elementMap := make(map[string]int)

	for _, post := range posts{
		for _, hashTag := range post.HashTags {
			if strings.Contains(hashTag,hashTagValue){
				if _, ok := elementMap[hashTag]; ok {
					elementMap[hashTag] = elementMap[hashTag]+1
				}else{
					elementMap[hashTag] = 1
				}
			}
		}
	}

	return elementMap
}

func (p postService) GetPostsByHashTagForGuest(ctx context.Context, hashtag string) ([]*model.GuestTimelinePostResponse, error) {
	posts, err := p.PostRepository.GetPostsByHashTag(ctx, hashtag)

	if err!=nil{
		return nil,err
	}

	var publicPosts []*model.Post

	for _,post := range posts{
		value, err := p.UserClient.IsProfilePrivate(post.UserInfo.Id)

		if err!=nil{
			log.Println(err)
			return nil,err
		}

		if !value {
			publicPosts=append(publicPosts, post)
		}
	}

	retVal := p.mapPostsForGuestTimelineToResponseGuestTimelinePostDTO(publicPosts)

	return retVal, nil
}

func (p postService) mapPostsForGuestTimelineToResponseGuestTimelinePostDTO(posts []*model.Post) []*model.GuestTimelinePostResponse {
	var retVal []*model.GuestTimelinePostResponse

	for _, post := range posts {
		res, err := model.NewGuestTimelinePostResponse(post)

		if err != nil { return nil}

		retVal = append(retVal, res)
	}

	return retVal
}

func (p postService) GetPostForUserTimelineByHashTag(ctx context.Context, hashtag string, bearer string) ([]*model.PostResponse, error) {
	posts, err := p.PostRepository.GetPostsByHashTag(ctx, hashtag)
	if err!=nil{
		return nil,err
	}

	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil {
		return nil, err
	}

	var publicPosts []*model.Post
	for _,post := range posts{
		value, err := p.UserClient.IsProfilePrivate(post.UserInfo.Id)
		if err!=nil{
			return nil,err
		}

		if !value {
			publicPosts=append(publicPosts, post)
		}
	}

	retVal := p.mapPostsToResponsePostDTO(bearer, publicPosts, userInfo.Id)

	return retVal, nil
}

func (p postService) SearchPostsByLocation(ctx context.Context, locationValue string) ([]*model.LocationSearchResponseDTO, error) {
	posts, err := p.PostRepository.GetPostsThatContainLocation(ctx, locationValue)
	if err != nil {
		return nil, errors.New("Couldn't find any posts")
	}

	elementMap := makeLocationMap(posts, locationValue)

	retVal := mapLocationMapToLocationResponseDTO(elementMap)

	return retVal, err
}

func makeLocationMap(posts []*model.Post, locationValue string) map[string]int {
	elementMap := make(map[string]int)

	for _, post := range posts{
		if strings.Contains(post.Location,locationValue){
			if _, ok := elementMap[post.Location]; ok {
				elementMap[post.Location] = elementMap[post.Location]+1
			}else {
				elementMap[post.Location] = 1
			}
		}
	}

	return elementMap
}

func mapLocationMapToLocationResponseDTO(hashTagMap map[string]int) []*model.LocationSearchResponseDTO {
	var retVal []*model.LocationSearchResponseDTO

	for key, element := range hashTagMap {
		res := model.LocationSearchResponseDTO{Hashtag: key, NumberOfPosts: element}
		retVal = append(retVal, &res)
	}

	return retVal
}

func (p postService) GetPostForGuestTimelineByLocation(ctx context.Context, location string) ([]*model.GuestTimelinePostResponse, error) {
	posts, err := p.PostRepository.GetPostsByLocation(ctx, location)

	if err!=nil{
		return nil,err
	}

	var publicPosts []*model.Post

	for _,post := range posts{
		value, err := p.UserClient.IsProfilePrivate(post.UserInfo.Id)

		if err!=nil{
			log.Println(err)
			return nil,err
		}

		if !value {
			publicPosts=append(publicPosts, post)
		}
	}

	retVal := p.mapPostsForGuestTimelineToResponseGuestTimelinePostDTO(publicPosts)

	return retVal, nil
}

func (p postService) GetPostForUserTimelineByLocation(ctx context.Context, location string, bearer string) ([]*model.PostResponse, error) {
	posts, err := p.PostRepository.GetPostsByLocation(ctx, location)
	if err!=nil{
		return nil,err
	}

	userInfo, err := p.UserClient.GetLoggedUserInfo(bearer)
	if err != nil {
		return nil, err
	}

	var publicPosts []*model.Post
	for _,post := range posts{
		value, err := p.UserClient.IsProfilePrivate(post.UserInfo.Id)
		if err!=nil{
			return nil,err
		}

		if !value {
			publicPosts=append(publicPosts, post)
		}
	}

	retVal := p.mapPostsToResponsePostDTO(bearer, publicPosts, userInfo.Id)

	return retVal, nil}