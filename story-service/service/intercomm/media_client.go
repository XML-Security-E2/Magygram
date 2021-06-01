package intercomm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"story-service/conf"
	"story-service/domain/model"
)

type MediaClient interface {
	SaveMedia(media *multipart.FileHeader) (model.Media, error)

}

type mediaClient struct {}

func (m mediaClient) SaveMedia(media *multipart.FileHeader) (model.Media, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	var fw io.Writer
	defer writer.Close()

	fw, _ = writeFileToRequestBody(media, fw, writer)
	writer.Close()

	retMedia, statusCode, err := handleSaveMediaRequest(body, writer)
	if err != nil || statusCode != http.StatusCreated {
		return model.Media{}, err
	}
	return retMedia, nil
}

func handleSaveMediaRequest(body *bytes.Buffer, writer *multipart.Writer) (model.Media, int, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", baseMediaUrl, bytes.NewReader(body.Bytes()))
	if err != nil {
		return model.Media{}, 0, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
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