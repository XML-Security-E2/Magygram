package service

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/beevik/guid"
	"io"
	"io/ioutil"
	"magyAgent/conf"
	"magyAgent/domain/model"
	"magyAgent/domain/repository"
	"magyAgent/domain/service-contracts"
	"magyAgent/service/intercomm"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type productService struct {
	repository.ProductRepository
	intercomm.MagygramClient
}

func NewProductService(r repository.ProductRepository, mc intercomm.MagygramClient) service_contracts.ProductService {
	return &productService{r, mc }
}

var (
	FileDirectory = "files"
	FileRequestPrefix = "/api/media/"
)

func (p productService) GetProductCampaignStatistics(ctx context.Context) ([]*model.CampaignStatisticResponse, error) {
	stats, err :=  p.MagygramClient.GetCampaignStatistics()
	list := model.CampaignStatisticReport{
		XMLName:   xml.Name{},
		FileId:    guid.New().String(),
		Campaigns: nil,
		DateCreating: time.Now(),
	}

	var campaigns []model.CampaignStatisticInfo
	for _, stat := range stats {
		media := stat.Media
		media.Url = fmt.Sprintf("%s%s:%s", conf.Current.Magygram.Protocol, conf.Current.Magygram.Domain, conf.Current.Magygram.Port) + media.Url
		stat.Media = media

		campaigns = append(campaigns, model.CampaignStatisticInfo{
			ExposeOnceDate:           stat.ExposeOnceDate,
			MinDisplaysForRepeatedly: stat.MinDisplaysForRepeatedly,
			Type:                     stat.Type,
			Frequency:                stat.Frequency,
			UserViews:                stat.UserViews,
			WebsiteClicks:            stat.WebsiteClicks,
			TargetGroup:              stat.TargetGroup,
			DateFrom:                 stat.DateFrom,
			DateTo:                   stat.DateTo,
			DisplayTime:              stat.DisplayTime,
			CampaignStatus:           stat.CampaignStatus,
			InfluencerUsername:       stat.InfluencerUsername,
			Media:                    stat.Media,
			Website:                  stat.Website,
			Likes:                    stat.Likes,
			Dislikes:                 stat.Dislikes,
			Comments:                 stat.Comments,
			StoryViews:               stat.StoryViews,
			DailyAverage:             stat.DailyAverage,
			Activity:                 stat.Activity,
		})
	}

	list.Campaigns = campaigns
	file, _ := xml.MarshalIndent(list, "", "	")

	_ = ioutil.WriteFile("./probica.xml", file, 0644)

	return stats, err
}

func (p productService) CreateProductCampaign(ctx context.Context, campaignReq *model.CampaignRequest) error {

	product, err := p.ProductRepository.GetById(ctx, campaignReq.ProductId)
	if err != nil {
		return err
	}

	path := strings.Replace(product.ImageURL, FileRequestPrefix, "./" + FileDirectory + "/", -1)


	return p.MagygramClient.CreateCampaign(&model.CampaignApiRequest{
		MinDisplaysForRepeatedly: campaignReq.MinDisplaysForRepeatedly,
		Frequency:                campaignReq.Frequency,
		TargetGroup:              campaignReq.TargetGroup,
		DisplayTime:              campaignReq.DisplayTime,
		DateFrom:                 campaignReq.DateFrom,
		DateTo:                   campaignReq.DateTo,
		ExposeOnceDate:           campaignReq.ExposeOnceDate,
		Type:                     campaignReq.Type,
		FilePath:                 path,
	})
}

func (p productService) CreateProduct(ctx context.Context, productReq *model.ProductRequest) (*model.Product, error) {
	fileName , err := saveFile(productReq.Image)
	if err != nil {
		return nil, err
	}

	product := model.NewProduct(productReq, fileName)
	if err != nil {
		return nil, err
	}

	return p.ProductRepository.Create(ctx, product)
}

func saveFile(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	filename := guid.New().String() + filepath.Ext(file.Filename)
	dst, err := os.Create(filepath.Join(FileDirectory, filename))
	if err != nil {
		return "", err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return FileRequestPrefix + filename, nil
}

func (p productService) UpdateProduct(ctx context.Context, productId string, productReq *model.ProductUpdateRequest) (*model.Product, error) {
	product, err := p.ProductRepository.GetById(ctx, productId)
	if err != nil {
		return nil, err
	}
	product.Name = productReq.Name
	product.Price = productReq.Price
	product.Quantity = productReq.Quantity

	return p.ProductRepository.Update(ctx, product)
}

func (p productService) UpdateProductImage(ctx context.Context, productId string, image *multipart.FileHeader) (*model.Product, error) {
	product, err := p.ProductRepository.GetById(ctx, productId)
	if err != nil {
		return nil, err
	}

	fileName , err := saveFile(image)
	if err != nil {
		return nil, err
	}

	product.ImageURL = fileName
	return p.ProductRepository.Update(ctx, product)
}

func (p productService) GetProductById(ctx context.Context, id string) (*model.Product, error) {
	return p.ProductRepository.GetById(ctx, id)
}

func (p productService) DeleteProductById(ctx context.Context, id string) error {
	return p.ProductRepository.DeleteById(ctx, id)
}

func (p productService) GetAllProducts(ctx context.Context) (*[]model.Product, error) {
	return p.ProductRepository.GetAll(ctx)
}

