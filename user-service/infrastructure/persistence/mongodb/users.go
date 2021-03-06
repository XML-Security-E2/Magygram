package mongodb

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"user-service/domain/model"
	"user-service/domain/repository"
	"user-service/logger"
	"user-service/tracer"
)

type userRepository struct {
	Col *mongo.Collection
}

func NewUserRepository(Col *mongo.Collection) repository.UserRepository {
	return &userRepository{Col}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) (*mongo.InsertOneResult, error) {
	span := tracer.StartSpanFromContext(ctx, "UserRepositoryRegisterUser")
	defer span.Finish()
	return r.Col.InsertOne(ctx, user)
}

func (r *userRepository) Update(ctx context.Context, user *model.User) (*mongo.UpdateResult, error) {
	return r.Col.UpdateOne(ctx, bson.M{"_id":  user.Id},bson.D{{"$set", bson.D{{"email" , user.Email},
																{"username" , user.Username},
																{"name" , user.Name},
																{"favouritePosts" , user.FavouritePosts},
																{"highlightsStory" , user.HighlightsStory},
																{"surname" , user.Surname},
																{"website" , user.Website},
																{"bio" , user.Bio},
																{"imageUrl" , user.ImageUrl},
																{"number" , user.Number},
																{"gender" , user.Gender},
																{"birth_date" , user.BirthDate},
																{"liked_posts" , user.LikedPosts},
																{"disliked_posts" , user.DislikedPosts},
																{"commented_posts" , user.CommentedPosts},
																{"blocked_users" , user.BlockedUsers},
																{"notification_settings" , user.NotificationSettings},
																{"privacy_settings" , user.PrivacySettings},
																{"verified_profile" , user.IsVerified},
																{"category" , user.Category},
																{"private_profile", user.IsPrivate},
	}}})
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*model.User, error) {

	var user = model.User{}
	err := r.Col.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		logger.LoggingEntry.WithFields(logrus.Fields{"user_id" : id}).Warn("Invalid user id")
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) SearchForUsersByUsername(ctx context.Context, username string, loggedUserId string) ([]model.User, error) {
	var users []model.User
	log.Println("param: " + username)
	cursor, err := r.Col.Find(ctx, bson.M{"username": bson.M{"$regex": username, "$options": "i"}, "_id" : bson.M{ "$ne": loggedUserId}, "blocked_users": bson.M{"$ne": loggedUserId}})
	if err != nil {
		return nil, err
	} else {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var record model.User
			err := cursor.Decode(&record)
			users = append(users, record)
			if err != nil {
				return nil, err
			}
		}
	}

	return users, nil
}

func (r *userRepository) SearchForInfluencerByUsername(ctx context.Context, username string, loggedUserId string) ([]model.User, error) {
	var users []model.User
	log.Println("param: " + username)
	cursor, err := r.Col.Find(ctx, bson.M{"username": bson.M{"$regex": username, "$options": "i"}, "category": "INFLUENCER" ,"_id" : bson.M{ "$ne": loggedUserId}, "blocked_users": bson.M{"$ne": loggedUserId}})
	if err != nil {
		return nil, err
	} else {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var record model.User
			err := cursor.Decode(&record)
			users = append(users, record)
			if err != nil {
				return nil, err
			}
		}
	}

	return users, nil
}

func (r *userRepository) SearchForUsersByUsernameByGuest(ctx context.Context, username string) ([]model.User, error) {
	var users []model.User
	cursor, err := r.Col.Find(ctx, bson.M{"username": bson.M{"$regex": username, "$options": "i"}, "private_profile" : false})
	if err != nil {
		return nil, err
	} else {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var record model.User
			err := cursor.Decode(&record)
			users = append(users, record)
			if err != nil {
				return nil, err
			}
		}
	}

	return users, nil
}

func (r userRepository) DeleteUser(ctx context.Context, request *model.User) (*mongo.UpdateResult, error) {
	return r.Col.UpdateOne(ctx, bson.M{"_id":  request.Id},bson.D{{"$set", bson.D{
		{"deleted" , true}}}})
}

func (r userRepository) PhysicalDelete(ctx context.Context, userId string) (*mongo.DeleteResult,error){
	return 	r.Col.DeleteOne(ctx, bson.M{"_id":  userId})
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {

	var user = model.User{}
	err := r.Col.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) IsBlocked(ctx context.Context, subjectId string, objectId string) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "UserRepositoryIsBlocked")
	defer span.Finish()

	var user model.User
	err := r.Col.FindOne(ctx, bson.M{"_id": subjectId, "blocked_users": objectId}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, err
		}
	}

	return true, nil
}
