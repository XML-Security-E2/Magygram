package main

import (
	"auth-service/conf"
	"auth-service/controller/middleware"
	"auth-service/controller/router"
	"auth-service/interactor"
	"auth-service/logger"
	"context"
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var runServer = flag.Bool("auth-service", os.Getenv("IS_PRODUCTION") == "true", "production is -server option require")

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
			Username: conf.Current.Database.User,
			Password: conf.Current.Database.Password,
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

	usersCol := client.Database(*mongoDatabase).Collection("users")
	loginEventsCol := client.Database(*mongoDatabase).Collection("login-events")

	e := echo.New()
	i := interactor.NewInteractor(usersCol, loginEventsCol)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)

	logger.Logger.WithFields(logrus.Fields{
		"host": conf.Current.Server.Host,
		"port":   conf.Current.Server.Port,
	}).Info("Server started")

	if os.Getenv("IS_PRODUCTION") == "true" {
		e.Start(":"+ conf.Current.Server.Port)
	} else {
		e.Logger.Fatal(e.StartTLS(":" + conf.Current.Server.Port, "certificate.pem", "certificate-key.pem"))
	}
}
