package main

import (
	"flag"
	"github.com/labstack/echo"
	"media-service/conf"
	"media-service/controller/router"
	"media-service/interactor"
)

var runServer = flag.Bool("server", false, "production is -server option require")

func main() {

	conf.NewConfig(*runServer)

	e := echo.New()

	i := interactor.NewInteractor()
	h := i.NewAppHandler()

	router.NewRouter(e, h)

	e.Logger.Fatal(e.StartTLS(":" + conf.Current.Server.Port, "certificate.pem", "certificate-key.pem"))
}
