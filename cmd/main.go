package main

import (
	"github.com/nogavadu/todo-app"
	"github.com/nogavadu/todo-app/pkg/handler"
	"github.com/nogavadu/todo-app/pkg/repository"
	"github.com/nogavadu/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewService(services)
	srv := new(todo.Sever)

	if err := srv.Start("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("failed to start server: %s", err.Error())
	}
}
