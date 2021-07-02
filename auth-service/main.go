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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
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

	usersCol.Indexes().CreateMany(context.Background(),
		[]mongo.IndexModel{{
			Keys: bson.M{
				"email": 1,
			},
			Options: options.Index().SetUnique(true),
		}})

	_, err = usersCol.InsertOne(ctx, bson.D{
		{Key: "_id", Value: "2bdfa5ed-be19-4573-8fe9-49bdfbd8dc12"},
		{Key: "active", Value: true},
		{Key: "email", Value: "admin1@admin.com"},
		{Key: "password", Value: HashAndSaltPasswordIfStrongAndMatching("Adminadmin1!")},
		{Key: "roles", Value: bson.A{
			bson.D{
				{Key: "name", Value: "admin"},
				{Key: "permissions", Value: bson.A{
					bson.D{{Key: "name", Value: "get_verification_request"}},
					bson.D{{Key: "name", Value: "confirm_verification_request"}},
					bson.D{{Key: "name", Value: "reject_verification_request"}},
					bson.D{{Key: "name", Value: "visit_private_profiles"}},
					bson.D{{Key: "name", Value: "delete_posts"}},
					bson.D{{Key: "name", Value: "delete_story"}},
					bson.D{{Key: "name", Value: "search"}},
					bson.D{{Key: "name", Value: "get_logged_info"}},
					bson.D{{Key: "name", Value: "get_user_stories"}},
					bson.D{{Key: "name", Value: "get_storyline_stories"}},
					bson.D{{Key: "name", Value: "visit_user_story"}},
					bson.D{{Key: "name", Value: "search_all_post_by_hashtag"}},
					bson.D{{Key: "name", Value: "search_all_post_by_location"}},
					bson.D{{Key: "name", Value: "view_profile_highlights"}},
					bson.D{{Key: "name", Value: "create_agent"}},
				}}}}},
		{Key: "totp_token", Value: "123"},
	})

	_, err = usersCol.InsertOne(ctx, bson.D{
		{Key: "_id", Value: "37f66e74-71cd-45b5-8070-daeb01323a3b"},
		{Key: "active", Value: true},
		{Key: "email", Value: "admin2@admin.com"},
		{Key: "password", Value: HashAndSaltPasswordIfStrongAndMatching("Adminadmin2!")},
		{Key: "roles", Value: bson.A{
			bson.D{
				{Key: "name", Value: "admin"},
				{Key: "permissions", Value: bson.A{
					bson.D{{Key: "name", Value: "get_verification_request"}},
					bson.D{{Key: "name", Value: "confirm_verification_request"}},
					bson.D{{Key: "name", Value: "reject_verification_request"}},
					bson.D{{Key: "name", Value: "visit_private_profiles"}},
					bson.D{{Key: "name", Value: "delete_posts"}},
					bson.D{{Key: "name", Value: "delete_story"}},
					bson.D{{Key: "name", Value: "search"}},
					bson.D{{Key: "name", Value: "get_user_stories"}},
					bson.D{{Key: "name", Value: "visit_user_story"}},
					bson.D{{Key: "name", Value: "get_logged_info"}},
					bson.D{{Key: "name", Value: "get_storyline_stories"}},
					bson.D{{Key: "name", Value: "search_all_post_by_hashtag"}},
					bson.D{{Key: "name", Value: "search_all_post_by_location"}},
					bson.D{{Key: "name", Value: "view_profile_highlights"}},
					bson.D{{Key: "name", Value: "create_agent"}},
				}}}}},
		{Key: "totp_token", Value: "123"},
	})


	e := echo.New()
	i := interactor.NewInteractor(usersCol, loginEventsCol)
	h := i.NewAppHandler()

	userService := i.NewUserService()
	go userService.RedisConnection()

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

func HashAndSaltPasswordIfStrongAndMatching(password string) string {
	pwd := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	return string(hash)
}

