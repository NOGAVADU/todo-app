package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nogavadu/todo-app"
	"github.com/nogavadu/todo-app/pkg/handler"
	"github.com/nogavadu/todo-app/pkg/repository"
	"github.com/nogavadu/todo-app/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// @title TODOLIST RESTFUL APP
// @version 1.0
// @description Just API Server for TodoList application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	initLogger()

	if err := initConfig(); err != nil {
		logrus.Fatalf("failed to initialize config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("failed to load enviroment variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewService(services)

	srv := new(todo.Server)
	go func() {
		if err = srv.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("failed to start server: %s", err.Error())
		}
	}()

	logrus.Printf("server started on port=%s", viper.GetString("port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("server shutting down")

	if err = srv.Stop(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err = db.Close(); err != nil {
		logrus.Errorf("error occured on db connetction close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func initLogger() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
}
