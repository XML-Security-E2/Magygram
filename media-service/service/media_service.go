package service

import (
	"context"
	"github.com/beevik/guid"
	"github.com/sirupsen/logrus"
	"io"
	"media-service/domain/model"
	"media-service/logger"
	"mime/multipart"
	"os"
	"path/filepath"
)

type MediaService interface {
	SaveMedia(ctx context.Context, files []*multipart.FileHeader) ([]model.Media, error)
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
	var mediaIdLog []string
	for _, file := range files {
		media , err := saveFile(file)
		if err != nil {
			logger.LoggingEntry.Error("Error while saving file")
			return nil, err
		}
		mediaIds = append(mediaIds, *media)
		mediaIdLog = append(mediaIdLog, media.Url)
	}

	logger.LoggingEntry.WithFields(logrus.Fields{"file_paths" : mediaIds}).Info("Files saved")

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