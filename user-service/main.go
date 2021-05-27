package main

import (
	"context"
	"flag"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
	"user-service/conf"
	"user-service/domain/model"
	"user-service/domain/repository"
)

type application struct {
	users    *repository.UserModel
}

var runServer = flag.Bool("server", false, "production is -server option require")

func main() {

	conf.NewConfig(*runServer)
	mongoDbInfo := fmt.Sprintf("%s:%s",
		conf.Current.Database.Host, conf.Current.Database.Port)


	mongoURI := flag.String("mongoURI", mongoDbInfo, "Database hostname url")
	mongoDatabse := flag.String("mongoDatabse", conf.Current.Database.Database, "Database name")
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

	app := &application{
		users: &repository.UserModel{
			C: client.Database(*mongoDatabse).Collection("users"),
		},
	}

	var user = model.User{Id: "test", Name: "Dusan"}

	app.users.Insert(user)

  	fmt.Println(user.Id)
}
