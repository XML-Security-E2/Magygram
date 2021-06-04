package main

import (
	"flag"
	"os"
	"github.com/labstack/echo"
	"media-service/conf"
	"media-service/controller/router"
	"media-service/interactor"
)

var runServer = flag.Bool("media-service", os.Getenv("IS_PRODUCTION") == "true", "production is -server option require")

func main() {

	conf.NewConfig(*runServer)

	e := echo.New()

	i := interactor.NewInteractor()
	h := i.NewAppHandler()

	router.NewRouter(e, h)

	if os.Getenv("IS_PRODUCTION") == "true" {
		e.Start(":"+ conf.Current.Server.Port)
	} else {
		e.Logger.Fatal(e.StartTLS(":" + conf.Current.Server.Port, "certificate.pem", "certificate-key.pem"))
	}}
