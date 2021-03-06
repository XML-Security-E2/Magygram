package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string
		Host string
		Secret string
		Handshake string
		Apikey string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
	Mail struct {
		Sender   string
		Password string
		Host 	 string
		Port 	 string
	}

	Magygram struct {
		Protocol   string
		Domain string
		Port 	 string
	}

	Xmldb struct {
		Protocol   string
		Domain string
		Username string
		Password string
		Port 	 string
	}
}

var Current *Config

func NewConfig(runServer bool) {
	var C Config
	Current = &C
	viper.AddConfigPath("./conf")
	viper.SetConfigType("yml")

	if runServer {
		viper.SetConfigName("production")
	} else {
		viper.SetConfigName("local")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal config file error: %s", err))
	}

	if err := viper.Unmarshal(&C); err != nil {
		panic(fmt.Errorf("fatal config file error: %s", err))
	}
	return
}
