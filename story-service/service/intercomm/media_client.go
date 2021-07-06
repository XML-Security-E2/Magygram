package intercomm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"story-service/conf"
	"story-service/domain/model"
	"story-service/logger"
	"story-service/tracer"
)

type MediaClient interface {
	SaveMedia(ctx context.Context, media *multipart.FileHeader) (model.Media, error)

}

type mediaClient struct {}

func (m mediaClient) SaveMedia(ctx context.Context, media *multipart.FileHeader) (model.Media, error) {
	span := tracer.StartSpanFromContext(ctx, "MediaClientSaveMedia")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(ctx, span)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	var fw io.Writer
	defer writer.Close()

	fw, _ = writeFileToRequestBody(media, fw, writer)
	writer.Close()

	retMedia, statusCode, err := handleSaveMediaRequest(ctx, body, writer)

	if err != nil || statusCode != http.StatusCreated {
		return model.Media{}, err
	}
	return retMedia, nil
}

func handleSaveMediaRequest(ctx context.Context, body *bytes.Buffer, writer *multipart.Writer) (model.Media, int, error) {
	span := tracer.StartSpanFromContext(ctx, "MediaClientHandleSaveMediaRequest")
	defer span.Finish()

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx,"POST", baseMediaUrl, bytes.NewReader(body.Bytes()))
	if err != nil {
		return model.Media{}, 0, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))
	tracer.Inject(span, req)

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 201 {
		if resp == nil {
			logger.LoggingEntry.Error("Media-service not available")
			return model.Media{}, 0, err
		}

		logger.LoggingEntry.Error("Media-service save media")
		return model.Media{}, 0, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Media{}, 0, err
	}

	var retMedia []model.Media
	json.Unmarshal(bodyBytes, &retMedia)

	return retMedia[0], resp.StatusCode, nil
}

func writeFileToRequestBody(media *multipart.FileHeader, fw io.Writer, writer *multipart.Writer) (io.Writer, error) {
	src, err := media.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()
	fw, err = writer.CreateFormFile("images", media.Filename)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fw, src)
	if err != nil {
		return nil, err
	}
	return fw, nil
}

func NewMediaClient() MediaClient {
	baseMediaUrl = fmt.Sprintf("%s%s:%s/api/media", conf.Current.Mediaservice.Protocol, conf.Current.Mediaservice.Domain, conf.Current.Mediaservice.Port)
	return &mediaClient{}
}

var (
	baseMediaUrl = ""
)