package main

import (
	"os"
	"flag"
	"github.com/labstack/echo"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"relationship-service/conf"
	"relationship-service/controller/middleware"
	"relationship-service/controller/router"
	"relationship-service/interactor"
)

var runServer = flag.Bool("relationship-service", os.Getenv("IS_PRODUCTION") == "true", "production is -server option require")

func main(){
	conf.NewConfig(*runServer)
	driver, err := neo4j.NewDriver(conf.Current.Database.Host, neo4j.BasicAuth(conf.Current.Database.User, conf.Current.Database.Password, ""))
	if err != nil {
		panic(err)
	}
	e := echo.New()
	i := interactor.NewInteractor(driver)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)

	if os.Getenv("IS_PRODUCTION") == "true" {
		e.Start(":"+ conf.Current.Server.Port)
	} else {
		e.Logger.Fatal(e.StartTLS(":" + conf.Current.Server.Port, "certificate.pem", "certificate-key.pem"))
	}}