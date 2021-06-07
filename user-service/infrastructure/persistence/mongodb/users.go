package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"user-service/domain/model"
	"user-service/domain/repository"
)

type userRepository struct {
	Col *mongo.Collection
}

func NewUserRepository(Col *mongo.Collection) repository.UserRepository {
	return &userRepository{Col}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) (*mongo.InsertOneResult, error) {
	return r.Col.InsertOne(ctx, user)
}

func (r *userRepository) Update(ctx context.Context, user *model.User) (*mongo.UpdateResult, error) {
	return r.Col.UpdateOne(ctx, bson.M{"_id":  user.Id},bson.D{{"$set", bson.D{{"email" , user.Email},
																{"username" , user.Username},
																{"name" , user.Name},
																{"favouritePosts" , user.FavouritePosts},
																{"highlightsStory" , user.HighlightsStory},
																{"surname" , user.Surname}}}})
}

func (r *userRepository) UpdateUserDetails(ctx context.Context, user *model.User) (*mongo.UpdateResult, error) {
	return r.Col.UpdateOne(ctx, bson.M{"_id":  user.Id},bson.D{{"$set", bson.D{{"email" , user.Email},
		{"username" , user.Username},
		{"name" , user.Name},
		{"surname" , user.Surname},
		{"website" , user.Website},
		{"bio" , user.Bio},
		{"number" , user.Number},
		{"gender" , user.Gender},
	}}})
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*model.User, error) {

	var user = model.User{}
	err := r.Col.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
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
	cursor, err := r.Col.Find(ctx, bson.M{"username": bson.M{"$regex": username, "$options": "i"}, "_id" : bson.M{ "$ne": loggedUserId}})
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
