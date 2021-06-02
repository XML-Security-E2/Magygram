package service

import (
	"context"
	"github.com/beevik/guid"
	"io"
	"media-service/domain/model"
	"mime/multipart"
	"os"
	"path/filepath"
)

type MediaService interface {
	SaveMedia(ctx context.Context, files []*multipart.FileHeader) ([]model.Media, error)
	GetMedia(ctx context.Context, mediaId string) (*os.File, error)
}

type mediaService struct {
}

var (
	FileDirectory = "files"
	FileRequestPrefix = "/api/media/"
	IMAGE = "IMAGE"
	VIDEO = "VIDEO"
	ImageExtensions = []string{".jpg", ".jpeg", ".png"}
)

func NewMediaService() MediaService {
	return &mediaService{}
}

func (m mediaService) SaveMedia(ctx context.Context, files []*multipart.FileHeader) ([]model.Media, error) {
	var mediaIds []model.Media
	for _, file := range files {
		media , err := saveFile(file)
		if err != nil {
			return nil, err
		}
		mediaIds = append(mediaIds, *media)
	}
	return mediaIds, nil
}

func saveFile(file *multipart.FileHeader) (*model.Media, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	filename := guid.New().String() + filepath.Ext(file.Filename)
	dst, err := os.Create(filepath.Join(FileDirectory, filename))
	if err != nil {
		return nil, err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	if checkIfImage(filepath.Ext(file.Filename)) == true {
		return &model.Media{
			Url:       FileRequestPrefix + filename,
			MediaType: IMAGE,
		} , nil
	} else {
		return &model.Media{
			Url:       FileRequestPrefix + filename,
			MediaType: VIDEO,
		}, nil

	}
}

func checkIfImage(extension string) bool {
	for _, ext := range ImageExtensions {
		if ext == extension {
			return true
		}
	}
	return false
}

func (m mediaService) GetMedia(ctx context.Context, mediaId string) (*os.File, error) {
	panic("implement me")
}