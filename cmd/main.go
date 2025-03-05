package main

import (
	"github.com/nogavadu/todo-app"
	"github.com/nogavadu/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Sever)

	if err := srv.Start("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("failed to start server: %s", err.Error())
	}
}
