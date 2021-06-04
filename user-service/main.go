package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
	"user-service/conf"
	"user-service/controller/middleware"
	"user-service/controller/router"
	"user-service/interactor"
)


var runServer = flag.Bool("user-service", os.Getenv("IS_PRODUCTION") == "true", "production is -server option require")

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
	accActivationsCol := client.Database(*mongoDatabase).Collection("account-activations")
	resetPasswordsCol := client.Database(*mongoDatabase).Collection("reset-passwords")
	usersCol.Indexes().CreateMany(context.Background(),
		[]mongo.IndexModel{{
			Keys: bson.M{
				"email": 1,
			},
			Options: options.Index().SetUnique(true),
		}, {
			Keys: bson.M{
				"username": 1,
			},
			Options: options.Index().SetUnique(true),
		}})


	e := echo.New()
	i := interactor.NewInteractor(usersCol, accActivationsCol, resetPasswordsCol)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)

	if os.Getenv("IS_PRODUCTION") == "true" {
		e.Start(":"+ conf.Current.Server.Port)
	} else {
		e.Logger.Fatal(e.StartTLS(":" + conf.Current.Server.Port, "certificate.pem", "certificate-key.pem"))
	}
}
