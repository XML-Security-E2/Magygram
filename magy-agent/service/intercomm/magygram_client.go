package intercomm

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"magyAgent/conf"
	"magyAgent/domain/model"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type MagygramClient interface {
	CreateCampaign(request *model.CampaignApiRequest) error
}

type magygramClient struct {

}

var (
	baseMagygramUrl = ""
)

func NewMagygramClient() MagygramClient {
	baseMagygramUrl = fmt.Sprintf("%s%s:%s/api/ads/campaign/agent", conf.Current.Magygram.Protocol, conf.Current.Magygram.Domain, conf.Current.Magygram.Port)
	return &magygramClient{}
}

func (m magygramClient) CreateCampaign(request *model.CampaignApiRequest) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	var fw io.Writer

	src, err :=  os.Open(request.FilePath)
	if err != nil {
		return err
	}
	defer src.Close()
	fw, err = writer.CreateFormFile("images", filepath.Base(request.FilePath))
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, src)
	if err != nil{
		return err
	}

	writer.WriteField("minDisplaysForRepeatedly", strconv.Itoa(request.MinDisplaysForRepeatedly))
	writer.WriteField("frequency", string(request.Frequency))
	writer.WriteField("minAge", strconv.Itoa(request.TargetGroup.MinAge))
	writer.WriteField("maxAge", strconv.Itoa(request.TargetGroup.MaxAge))
	writer.WriteField("gender", string(request.TargetGroup.Gender))

	writer.WriteField("displayTime", strconv.Itoa(request.DisplayTime))
	writer.WriteField("dateFrom", strconv.FormatInt(request.DateFrom, 10))
	writer.WriteField("dateTo", strconv.FormatInt(request.DateTo, 10))
	writer.WriteField("exposeOnceDate", strconv.FormatInt(request.ExposeOnceDate, 10))
	writer.WriteField("campaignType", string(request.Type))

	writer.Close()
	client := &http.Client{}
	req, err := http.NewRequest("POST", baseMagygramUrl, bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", "Bearer " + conf.Current.Server.Apikey)

	hash, _ := bcrypt.GenerateFromPassword([]byte(conf.Current.Server.Secret), bcrypt.MinCost)
	req.Header.Add(conf.Current.Server.Handshake, string(hash))

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 201 {
		if resp == nil {
			return err
		}

		return errors.New("error while creating post campaign")
	}

	return nil
}