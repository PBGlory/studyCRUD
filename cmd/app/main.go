package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"studyCRUD/internal/database"
	"studyCRUD/internal/handlers"
	"studyCRUD/internal/taskService"
	"studyCRUD/internal/web/tasks"
)

func main() {

	database.InitDB()

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewTaskService(repo)

	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatal("failed to start with error", err)
	}
}
