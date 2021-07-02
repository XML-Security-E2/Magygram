package scheduler

import (
	"ads-service/domain/repository"
	"context"
	"fmt"
	"time"
)

type SchedulerService interface {
	UpdateAllPendingCampaigns(ticker *time.Ticker)
}

type schedulerService struct {
	repository.CampaignRepository
	repository.CampaignUpdateRequestsRepository
}

func NewSchedulerService(r repository.CampaignRepository, re repository.CampaignUpdateRequestsRepository) SchedulerService {
	return &schedulerService{r,re}
}

func (s schedulerService) UpdateAllPendingCampaigns(ticker *time.Ticker) {
	//	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				pending, _ := s.CampaignUpdateRequestsRepository.GetAllPending(context.TODO())
				fmt.Println(len(pending))

				for _, request := range pending {


					if campaign, err := s.CampaignRepository.GetByID(context.TODO(), request.CampaignId); err == nil {
						campaign.DateFrom = request.DateFrom
						campaign.DateTo = request.DateTo
						campaign.TargetGroup = request.TargetGroup
						campaign.MinDisplaysForRepeatedly = request.MinDisplaysForRepeatedly

						if _, err = s.CampaignRepository.Update(context.TODO(), campaign); err == nil {

							request.CampaignUpdateStatus = "DONE"
							s.CampaignUpdateRequestsRepository.Update(context.TODO(), request)
						}

					}
				}
			}
		}
	}()



}