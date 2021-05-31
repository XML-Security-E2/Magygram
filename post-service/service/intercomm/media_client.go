package intercomm

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"post-service/conf"
)

type MediaClient interface {
	SaveMedia(mediaList []*multipart.FileHeader) error
	
}

type mediaClient struct {}

func (m mediaClient) SaveMedia(mediaList []*multipart.FileHeader) error {
	body := &bytes.Buffer{}
	client := &http.Client{}
	writer := multipart.NewWriter(body)
	var fw io.Writer
	defer writer.Close()
	fmt.Println(len(mediaList))

	for idx, media := range mediaList {
		fw, _ = writeFileToRequestBody(media, fw, writer, idx)
	}
	writer.Close()

	baseUrl := fmt.Sprintf("%s%s:%s/api/media", conf.Current.Mediaservice.Protocol, conf.Current.Mediaservice.Domain, conf.Current.Mediaservice.Port)
	fmt.Println(baseUrl)

	req, err := http.NewRequest("POST", baseUrl, bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return err
	}

	return nil
}

func writeFileToRequestBody(media *multipart.FileHeader, fw io.Writer, writer *multipart.Writer, idx int) (io.Writer, error) {
	src, err := media.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()
	fw, err = writer.CreateFormFile("images[" + string(idx)+"]", "test.png")
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
	baseUrl = fmt.Sprintf("%s%s:%s/api/media", conf.Current.Mediaservice.Protocol, conf.Current.Mediaservice.Domain, conf.Current.Mediaservice.Port)
	return &mediaClient{}
}

var (
	baseUrl = ""
)