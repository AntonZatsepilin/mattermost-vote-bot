package main

import (
	"context"
	"os"
	"time"

	"github.com/AntonZatsepilin/mattermost-vote-bot.git/internal/repository"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tarantool/go-tarantool/v2"
)

func main() {
logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	tarantoolCfg := repository.TarantoolConfig{
		Host:     os.Getenv("TARANTOOL_HOST"),
		Port:     os.Getenv("TARANTOOL_PORT"),
		User:     os.Getenv("TARANTOOL_USER_NAME"),
		Password: os.Getenv("TARANTOOL_USER_PASSWORD"),
		Timeout: viper.GetInt("tarantool.timeout"),
	}

	logrus.Infof("Tarantool config: %v", tarantoolCfg)

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var db *tarantool.Connection
	var err error
	db, err = repository.NewTarantoolDB(ctx, tarantoolCfg)
	if err != nil {
		logrus.Fatalf("failed to initialize Tarantool: %s", err.Error())
	}

	defer db.CloseGraceful()
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}