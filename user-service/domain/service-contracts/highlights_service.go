package service_contracts

import (
	"context"
	"user-service/domain/model"
)

type HighlightsService interface {
	CreateHighlights(ctx context.Context, bearer string, highlights *model.HighlightRequest) (*model.HighlightProfileResponse,error)
	GetProfileHighlights(ctx context.Context, bearer string, userId string) ([]*model.HighlightProfileResponse, error)
	GetProfileHighlightsByHighlightName(ctx context.Context, bearer string, name string) (*model.HighlightImageWithMedia, error)
}