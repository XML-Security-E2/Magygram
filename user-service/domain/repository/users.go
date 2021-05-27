package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"user-service/domain/model"
)

type UserModel struct {
	C *mongo.Collection
}

func (m *UserModel) Insert(user model.User) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), user)
}