package main

import (
	"fmt"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
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

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Db, _ = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	// Auto Migrate
	Db.AutoMigrate(&model.User{}, &model.AccountActivation{})

	e := echo.New()

	i := interactor.NewInteractor(Db)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)

	e.Logger.Fatal(e.StartTLS(":443", "certificate.pem", "certificate-key.pem"))

}
