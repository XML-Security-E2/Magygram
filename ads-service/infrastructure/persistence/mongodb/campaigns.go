package mongodb

import (
	"ads-service/domain/model"
	"ads-service/domain/repository"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type campaignRepository struct {
	Col *mongo.Collection
}

func NewCampaignRepository(Col *mongo.Collection) repository.CampaignRepository {
	return &campaignRepository{Col}
}

func (c campaignRepository) GetAllByOwnerID(ctx context.Context, id string, campaignType string) ([]*model.Campaign, error) {
	cursor, err := c.Col.Find(ctx, bson.M{"owner_id": id, "campaign_type": campaignType})
	var results []*model.Campaign

	if err != nil {
		if cursor != nil {
			defer cursor.Close(ctx)
		}
		return nil, err
	} else {
		if cursor != nil {
			for cursor.Next(ctx) {
				var result model.Campaign

				_ = cursor.Decode(&result)
				results = append(results, &result)
			}
		} else {
			return nil, err
		}
	}
	return results, nil
}

func (c campaignRepository) Create(ctx context.Context, campaign *model.Campaign) (*mongo.InsertOneResult, error) {
	return c.Col.InsertOne(ctx, campaign)
}

func (c campaignRepository) DeleteByID(ctx context.Context, id string) error {
	_, err := c.Col.UpdateOne(ctx, bson.M{"_id":  id},bson.D{{"$set", bson.D{
		{"deleted" , true}}}})
	return err
}

func (c campaignRepository) GetFutureByContentIDAndType(ctx context.Context, contentId string, campaignType string) (*model.Campaign, error) {
	var campaign = model.Campaign{}
	err := c.Col.FindOne(ctx, bson.M{"content_id": contentId,
		"campaign_type": campaignType,
		"deleted": false,
		"date_to" : bson.M{"$gte" : primitive.NewDateTimeFromTime(time.Now())}}).Decode(&campaign)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}
	return &campaign, nil
}

func (c campaignRepository) GetByContentIDAndType(ctx context.Context, contentId string, campaignType string) (*model.Campaign, error) {
	var campaign = model.Campaign{}
	err := c.Col.FindOne(ctx, bson.M{"content_id": contentId,
		"deleted": false,
		"campaign_type": campaignType}).Decode(&campaign)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}
	return &campaign, nil
}

func (c campaignRepository) GetAllFutureByOwnerIDAndType(ctx context.Context, ownerId string, campaignType string) ([]*model.Campaign, error) {

	cursor, err := c.Col.Find(ctx, bson.M{"owner_id": ownerId,
										  "deleted": false,
										  "campaign_type": campaignType,
										  "date_to" : bson.M{"$gte" : primitive.NewDateTimeFromTime(time.Now())}})

	var results []*model.Campaign

	if err != nil {
		if cursor != nil {
			defer cursor.Close(ctx)
		}
		return []*model.Campaign{}, err
	} else {
		if cursor != nil {

			for cursor.Next(ctx) {
				var result model.Campaign

				err := cursor.Decode(&result)
				results = append(results, &result)

				if err != nil {
					return []*model.Campaign{}, err
				}
			}
		}else {
			return []*model.Campaign{}, err
		}
	}
	return results, nil
}


func (c campaignRepository) GetAll(ctx context.Context) ([]*model.Campaign, error) {
	cursor, err := c.Col.Find(ctx, bson.D{})
	var results []*model.Campaign

	if err != nil {
		if cursor != nil {
			defer cursor.Close(ctx)
		}
		return nil, err
	} else {
		for cursor.Next(ctx) {
			var result model.Campaign

			err := cursor.Decode(&result)
			results = append(results, &result)

			if err != nil {
				return nil, err
			}
		}
	}
	return results, nil
}

func (c campaignRepository) Update(ctx context.Context, campaign *model.Campaign) (*mongo.UpdateResult, error) {
	return c.Col.UpdateOne(ctx, bson.M{"_id":  campaign.Id},bson.D{{"$set", bson.D{
		{"min_displays_for_repeatedly" , campaign.MinDisplaysForRepeatedly},
		{"seen_by" , campaign.SeenBy},
		{"daily_seen_by", campaign.DailySeenBy},
		{"target_group" , campaign.TargetGroup},
		{"website_click_count", campaign.WebsiteClickCount},
		{"date_from" , campaign.DateFrom},
		{"date_to" , campaign.DateTo}}}})}

func (c campaignRepository) GetByID(ctx context.Context, id string) (*model.Campaign, error) {
	var campaign = model.Campaign{}
	err := c.Col.FindOne(ctx, bson.M{"_id": id}).Decode(&campaign)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}
	return &campaign, nil
}

func (c campaignRepository) GetUnseenContentIdsCampaignsForUser(ctx context.Context, targetGroup *model.UserTargetGroup, contentType string, count int) ([]*model.Campaign, error) {
	y,m,d := time.Now().Date()
	timee := time.Date(y,m,d,0,0,0,0, time.UTC)


	cursor, err := c.Col.Find(ctx,bson.M{"seen_by": bson.M{"$ne": targetGroup.Id},
										 "campaign_type": contentType,
										 "target_group.min_age": bson.M{"$lte" : targetGroup.Age},
										 "target_group.max_age": bson.M{"$gte" : targetGroup.Age},
										 "$and" : []interface{}{
											 bson.M{"$or" : []interface{}{ bson.M{"target_group.gender": "ANY"}, bson.M{"target_group.gender": targetGroup.Gender}}},
											 bson.M{"$or" : []interface{}{ bson.M{"frequency": "ONCE", "expose_once_date" : bson.M{"$gte": primitive.NewDateTimeFromTime(timee),
													 "$lt": primitive.NewDateTimeFromTime(timee.AddDate(0,0,1))},
													 "$or" : []interface{}{bson.M{"display_time": time.Now().Hour() + 1}, bson.M{"display_time": time.Now().Hour()}}},
													 bson.M{"frequency": "REPEATEDLY", "date_from" : bson.M{"$lte": primitive.NewDateTimeFromTime(time.Now())},
														 "date_to" : bson.M{"$gte": primitive.NewDateTimeFromTime(time.Now())}}}},
										 }})

	var results []*model.Campaign
	var resultsTmp []*model.Campaign

	if err != nil {
		if cursor != nil {
			defer cursor.Close(ctx)
		}
		return nil, err
	} else {
		if cursor != nil {
			for cursor.Next(ctx) {
			var result model.Campaign

			fmt.Println(result.Id)
			err := cursor.Decode(&result)
			results = append(results, &result)

			if err != nil {
				return nil, err
			}

			}
		} else {
			return nil, err
		}
	}
	fmt.Println(len(results))
	if len(results) > 0 {
		for _, res := range results {
			for _, dates := range res.DailySeenBy {
				if dates.Date == timee && len(dates.SeenBy) < res.MinDisplaysForRepeatedly {
					resultsTmp = append(resultsTmp, res)
					break
				}
			}
		}
	}

	if len(resultsTmp) < count {
		return results, nil
	} else {
		return resultsTmp, nil
	}
}