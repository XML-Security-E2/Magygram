package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"magyAgent/conf"
	"magyAgent/controller/middleware"
	"magyAgent/controller/router"
	"magyAgent/domain/model"
	"magyAgent/interactor"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "test"
)

var Db *gorm.DB
var runServer = flag.Bool("server", false, "production is -server option require")

func main() {
	conf.NewConfig(*runServer)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Current.Database.Host, conf.Current.Database.Port, conf.Current.Database.User, conf.Current.Database.Password, conf.Current.Database.Database)

	Db, _ = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	// Auto Migrate
	Db.AutoMigrate(&model.User{}, &model.AccountActivation{})

	e := echo.New()

	i := interactor.NewInteractor(Db)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)

	e.Logger.Fatal(e.StartTLS(":" + conf.Current.Server.Port, "certificate.pem", "certificate-key.pem"))

}
