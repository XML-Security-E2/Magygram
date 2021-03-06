package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	_ "net/http"
	"post-service/domain/model"
	"post-service/domain/repository"
	service_contracts "post-service/domain/service-contracts"
	"post-service/domain/service-contracts/exceptions"
	"post-service/logger"
	"post-service/service/intercomm"
	"post-service/tracer"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/beevik/guid"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)

type postService struct {
	repository.PostRepository
	intercomm.MediaClient
	intercomm.UserClient
	intercomm.RelationshipClient
	intercomm.AuthClient
	intercomm.MessageClient
	intercomm.AdsClient
}

func NewPostService(r repository.PostRepository, ic intercomm.MediaClient, uc intercomm.UserClient, ir intercomm.RelationshipClient, ac intercomm.AuthClient, mc intercomm.MessageClient, adsc intercomm.AdsClient) service_contracts.PostService {
	return &postService{r, ic, uc, ir, ac, mc, adsc}
}

func (p postService) GetPostsMediaAndWebsiteByIds(ctx context.Context, ids *model.FollowedUsersResponse) ([]*model.IdMediaWebsiteResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceGetPostsMediaAndWebsiteByIds")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	var resp []*model.IdMediaWebsiteResponse

	for _, id := range ids.Users {
		post, err := p.PostRepository.GetByID(ctx, id)
		if err == nil {
			resp = append(resp, &model.IdMediaWebsiteResponse{
				Id:         post.Id,
				Media:      post.Media[0],
				Website:    post.WebSite,
				Likes:      len(post.LikedBy),
				Dislikes:   len(post.DislikedBy),
				Comments:   len(post.Comments),
				StoryViews: 0,
			})
		}
	}

	if resp == nil {
		return []*model.IdMediaWebsiteResponse{}, nil
	}

	return resp, nil
}

func (p postService) CreatePost(ctx context.Context, bearer string, postRequest *model.PostRequest) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceCreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)
	if err != nil {
		return "", err
	}

	media, err := p.MediaClient.SaveMedia(ctx, postRequest.Media)
	if err != nil {
		return "", err
	}

	post, err := model.NewPost(postRequest, *userInfo, "REGULAR", media, "")
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"tags": postRequest.Tags,
			"description": postRequest.Description,
			"location":    postRequest.Location}).Warn("Post creating validation failure")

		return "", err
	}

	if err = validator.New().Struct(post); err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"tags": postRequest.Tags,
			"description": postRequest.Description,
			"location":    postRequest.Location}).Warn("Post creating validation failure")
		return "", err
	}

	result, err := p.PostRepository.Create(ctx, post)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"tags": postRequest.Tags,
			"description": postRequest.Description,
			"location":    postRequest.Location}).Error("Post database create failure")
		return "", err
	}

	err = p.MessageClient.CreateNotifications(ctx, &intercomm.NotificationRequest{
		Username:   userInfo.Username,
		UserId:     userInfo.Id,
		UserFromId: userInfo.Id,
		NotifyUrl:  "TODO",
		ImageUrl:   post.UserInfo.ImageURL,
		Type:       intercomm.PublishedPost,
	})

	if err != nil {
		return "", err
	}

	if postId, ok := result.InsertedID.(string); ok {
		logger.LoggingEntry.WithFields(logrus.Fields{"post_id": post.Id, "user_id": userInfo.Id}).Info("Post created")
		return postId, nil
	}

	return "", err
}

func (p postService) CreatePostCampaignFromApi(ctx context.Context, bearer string, postRequest *model.PostRequest) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceCreatePostCampaignFromApi")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userInfo, err := p.UserClient.GetLoggedAgentInfo(ctx, bearer)
	if err != nil { return "", err}

	media, err := p.MediaClient.SaveMedia(ctx, postRequest.Media)
	if err != nil { return "", err}

	postRequest.Description = "You can check our products at our website: " + userInfo.Website

	post, err := model.NewPost(postRequest, model.UserInfo{
		Id:       userInfo.Id,
		Username: userInfo.Username,
		ImageURL: userInfo.ImageURL,
	}, "CAMPAIGN", media, userInfo.Website)
	if err != nil {return "", err}

	if err = validator.New().Struct(post); err!= nil {
		return "", err
	}

	_, err = p.PostRepository.Create(ctx, post)
	if err != nil {
		return "", err
	}

	err = p.MessageClient.CreateNotifications(ctx, &intercomm.NotificationRequest{
		Username:  userInfo.Username,
		UserId:    userInfo.Id,
		UserFromId:userInfo.Id,
		NotifyUrl: "TODO",
		ImageUrl:  post.UserInfo.ImageURL,
		Type:      intercomm.PublishedPost,
	})

	if err != nil {
		return "", err
	}

	return post.Id, nil
}

func (p postService) CreatePostInfluencer(ctx context.Context, bearer string, request  *model.InfluencerRequest) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceCreatePostInfluencer")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)
	if err != nil { return "", err}

	agentPost, err := p.PostRepository.GetByID(ctx, request.PostIdInfluencer)
	if err != nil {
		return "", err
	}

	agentInfo, err := p.UserClient.GetLoggedAgentInfoById(agentPost.UserInfo.Id)
	if err != nil { return "", err}

	post, err := model.NewPostInfluencer(agentPost, *userInfo, agentInfo.Website)

	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"tags": post.Tags,
			"description" : post.Description,
			"location" : post.Location}).Warn("Post creating validation failure")

		return "", err}

	if err = validator.New().Struct(post); err!= nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"tags": agentPost.Tags,
			"description" : agentPost.Description,
			"location" : agentPost.Location}).Warn("Post creating validation failure")
		return "", err
	}

	result, err := p.PostRepository.Create(ctx, post)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"tags": agentPost.Tags,
			"description" : agentPost.Description,
			"location" : agentPost.Location}).Error("Post database create failure")
		return "", err}

	err = p.MessageClient.CreateNotifications(ctx, &intercomm.NotificationRequest{
		Username:  userInfo.Username,
		UserId:    userInfo.Id,
		UserFromId:userInfo.Id,
		NotifyUrl: "TODO",
		ImageUrl:  post.UserInfo.ImageURL,
		Type:      intercomm.PublishedPost,
	})

	if err != nil {
		return "", err
	}

	if postId, ok := result.InsertedID.(string); ok {
		logger.LoggingEntry.WithFields(logrus.Fields{"post_id": post.Id, "user_id" : userInfo.Id}).Info("Post created")
		return postId, nil
	}

	return post.Id, err
}
func (p postService) CreatePostCampaign(ctx context.Context, bearer string, postRequest *model.PostRequest, campaignReq *model.CampaignRequest) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceCreatePostCampaign")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userInfo, err := p.UserClient.GetLoggedAgentInfo(ctx, bearer)
	if err != nil {
		return "", err
	}

	media, err := p.MediaClient.SaveMedia(ctx, postRequest.Media)
	if err != nil {
		return "", err
	}

	post, err := model.NewPost(postRequest, model.UserInfo{
		Id:       userInfo.Id,
		Username: userInfo.Username,
		ImageURL: userInfo.ImageURL,
	}, "CAMPAIGN", media, userInfo.Website)
	if err != nil {
		return "", err
	}

	if err = validator.New().Struct(post); err != nil {
		return "", err
	}

	campaignReq.ContentId = post.Id
	err = p.AdsClient.CreatePostCampaign(ctx, bearer, campaignReq)
	if err != nil {
		return "", err
	}

	result, err := p.PostRepository.Create(ctx, post)
	if err != nil {
		return "", err
	}

	err = p.MessageClient.CreateNotifications(ctx, &intercomm.NotificationRequest{
		Username:   userInfo.Username,
		UserId:     userInfo.Id,
		UserFromId: userInfo.Id,
		NotifyUrl:  "TODO",
		ImageUrl:   post.UserInfo.ImageURL,
		Type:       intercomm.PublishedPost,
	})

	if err != nil {
		return "", err
	}

	if postId, ok := result.InsertedID.(string); ok {
		return postId, nil
	}

	return "", err
}

type timeSlice []*model.Post

func (p timeSlice) Len() int {
	return len(p)
}

func (p timeSlice) Less(i, j int) bool {
	return p[i].CreatedTime.After(p[j].CreatedTime)
}

func (p timeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p postService) GetPostsForTimeline(ctx context.Context, bearer string) ([]*model.PostResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceGetPostsForTimeline")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	var posts []*model.Post
	userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)
	if err != nil {
		return nil, err
	}

	var followedUsers model.FollowedUsersResponse
	followedUsers, err = p.RelationshipClient.GetUnmutedFollowedUsers(ctx, userInfo.Id)
	if err != nil {
		return nil, err
	}

	for _, userId := range followedUsers.Users {
		var newPosts []*model.Post
		newPosts, _ = p.PostRepository.GetPostsForUser(ctx, userId)
		posts = append(posts, newPosts...)
	}

	for _, post := range posts {
		if post.ContentType == "CAMPAIGN" {
			p.AdsClient.UpdateCampaignVisitor(ctx, bearer, post.Id)
		}
	}

	sortedPosts := sortPostPerTime(ctx, posts)

	campaignCount := int64(len(sortedPosts) / 5) // broj postova sortiranih, na svaki 5i

	campaignPostIds, err := p.AdsClient.GetPostCampaignSuggestion(ctx, bearer, strconv.Itoa(int(campaignCount)))

	log.Println("TEST")
	log.Println(campaignCount)
	log.Println(len(campaignPostIds))

	var retPosts []*model.Post
	var index int = 0

	if len(campaignPostIds) != 0 {
		var postNumberBeforeCampaign int32
		if len(campaignPostIds) == 1 {
			postNumberBeforeCampaign = int32(len(sortedPosts)/2) + 1
		} else {
			postNumberBeforeCampaign = int32(len(sortedPosts) / (len(campaignPostIds)))
		}

		for sortedIndex, post := range sortedPosts {
			if (sortedIndex+1)%int(postNumberBeforeCampaign) == 0 {
				value, err := p.PostRepository.GetByID(ctx, campaignPostIds[index])
				if err != nil {
					return nil, errors.New("Campaign not exist")
				}
				retPosts = append(retPosts, value)
				index++
			}

			retPosts = append(retPosts, post)
		}

		retVal := p.mapPostsToResponsePostDTO(ctx, bearer, retPosts, userInfo.Id)

		return retVal, nil

	} else {
		retVal := p.mapPostsToResponsePostDTO(ctx, bearer, sortedPosts, userInfo.Id)

		return retVal, nil
	}
}

func sortPostPerTime(ctx context.Context, posts []*model.Post) []*model.Post {
	span := tracer.StartSpanFromContext(ctx, "PostServiceSortPostPerTime")
	defer span.Finish()

	dateSortedPosts := make(timeSlice, 0, len(posts))
	for _, post := range posts {
		dateSortedPosts = append(dateSortedPosts, post)
	}

	sort.Sort(dateSortedPosts)

	return dateSortedPosts
}

func (p postService) DeletePost(ctx context.Context, bearer string, requestId string) error {
	span := tracer.StartSpanFromContext(ctx, "PostServiceDeletePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	retVal, err := p.AuthClient.HasRole(ctx, bearer, "delete_posts")
	if err != nil {
		return errors.New("auth service not found")
	}

	request, err := p.PostRepository.GetByID(ctx, requestId)
	if err != nil {
		return errors.New("post not found")
	}

	if !retVal {
		userId, err := p.AuthClient.GetLoggedUserId(ctx, bearer)
		if err != nil {
			return err
		}
		if request.UserInfo.Id != userId {
			return errors.New("user not authorized for post delete")
		}
	}

	err = p.AdsClient.DeleteCampaign(ctx, bearer, request.Id)
	if err != nil {
		return err
	}

	request.IsDeleted = true

	_, err = p.PostRepository.DeletePost(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (p postService) LikePost(ctx context.Context, bearer string, postId string) error {
	span := tracer.StartSpanFromContext(ctx, "PostServiceLikePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)
	if err != nil {
		return err
	}

	result, err := p.PostRepository.GetByID(ctx, postId)
	if err != nil {
		return err
	}
	var res model.UserInfo

	res.Id = userInfo.Id
	res.ImageURL = userInfo.ImageURL
	res.Username = userInfo.Username

	result.LikedBy = append(result.LikedBy, res)

	err = p.UserClient.UpdateLikedPosts(ctx, bearer, postId)
	if err != nil {
		return err
	}

	_, err = p.PostRepository.Update(ctx, result)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userInfo.Id,
			"post_id": postId}).Error("Post like, database update failure")
		return err
	}

	err = p.MessageClient.CreateNotification(ctx, &intercomm.NotificationRequest{
		Username:   userInfo.Username,
		UserId:     result.UserInfo.Id,
		UserFromId: userInfo.Id,
		NotifyUrl:  "TODO",
		ImageUrl:   userInfo.ImageURL,
		Type:       intercomm.Liked,
	})

	if err != nil {
		return err
	}

	logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userInfo.Id,
		"post_id": postId}).Info("Post liked")
	return nil
}

func (p postService) UnlikePost(ctx context.Context, bearer string, postId string) error {
	span := tracer.StartSpanFromContext(ctx, "PostServiceUnlikePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)
	if err != nil {
		return err
	}

	result, err := p.PostRepository.GetByID(ctx, postId)
	if err != nil {
		return err
	}

	result.LikedBy = findAndDeleteLikedBy(result, userInfo)

	err = p.UserClient.UpdateLikedPosts(ctx, bearer, postId)
	if err != nil {
		return err
	}

	_, err = p.PostRepository.Update(ctx, result)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userInfo.Id,
			"post_id": postId}).Error("Post dislike, database update failure")
		return err
	}

	logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userInfo.Id,
		"post_id": postId}).Info("Post unliked")
	return nil
}

func (p postService) DislikePost(ctx context.Context, bearer string, postId string) error {
	span := tracer.StartSpanFromContext(ctx, "PostServiceDislikePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)
	if err != nil {
		return err
	}

	result, err := p.PostRepository.GetByID(ctx, postId)
	if err != nil {
		return err
	}
	var res model.UserInfo

	res.Id = userInfo.Id
	res.ImageURL = userInfo.ImageURL
	res.Username = userInfo.Username

	result.DislikedBy = append(result.DislikedBy, res)

	err = p.UserClient.UpdateDislikedPosts(ctx, bearer, postId)
	if err != nil {
		return err
	}

	_, err = p.PostRepository.Update(ctx, result)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userInfo.Id,
			"post_id": postId}).Error("Post dislike, database update failure")
		return err
	}

	err = p.MessageClient.CreateNotification(ctx, &intercomm.NotificationRequest{
		Username:   userInfo.Username,
		UserId:     result.UserInfo.Id,
		UserFromId: userInfo.Id,
		NotifyUrl:  "TODO",
		ImageUrl:   userInfo.ImageURL,
		Type:       intercomm.Disliked,
	})

	if err != nil {
		return err
	}

	logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userInfo.Id,
		"post_id": postId}).Info("Post disliked")
	return nil
}

func (p postService) UndislikePost(ctx context.Context, bearer string, postId string) error {
	span := tracer.StartSpanFromContext(ctx, "PostServiceUndislikePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)
	if err != nil {
		return err
	}
	fmt.Println(postId)

	result, err := p.PostRepository.GetByID(ctx, postId)
	if err != nil {
		return err
	}

	result.DislikedBy = findAndDeleteDislikedBy(result, userInfo)

	err = p.UserClient.UpdateDislikedPosts(ctx, bearer, postId)
	if err != nil {
		return err
	}

	_, err = p.PostRepository.Update(ctx, result)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userInfo.Id,
			"post_id": postId}).Error("Post un-dislike, database update failure")
		return err
	}

	logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userInfo.Id,
		"post_id": postId}).Info("Post un-disliked")
	return nil
}

func (p postService) AddComment(ctx context.Context, postId string, content string, bearer string, tags []model.Tag) (*model.Comment, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceAddComment")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)
	if err != nil {
		return nil, err
	}

	result, err := p.PostRepository.GetByID(ctx, postId)
	if err != nil {
		return nil, err
	}

	var res model.Comment

	res.Id = guid.New().String()
	res.Content = content
	res.CreatedBy = *userInfo
	res.TimeCreated = time.Now()
	res.Tags = tags

	result.Comments = append(result.Comments, res)

	err = p.UserClient.AddComment(ctx, bearer, postId)
	if err != nil {
		return nil, err
	}

	_, err = p.PostRepository.Update(ctx, result)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userInfo.Id,
			"post_id": postId}).Error("Post comment, database update failure")
		return nil, err
	}

	err = p.MessageClient.CreateNotification(ctx, &intercomm.NotificationRequest{
		Username:   userInfo.Username,
		UserId:     result.UserInfo.Id,
		UserFromId: userInfo.Id,
		NotifyUrl:  "TODO",
		ImageUrl:   userInfo.ImageURL,
		Type:       intercomm.Commented,
	})

	if err != nil {
		return nil, err
	}

	logger.LoggingEntry.WithFields(logrus.Fields{"user_id": userInfo.Id,
		"post_id": postId}).Info("Post commented")

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

func (p postService) mapPostsToResponsePostDTO(ctx context.Context, bearer string, result []*model.Post, userId string) []*model.PostResponse {
	var retVal []*model.PostResponse

	postIdFavourites, err := p.UserClient.MapPostsToFavourites(ctx, bearer, getIdsFromPosts(result))
	if err != nil {
		return nil
	}

	for _, post := range result {
		res, err := model.NewPostResponse(post, hasUserLikedPost(post, userId), hasUserDislikedPost(post, userId), isInFavourites(post, postIdFavourites))

		if err != nil {
			return nil
		}

		retVal = append(retVal, res)
	}

	return retVal
}

func (p postService) mapPostsToResponsePostDTOForAdmin(bearer string, result []*model.Post) []*model.PostResponse {
	var retVal []*model.PostResponse

	for _, post := range result {
		res, err := model.NewPostResponse(post, false, false, false)

		if err != nil {
			return nil
		}

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
		if likedUserInfo.Id == usedId {
			retVal = true
			break
		}
	}

	return retVal
}

func hasUserDislikedPost(post *model.Post, usedId string) bool {
	var retVal = false

	for _, dislikedUserInfo := range post.DislikedBy {
		if dislikedUserInfo.Id == usedId {
			retVal = true
			break
		}
	}

	return retVal
}

func (p postService) GetPostsFirstImage(ctx context.Context, postId string) (*model.Media, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceGetPostsFirstImage")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

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
	span := tracer.StartSpanFromContext(ctx, "PostServiceEditPost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	post, err := p.PostRepository.GetByID(ctx, postRequest.Id)
	if err != nil {
		return errors.New("invalid post id")
	}

	isOwner, err := p.CheckIfUsersPostFromBearer(ctx, bearer, post.UserInfo.Id)
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
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id": post.UserInfo.Id,
			"post_id":     post.Id,
			"tags":        postRequest.Tags,
			"description": postRequest.Description,
			"location":    postRequest.Location}).Error("Post edit, database update failure")
		return err
	}

	logger.LoggingEntry.WithFields(logrus.Fields{"user_id": post.UserInfo.Id,
		"post_id": post.Id,
	}).Info("Post edited")

	return nil
}

func (p postService) CheckIfUsersPostFromBearer(ctx context.Context, bearer string, postOwnerId string) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceCheckIfUsersPostFromBearer")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"post_owner_id": postOwnerId}).Warn("Unauthorized access")
		return false, err
	}

	if postOwnerId != userInfo.Id {
		logger.LoggingEntry.WithFields(logrus.Fields{"post_owner_id": postOwnerId, "user_id": userInfo.Id}).Warn("Unauthorized access")
		return false, nil
	}
	return true, nil
}

func (p postService) GetUsersPosts(ctx context.Context, bearer string, postOwnerId string) ([]*model.PostProfileResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceGetUsersPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	retVal, err := p.AuthClient.HasRole(ctx, bearer, "visit_private_profiles")
	if err != nil {
		return nil, errors.New("auth service not found")
	}

	if !p.checkIfUserContentIsAccessible(ctx, bearer, postOwnerId) {
		if !retVal {
			return nil, &exceptions.UnauthorizedAccessError{Msg: "User not authorized"}
		}
	}

	userPosts, err := p.PostRepository.GetPostsForUser(ctx, postOwnerId)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"post_owner_id": postOwnerId}).Warn("Error while getting user posts")
		return nil, errors.New("invalid user id")
	}

	var userPostsResponse []*model.PostProfileResponse
	for _, post := range userPosts {
		userPostsResponse = append(userPostsResponse, &model.PostProfileResponse{
			Id:    post.Id,
			Media: post.Media[0],
		})
	}

	return userPostsResponse, nil
}

func (p postService) checkIfUserContentIsAccessible(ctx context.Context, bearer string, postOwnerId string) bool {
	isPrivate, err := p.UserClient.IsUserPrivate(ctx, postOwnerId)
	if err != nil {
		return false
	}

	if isPrivate {
		if bearer == "" {
			logger.LoggingEntry.WithFields(logrus.Fields{"post_owner_id": postOwnerId}).Warn("Unauthorized access")
			return false
		}
		userId, err := p.AuthClient.GetLoggedUserId(ctx, bearer)
		if err != nil {
			logger.LoggingEntry.WithFields(logrus.Fields{"post_owner_id": postOwnerId}).Warn("Unauthorized access")
			return false
		}

		if userId != postOwnerId {
			followedUsers, err := p.RelationshipClient.GetFollowedUsers(ctx, userId)
			if err != nil {
				logger.LoggingEntry.WithFields(logrus.Fields{"post_owner_id": postOwnerId, "user_id": userId}).Warn("Unauthorized access")
				return false
			}

			for _, usrId := range followedUsers.Users {
				if postOwnerId == usrId {
					return true
				}
			}

			logger.LoggingEntry.WithFields(logrus.Fields{"post_owner_id": postOwnerId, "user_id": userId}).Warn("Unauthorized access")
			return false
		}
	}

	return true
}

func (p postService) GetUsersPostsCount(ctx context.Context, postOwnerId string) (int, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceGetUsersPostsCount")
	defer span.Finish()

	userPosts, err := p.PostRepository.GetPostsForUser(ctx, postOwnerId)
	if err != nil {
		return 0, errors.New("invalid user id")
	}

	return len(userPosts), nil
}

func (p postService) GetPostForMessagesById(ctx context.Context, bearer string, postId string) (*model.PostResponse, *model.UserInfo, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceGetPostForMessagesById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userId, err := p.AuthClient.GetLoggedUserId(ctx, bearer)
	if err != nil {
		return nil, nil, err
	}

	post, err := p.PostRepository.GetByID(ctx, postId)
	if err != nil {
		return nil, nil, errors.New("invalid post id")
	}

	retVal, err := p.AuthClient.HasRole(ctx, bearer, "visit_private_profiles")
	if err != nil {
		return nil, nil, errors.New("auth service not found")
	}

	if !retVal {
		if !p.checkIfUserContentIsAccessible(ctx, bearer, post.UserInfo.Id) {
			return nil, &post.UserInfo, &exceptions.UnauthorizedAccessError{Msg: "User not authorized"}
		}

		postIdFavourite, err := p.UserClient.MapPostsToFavourites(ctx, bearer, []string{post.Id})
		if err != nil {
			return nil, nil, err
		}

		res, err := model.NewPostResponse(post, hasUserLikedPost(post, userId), hasUserDislikedPost(post, userId), isInFavourites(post, postIdFavourite))
		if err != nil {
			return nil, nil, err
		}

		return res, nil, nil
	}

	res, err := model.NewPostResponse(post, false, false, false)
	if err != nil {
		return nil, nil, err
	}

	return res, nil, nil
}

func (p postService) GetPostById(ctx context.Context, bearer string, postId string) (*model.PostResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceGetPostById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userId, err := p.AuthClient.GetLoggedUserId(ctx, bearer)
	if err != nil {
		return nil, err
	}

	post, err := p.PostRepository.GetByID(ctx, postId)
	if err != nil {
		return nil, errors.New("invalid post id")
	}

	retVal, err := p.AuthClient.HasRole(ctx, bearer, "visit_private_profiles")
	if err != nil {
		return nil, errors.New("auth service not found")
	}

	if !retVal {
		if !p.checkIfUserContentIsAccessible(ctx, bearer, post.UserInfo.Id) {
			return nil, &exceptions.UnauthorizedAccessError{Msg: "User not authorized"}
		}

		postIdFavourite, err := p.UserClient.MapPostsToFavourites(ctx, bearer, []string{post.Id})
		if err != nil {
			return nil, err
		}

		res, err := model.NewPostResponse(post, hasUserLikedPost(post, userId), hasUserDislikedPost(post, userId), isInFavourites(post, postIdFavourite))
		if err != nil {
			return nil, err
		}

		return res, nil
	}

	res, err := model.NewPostResponse(post, false, false, false)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p postService) SearchForPostsByHashTagByGuest(ctx context.Context, hashTagValue string) ([]*model.HashTageSearchResponseDTO, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceSearchForPostsByHashTagByGuest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

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

	for _, post := range posts {
		for _, hashTag := range post.HashTags {
			if strings.Contains(hashTag, hashTagValue) {
				if _, ok := elementMap[hashTag]; ok {
					elementMap[hashTag] = elementMap[hashTag] + 1
				} else {
					elementMap[hashTag] = 1
				}
			}
		}
	}

	return elementMap
}

func (p postService) GetPostsByHashTagForGuest(ctx context.Context, hashtag string) ([]*model.GuestTimelinePostResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceGetPostsByHashTagForGuest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	posts, err := p.PostRepository.GetPostsByHashTag(ctx, hashtag)

	if err != nil {
		return nil, err
	}

	var publicPosts []*model.Post

	for _, post := range posts {
		value, err := p.UserClient.IsUserPrivate(ctx, post.UserInfo.Id)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		if !value {
			publicPosts = append(publicPosts, post)
		}
	}

	retVal := p.mapPostsForGuestTimelineToResponseGuestTimelinePostDTO(publicPosts)

	return retVal, nil
}

func (p postService) mapPostsForGuestTimelineToResponseGuestTimelinePostDTO(posts []*model.Post) []*model.GuestTimelinePostResponse {
	var retVal []*model.GuestTimelinePostResponse

	for _, post := range posts {
		res, err := model.NewGuestTimelinePostResponse(post)

		if err != nil {
			return nil
		}

		retVal = append(retVal, res)
	}

	return retVal
}

func (p postService) GetPostForUserTimelineByHashTag(ctx context.Context, hashtag string, bearer string) ([]*model.PostResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceGetPostForUserTimelineByHashTag")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	posts, err := p.PostRepository.GetPostsByHashTag(ctx, hashtag)
	if err != nil {
		return nil, err
	}

	var publicPosts []*model.Post

	retValRole, err := p.AuthClient.HasRole(ctx, bearer, "search_all_post_by_hashtag")
	if err != nil {
		return nil, errors.New("auth service not found")
	}

	if retValRole {
		publicPosts = posts
		retVal := p.mapPostsToResponsePostDTOForAdmin(bearer, publicPosts)

		return retVal, nil
	} else {
		userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)
		if err != nil {
			return nil, err
		}

		for _, post := range posts {
			value, err := p.UserClient.IsUserPrivate(ctx, post.UserInfo.Id)
			if err != nil {
				return nil, err
			}

			if !value {
				publicPosts = append(publicPosts, post)
			}
		}
		retVal := p.mapPostsToResponsePostDTO(ctx, bearer, publicPosts, userInfo.Id)

		return retVal, nil
	}
}

func (p postService) SearchPostsByLocation(ctx context.Context, locationValue string) ([]*model.LocationSearchResponseDTO, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceSearchPostsByLocation")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

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

	for _, post := range posts {
		if strings.Contains(post.Location, locationValue) {
			if _, ok := elementMap[post.Location]; ok {
				elementMap[post.Location] = elementMap[post.Location] + 1
			} else {
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
	span := tracer.StartSpanFromContext(ctx, "PostServiceGetPostForGuestTimelineByLocation")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	posts, err := p.PostRepository.GetPostsByLocation(ctx, location)

	if err != nil {
		return nil, err
	}

	var publicPosts []*model.Post

	for _, post := range posts {
		value, err := p.UserClient.IsUserPrivate(ctx, post.UserInfo.Id)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		if !value {
			publicPosts = append(publicPosts, post)
		}
	}

	retVal := p.mapPostsForGuestTimelineToResponseGuestTimelinePostDTO(publicPosts)

	return retVal, nil
}

func (p postService) GetPostForUserTimelineByLocation(ctx context.Context, location string, bearer string) ([]*model.PostResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceGetPostForUserTimelineByLocation")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	posts, err := p.PostRepository.GetPostsByLocation(ctx, location)
	if err != nil {
		return nil, err
	}

	retValRole, err := p.AuthClient.HasRole(ctx, bearer, "search_all_post_by_location")
	if err != nil {
		return nil, errors.New("auth service not found")
	}

	if retValRole {
		retVal := p.mapPostsToResponsePostDTOForAdmin(bearer, posts)

		return retVal, nil
	} else {
		userInfo, err := p.UserClient.GetLoggedUserInfo(ctx, bearer)
		if err != nil {
			return nil, err
		}

		var publicPosts []*model.Post
		for _, post := range posts {
			value, err := p.UserClient.IsUserPrivate(ctx, post.UserInfo.Id)
			if err != nil {
				return nil, err
			}

			if !value {
				publicPosts = append(publicPosts, post)
			}
		}

		retVal := p.mapPostsToResponsePostDTO(ctx, bearer, publicPosts, userInfo.Id)

		return retVal, nil
	}

}

func (p postService) GetPostByIdForGuest(ctx context.Context, postId string) (*model.GuestTimelinePostResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceGetPostByIdForGuest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	post, err := p.PostRepository.GetByID(ctx, postId)
	if err != nil {
		return nil, errors.New("invalid post id")
	}

	var retVal, _ = model.NewGuestTimelinePostResponse(post)

	return retVal, nil
}

func (p postService) GetUserPostCampaigns(ctx context.Context, bearer string) ([]*model.PostProfileResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceGetUserPostCampaigns")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	posts, err := p.AdsClient.GetAllActiveAgentsPostCampaigns(ctx, bearer)
	if err != nil {
		return []*model.PostProfileResponse{}, err
	}

	userPosts, err := p.PostRepository.GetPostsByPostIdArray(ctx, posts)

	if userPosts == nil {
		return nil, nil
	}

	if err != nil {
		return []*model.PostProfileResponse{}, err
	}

	var userPostsResponse []*model.PostProfileResponse
	for _, post := range userPosts {
		fmt.Println(post.Id)
		userPostsResponse = append(userPostsResponse, &model.PostProfileResponse{
			Id:    post.Id,
			Media: post.Media[0],
		})
	}

	return userPostsResponse, nil
}

func (p postService) GetUserLikedPosts(ctx context.Context, bearer string) ([]*model.PostProfileResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceGetUserLikedPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userLikedPostIds, err := p.UserClient.GetLikedPosts(ctx, bearer)
	if err != nil {
		return []*model.PostProfileResponse{}, err
	}

	userPosts, err := p.PostRepository.GetPostsByPostIdArray(ctx, userLikedPostIds)

	var userPostsResponse []*model.PostProfileResponse
	for _, post := range userPosts {
		userPostsResponse = append(userPostsResponse, &model.PostProfileResponse{
			Id:    post.Id,
			Media: post.Media[0],
		})
	}

	return userPostsResponse, nil
}

func (p postService) GetUserDislikedPosts(ctx context.Context, bearer string) ([]*model.PostProfileResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "PostServiceGetUserDislikedPosts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userLikedPostIds, err := p.UserClient.GetDislikedPosts(ctx, bearer)
	if err != nil {
		return []*model.PostProfileResponse{}, err
	}

	userPosts, err := p.PostRepository.GetPostsByPostIdArray(ctx, userLikedPostIds)

	var userPostsResponse []*model.PostProfileResponse
	for _, post := range userPosts {
		userPostsResponse = append(userPostsResponse, &model.PostProfileResponse{
			Id:    post.Id,
			Media: post.Media[0],
		})
	}

	return userPostsResponse, nil
}

func (p postService) EditPostOwnerInfo(ctx context.Context, bearer string, userInfo *model.UserInfo) error {
	span := tracer.StartSpanFromContext(ctx, "PostServiceEditPostOwnerInfo")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userId, err := p.AuthClient.GetLoggedUserId(ctx, bearer)
	if err != nil {
		return err
	}
	fmt.Println(userId)

	if userId != userInfo.Id {
		return errors.New("unauthorized edit")
	}

	userPosts, err := p.PostRepository.GetPostsForUser(ctx, userId)
	if err != nil {
		return errors.New("invalid user id")
	}

	for _, userPost := range userPosts {
		userPost.UserInfo = *userInfo

		_, err = p.PostRepository.Update(ctx, userPost)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p postService) EditLikedByInfo(ctx context.Context, bearer string, userInfoEdit *model.UserInfoEdit) error {
	span := tracer.StartSpanFromContext(ctx, "PostServiceEditLikedByInfo")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userId, err := p.AuthClient.GetLoggedUserId(ctx, bearer)
	if err != nil {
		return err
	}

	for _, postId := range userInfoEdit.PostIds {
		fmt.Println(postId)
		post, err := p.PostRepository.GetByID(ctx, postId)
		if err == nil {
			for idx, likedBy := range post.LikedBy {
				if userId == likedBy.Id {
					post.LikedBy[idx] = model.UserInfo{
						Id:       userId,
						Username: userInfoEdit.Username,
						ImageURL: userInfoEdit.ImageURL,
					}
					break
				}
			}
			_, err = p.PostRepository.Update(ctx, post)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (p postService) EditDislikedByInfo(ctx context.Context, bearer string, userInfoEdit *model.UserInfoEdit) error {
	span := tracer.StartSpanFromContext(ctx, "PostServiceEditDislikedByInfo")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userId, err := p.AuthClient.GetLoggedUserId(ctx, bearer)
	if err != nil {
		return err
	}

	for _, postId := range userInfoEdit.PostIds {
		post, err := p.PostRepository.GetByID(ctx, postId)
		if err == nil {
			for idx, dislikedBy := range post.DislikedBy {
				if userId == dislikedBy.Id {
					post.DislikedBy[idx] = model.UserInfo{
						Id:       userId,
						Username: userInfoEdit.Username,
						ImageURL: userInfoEdit.ImageURL,
					}
					break
				}
			}
			_, err = p.PostRepository.Update(ctx, post)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (p postService) EditCommentedByInfo(ctx context.Context, bearer string, userInfoEdit *model.UserInfoEdit) error {
	span := tracer.StartSpanFromContext(ctx, "PostServiceEditCommentedByInfo")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	userId, err := p.AuthClient.GetLoggedUserId(ctx, bearer)
	if err != nil {
		return err
	}

	for _, postId := range userInfoEdit.PostIds {
		post, err := p.PostRepository.GetByID(ctx, postId)
		if err == nil {
			for idx, comment := range post.Comments {
				if userId == comment.CreatedBy.Id {
					post.Comments[idx].CreatedBy = model.UserInfo{
						Id:       userId,
						Username: userInfoEdit.Username,
						ImageURL: userInfoEdit.ImageURL,
					}
				}
			}
			_, err = p.PostRepository.Update(ctx, post)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
