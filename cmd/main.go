package main

import (
	"blockchain_api/internal/handlers"
	"blockchain_api/internal/repository"
	"blockchain_api/internal/service"
	"blockchain_api/pkg/server"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Blockchain HTTP API
// @version 1.0
// @description API Server for Blockchain Application

// @host localhost:8000
// @BasePath /

func main() {

	// a recover function from panics
	defer func() {
		if errConnect := recover(); errConnect != nil {
			logrus.Printf("error to the connection: %s \n", errConnect)
		}
	}()

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if errConf := initConfig(); errConf != nil {
		logrus.Fatalf("error initializating configs: %s", errConf.Error())
	}

	if errLoadEnv := godotenv.Load(); errLoadEnv != nil {
		logrus.Fatalf("error loading env variables: %s", errLoadEnv.Error())
	}

	db := repository.NewConnectionPostgresDB(&repository.ConfigPostgres{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	go service.GetAndSaveBlockchainList(db)

	repos := repository.NewRepositiry(db)
	service := service.NewService(repos)
	handlers := handlers.NewHandler(service)

	srv := new(server.Server)
	if errRunServer := srv.RunServer(handlers.InitRoutes()); errRunServer != nil {
		logrus.Fatalf("error occurred while running http server: %s", errRunServer)
	}
}

func initConfig() error {
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
