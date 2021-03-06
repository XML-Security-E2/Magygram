package main

import (
	"ads-service/conf"
	"ads-service/controller/middleware"
	"ads-service/controller/router"
	"ads-service/interactor"
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var runServer = flag.Bool("ads-service", os.Getenv("IS_PRODUCTION") == "true", "production is -server option require")

func main() {

	conf.NewConfig(*runServer)
	mongoDbInfo := fmt.Sprintf("%s:%s", conf.Current.Database.Host, conf.Current.Database.Port)

	mongoURI := flag.String("mongoURI", mongoDbInfo, "Database hostname url")
	mongoDatabase := flag.String("mongoDatabse", conf.Current.Database.Database, "Database name")
	enableCredentials := flag.Bool("enableCredentials", false, "Enable the use of credentials for mongo connection")
	flag.Parse()

	co := options.Client().ApplyURI(*mongoURI)
	if *enableCredentials {
		co.Auth = &options.Credential{
			Username: os.Getenv(conf.Current.Database.User),
			Password: os.Getenv(conf.Current.Database.Password),
		}
	}

	client, err := mongo.NewClient(co)
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	campaignCol := client.Database(*mongoDatabase).Collection("campaigns")
	influencerCampaignCol := client.Database(*mongoDatabase).Collection("influencer-campaigns")
	updateReqCampaignCol := client.Database(*mongoDatabase).Collection("update-campaign-requests")

	e := echo.New()
	i := interactor.NewInteractor(campaignCol, influencerCampaignCol, updateReqCampaignCol)
	h := i.NewAppHandler()
	i.NewSchedulerService()

	router.NewRouter(e, h)

	metricsMiddleware := middleware.NewMetricsMiddleware()
	e.Use(metricsMiddleware.Metrics)

	if os.Getenv("IS_PRODUCTION") == "true" {
		e.Start(":" + conf.Current.Server.Port)
	} else {
		e.Logger.Fatal(e.StartTLS(":"+conf.Current.Server.Port, "certificate.pem", "certificate-key.pem"))
	}
}
