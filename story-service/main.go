package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"story-service/conf"
	"story-service/controller/middleware"
	"story-service/controller/router"
	"story-service/interactor"
	"story-service/logger"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var runServer = flag.Bool("story-service", os.Getenv("IS_PRODUCTION") == "true", "production is -server option require")

func main() {

	logger.InitLogger()

	conf.NewConfig(*runServer)
	mongoDbInfo := fmt.Sprintf("%s:%s",
		conf.Current.Database.Host, conf.Current.Database.Port)

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

	storyCol := client.Database(*mongoDatabase).Collection("stories")

	e := echo.New()
	i := interactor.NewInteractor(storyCol)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)

	metricsMiddleware := middleware.NewMetricsMiddleware()
	e.Use(metricsMiddleware.Metrics)

	logger.Logger.WithFields(logrus.Fields{
		"host": conf.Current.Server.Host,
		"port": conf.Current.Server.Port,
	}).Info("Server started")

	if os.Getenv("IS_PRODUCTION") == "true" {
		e.Start(":" + conf.Current.Server.Port)
	} else {
		e.Logger.Fatal(e.StartTLS(":"+conf.Current.Server.Port, "certificate.pem", "certificate-key.pem"))
	}
}
