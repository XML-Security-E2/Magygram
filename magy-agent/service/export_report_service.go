package service

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"magyAgent/conf"
	"magyAgent/domain/model"
	"magyAgent/service/intercomm"
	"strconv"
	"os"
	"strings"
)

type ExportService interface {
	ExportToPdf(report *model.CampaignStatisticReport) (bytes.Buffer, error)
}


type exportService struct {
	intercomm.MediaClient
}

func NewExportService (mc intercomm.MediaClient) ExportService {
	return &exportService{mc}
}


func (e exportService) ExportToPdf(report *model.CampaignStatisticReport) (bytes.Buffer, error) {

	m := pdf.NewMaroto(consts.Portrait, consts.A4)


		m.RegisterHeader(func() {
			m.Row(20, func() {
				m.Col(4, func() {
					_ = m.FileImage("./magygramlogo.png", props.Rect{
						Center:  true,
						Percent: 80,
					})
				})
			})
		})


	m.Row(20, func() {
		m.Col(8, func() {
			m.Text("Campaign statistics report for date: " + report.DateCreating.Format("01-02-2006"), props.Text{
				Top:   10,
				Style: consts.Bold,
				Size:  12,
				Align: consts.Left,
			})
		})
	})


	for _, item := range report.Campaigns {
		str := ""
		if os.Getenv("IS_PRODUCTION") == "true" {
			strs := strings.Split(item.Media.Url, "/")
			tmp := fmt.Sprintf("%s%s:%s/api/media/", conf.Current.Magygram.Protocol, conf.Current.Magygram.Domain, conf.Current.Magygram.Port) + strs[len(strs) - 1]
			mediaa, _ := e.MediaClient.GetMedia(tmp)
			str = b64.StdEncoding.EncodeToString(mediaa)
		} else {
			media, _ := e.MediaClient.GetMedia(item.Media.Url)
			str = b64.StdEncoding.EncodeToString(media)
		}

		format := item.Media.Url[len(item.Media.Url) - 3 :len(item.Media.Url)]

		m.Row(30, func() {
			m.Col(3, func() {
				m.Base64Image(str, consts.Extension(format), props.Rect{
					Center:  true,
					Top:     0,
					Percent: 80,
				})
			})
			if item.Frequency == "REPEATEDLY" {
				if item.CampaignStatus == "REGULAR" {
					m.Col(4, func() {
						m.Text("Campaign info", props.Text{
							Top: 3,
							Style: consts.Bold,
							Size:  10,
							Align: consts.Left,
						})
						m.Text("Campaign type: " + string(item.Type), props.Text{
							Top: 7,
							Size:  8,
							Align: consts.Left,
						})
						m.Text("Website: " + item.Website, props.Text{
							Top: 10,
							Size:  8,
							Align: consts.Left,
						})
						m.Text("Date from: " + item.DateFrom.Format("01-02-2006"), props.Text{
							Top: 13,
							Size:  8,
							Align: consts.Left,
						})
						m.Text("Date to: " + item.DateTo.Format("01-02-2006"), props.Text{
							Top: 16,
							Size:  8,
							Align: consts.Left,
						})
						m.Text("Minimum times to display: " + strconv.Itoa(item.MinDisplaysForRepeatedly), props.Text{
							Top: 19,
							Size:  8,
							Align: consts.Left,
						})
					})
				} else {
					m.Col(4, func() {
						m.Text("Campaign info", props.Text{
							Top: 3,
							Style: consts.Bold,
							Size:  10,
							Align: consts.Left,
						})
						m.Text("Campaign type: " + string(item.Type), props.Text{
							Top: 7,
							Size:  8,
							Align: consts.Left,
						})
						m.Text("Website: " + item.Website, props.Text{
							Top: 10,
							Size:  8,
							Align: consts.Left,
						})
						m.Text("Influencer: @" + item.InfluencerUsername, props.Text{
							Top: 13,
							Size:  8,
							Align: consts.Left,
						})
					})
				}

			} else {
				if item.CampaignStatus == "REGULAR" {
					m.Col(4, func() {
						m.Text("Campaign info", props.Text{
							Top: 3,
							Style: consts.Bold,
							Size:  10,
							Align: consts.Left,
						})
						m.Text("Campaign type: " + string(item.Type), props.Text{
							Top: 7,
							Size:  8,
							Align: consts.Left,
						})
						m.Text("Website: " + item.Website, props.Text{
							Top: 10,
							Size:  8,
							Align: consts.Left,
						})
						m.Text("Exposure date: " + item.ExposeOnceDate.Format("01-02-2006"), props.Text{
							Top: 13,
							Size:  8,
							Align: consts.Left,
						})
						m.Text("Exposure time: " + strconv.Itoa(item.DisplayTime) + " h", props.Text{
							Top: 16,
							Size:  8,
							Align: consts.Left,
						})
					})
				} else {
					m.Col(4, func() {
						m.Text("Campaign info", props.Text{
							Top: 3,
							Style: consts.Bold,
							Size:  10,
							Align: consts.Left,
						})
						m.Text("Campaign type: " + string(item.Type), props.Text{
							Top: 7,
							Size:  8,
							Align: consts.Left,
						})
						m.Text("Website: " + item.Website, props.Text{
							Top: 10,
							Size:  8,
							Align: consts.Left,
						})
						m.Text("Influencer: @" + item.InfluencerUsername, props.Text{
							Top: 13,
							Size:  8,
							Align: consts.Left,
						})
					})
				}

			}

			if item.Type == "POST" {
				m.Col(3, func() {
					m.Text("Campaign reach", props.Text{
						Top: 3,
						Style: consts.Bold,
						Size:  10,
						Align: consts.Left,
					})
					m.Text("Website clicks: " + strconv.Itoa(item.WebsiteClicks), props.Text{
						Top: 7,
						Size:  8,
						Align: consts.Left,
					})
					m.Text("Daily average views: " + fmt.Sprintf("%.2f", item.DailyAverage), props.Text{
						Top: 10,
						Size:  8,
						Align: consts.Left,
					})
					m.Text("Campaign views: " + strconv.Itoa(item.UserViews), props.Text{
						Top: 13,
						Size:  8,
						Align: consts.Left,
					})
					m.Text("Likes: " +  strconv.Itoa(item.Likes), props.Text{
						Top: 16,
						Size:  8,
						Align: consts.Left,
					})
					m.Text("Dislikes: " +  strconv.Itoa(item.Dislikes), props.Text{
						Top: 19,
						Size:  8,
						Align: consts.Left,
					})
					m.Text("Comments: " +  strconv.Itoa(item.Comments), props.Text{
						Top: 22,
						Size:  8,
						Align: consts.Left,
					})
				})
			} else {
				m.Col(3, func() {
					m.Text("Campaign reach", props.Text{
						Top: 3,
						Style: consts.Bold,
						Size:  10,
						Align: consts.Left,
					})
					m.Text("Website clicks: " + strconv.Itoa(item.WebsiteClicks), props.Text{
						Top: 7,
						Size:  8,
						Align: consts.Left,
					})
					m.Text("Daily average views: " + fmt.Sprintf("%.2f", item.DailyAverage), props.Text{
						Top: 10,
						Size:  8,
						Align: consts.Left,
					})
					m.Text("Campaign views: " + strconv.Itoa(item.UserViews), props.Text{
						Top: 13,
						Size:  8,
						Align: consts.Left,
					})
					m.Text("Story views: " +  strconv.Itoa(item.StoryViews), props.Text{
						Top: 16,
						Size:  8,
						Align: consts.Left,
					})
				})
			}

			m.Col(2, func() {
				m.Text("Target group", props.Text{
					Top: 3,
					Style: consts.Bold,
					Size:  10,
					Align: consts.Left,
				})
				m.Text("Min age: " + strconv.Itoa(item.TargetGroup.MinAge), props.Text{
					Top: 7,
					Size:  8,
					Align: consts.Left,
				})
				m.Text("Max age: " + strconv.Itoa(item.TargetGroup.MaxAge), props.Text{
					Top: 10,
					Size:  8,
					Align: consts.Left,
				})
				m.Text("Gender: " + string(item.TargetGroup.Gender), props.Text{
					Top: 13,
					Size:  8,
					Align: consts.Left,
				})

			})
		})
	}

	return m.Output()
}

