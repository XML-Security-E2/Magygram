package service

import (
	"context"
	"github.com/beevik/guid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type MediaService interface {
	SaveMedia(ctx context.Context, files []*multipart.FileHeader) ([]string, error)
	GetMedia(ctx context.Context, mediaId string) (*os.File, error)
}

type mediaService struct {
}

var (
	FileDirectory = "files"
)

func NewMediaService() MediaService {
	return &mediaService{}
}

func (m mediaService) SaveMedia(ctx context.Context, files []*multipart.FileHeader) ([]string, error) {
	var mediaIds []string
	for _, file := range files {
		filename, err := saveFile(file)
		if err != nil {
			return []string{""}, err
		}
		mediaIds = append(mediaIds, filename)

	}
	return mediaIds, nil
}

func saveFile(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	filename := guid.New().String() + ".jpg"
	dst, err := os.Create(filepath.Join(FileDirectory, filename))
	if err != nil {
		return "", err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return filename, nil
}

func (m mediaService) GetMedia(ctx context.Context, mediaId string) (*os.File, error) {
	panic("implement me")
}