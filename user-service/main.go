package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
	"user-service/conf"
	"user-service/controller/middleware"
	"user-service/controller/router"
	"user-service/interactor"
)


var runServer = flag.Bool("server", false, "production is -server option require")

func main() {

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

	usersCol := client.Database(*mongoDatabase).Collection("users")
	accActivationsCol := client.Database(*mongoDatabase).Collection("account-activations")
	loginEventsCol := client.Database(*mongoDatabase).Collection("login-events")
	resetPasswordsCol := client.Database(*mongoDatabase).Collection("reset-passwords")

	e := echo.New()
	i := interactor.NewInteractor(usersCol, accActivationsCol, loginEventsCol, resetPasswordsCol)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)

	e.Logger.Fatal(e.StartTLS(":" + conf.Current.Server.Port, "certificate.pem", "certificate-key.pem"))
}
