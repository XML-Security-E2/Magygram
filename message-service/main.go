package main

import (
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo"
	"message-service/conf"
	"message-service/controller/hub"
	"message-service/controller/router"
	"message-service/interactor"
	"os"
)

var Db *redis.Client
var NotifyHub *hub.NotifyHub
var MessageHub *hub.MessageHub
var MessageNotifyHub *hub.MessageNotificationsHub

var runServer = flag.Bool("message-service", os.Getenv("IS_PRODUCTION") == "true", "production is -server option require")

func main()  {

	conf.NewConfig(*runServer)

	Db = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Current.Database.Host, conf.Current.Database.Port),
		Password: conf.Current.Database.Password, // no password set
		DB:       conf.Current.Database.Database,  // use default DB
	})

	//err := Db.Set(context.TODO(), "key/123", "dusan", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = Db.Set(context.TODO(), "key/124", "dusancina", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//val, err := Db.Get(context.TODO(), "key/124").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key", val)
	//
	//keys, _, _ := Db.Scan(context.TODO(), 0, "key/*", 1).Result()
	//
	//fmt.Println(keys)
	//for _, key := range keys {
	//	fmt.Println(key)
	//}
	//iter := Db.Scan(context.TODO(), "key/*", "", 100).Iterator()
	//for iter.Next() {
	//	err := client.Del(iter.Val()).Err()
	//	if err != nil { panic(err) }
	//}
	//fmt.Println("key", val)
	NotifyHub = hub.NewNotifyHub()
	go NotifyHub.Run()

	MessageHub = hub.NewHub()
	go MessageHub.Run()

	MessageNotifyHub = hub.NewMessageNotificationsHub()
	go MessageNotifyHub.Run()

	e := echo.New()
	i := interactor.NewInteractor(Db, NotifyHub, MessageHub, MessageNotifyHub)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	//middleware.NewMiddleware(e)

	if os.Getenv("IS_PRODUCTION") == "true" {
		e.Start(":"+ conf.Current.Server.Port)
	} else {
		e.Logger.Fatal(e.StartTLS(":" + conf.Current.Server.Port, "certificate.pem", "certificate-key.pem"))
	}
}
