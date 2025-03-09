package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nogavadu/todo-app"
	"github.com/nogavadu/todo-app/pkg/handler"
	"github.com/nogavadu/todo-app/pkg/repository"
	"github.com/nogavadu/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("failed to initialize config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load enviroment variables: %s", err.Error())
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
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewService(services)
	srv := new(todo.Sever)

	if err := srv.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("failed to start server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
