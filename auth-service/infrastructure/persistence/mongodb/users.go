package mongodb

import (
	"auth-service/domain/model"
	"auth-service/domain/repository"
	"auth-service/logger"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
		{"active" , user.Active},
		{"password" , user.Password},}}})
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

func (r *userRepository) GetAllRolesByUserId(ctx context.Context, userId string) ([]model.Role, error) {
	user, err := r.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user.Roles, nil
}

func (r *userRepository) PhysicalDelete(ctx context.Context, userId string) (*mongo.DeleteResult, error) {
	return 	r.Col.DeleteOne(ctx, bson.M{"_id":  userId})
}