package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"user-service/domain/model"
	"user-service/domain/repository"
)

type resetPasswordRepository struct {
	Col *mongo.Collection}

func NewResetPasswordRepository(Col *mongo.Collection) repository.ResetPasswordRepository {
	return &resetPasswordRepository{Col}
}

func (r *resetPasswordRepository) Create(ctx context.Context, a *model.ResetPassword) (*mongo.InsertOneResult, error) {
	return r.Col.InsertOne(ctx, a)
}

func (r *resetPasswordRepository) GetById(ctx context.Context, id string) (*model.ResetPassword, error) {

	var resetPassword = model.ResetPassword{}
	err := r.Col.FindOne(ctx, bson.M{"_id": id}).Decode(&resetPassword)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}
	return &resetPassword, nil
}

func (r *resetPasswordRepository) Update(ctx context.Context, a *model.ResetPassword) (*mongo.UpdateResult, error) {
	return r.Col.UpdateOne(ctx, bson.M{"_id":  a.Id}, bson.D{{"$set", bson.D{{"used" , a.Used}}}})
}
