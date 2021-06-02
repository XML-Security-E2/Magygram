package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"post-service/conf"
	"post-service/controller/router"
	"post-service/interactor"
	"time"
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

	postCol := client.Database(*mongoDatabase).Collection("posts")
	locationCol := client.Database(*mongoDatabase).Collection("locations")
	tagCol := client.Database(*mongoDatabase).Collection("tags")

	e := echo.New()
	i := interactor.NewInteractor(postCol, locationCol, tagCol)
	h := i.NewAppHandler()

	router.NewRouter(e, h)

	e.Logger.Fatal(e.StartTLS(":" + conf.Current.Server.Port, "certificate.pem", "certificate-key.pem"))
}
