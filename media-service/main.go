package main

import (
	"flag"
	"media-service/conf"
	"media-service/controller/middleware"
	"media-service/controller/router"
	"media-service/interactor"
	"media-service/logger"
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

var runServer = flag.Bool("media-service", os.Getenv("IS_PRODUCTION") == "true", "production is -server option require")

func main() {

	logger.InitLogger()

	conf.NewConfig(*runServer)

	e := echo.New()

	i := interactor.NewInteractor()
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
