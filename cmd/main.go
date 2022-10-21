package main

import (
	"log"

	"github.com/asadbek/todo-app"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running httpserver :%s", err.Error())
	}

}
