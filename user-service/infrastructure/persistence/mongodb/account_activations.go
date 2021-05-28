package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"user-service/domain/model"
	"user-service/domain/repository"
)

type accountActivationRepository struct {
	Col *mongo.Collection
}

func NewAccountActivationRepository(Col *mongo.Collection) repository.AccountActivationRepository {
	return &accountActivationRepository{Col}
}

func (r *accountActivationRepository) Create(ctx context.Context, a *model.AccountActivation) (*mongo.InsertOneResult, error) {
	return r.Col.InsertOne(ctx, a)
}

func (r *accountActivationRepository) GetById(ctx context.Context, id string) (*model.AccountActivation, error) {

	var accountActivation = model.AccountActivation{}
	err := r.Col.FindOne(ctx, bson.M{"_id": id}).Decode(&accountActivation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}
	return &accountActivation, nil
}

func (r *accountActivationRepository) Update(ctx context.Context, a *model.AccountActivation)  (*mongo.UpdateResult, error) {
	return r.Col.UpdateOne(ctx, bson.M{"_id":  a.Id}, bson.D{{"$set", bson.D{{"used" , a.Used}}}})
}