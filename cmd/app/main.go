package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"studyCRUD/internal/database"
	"studyCRUD/internal/handlers"
	"studyCRUD/internal/taskService"
	"studyCRUD/internal/userService"
	"studyCRUD/internal/web/tasks"
	"studyCRUD/internal/web/users"
)

func main() {

	database.InitDB()

	repoTask := taskService.NewTaskRepository(database.DB)
	serviceTasks := taskService.NewTaskService(repoTask)

	repoUsers := userService.NewUserRepository(database.DB, repoTask)
	serviceUsers := userService.NewUserService(repoUsers)

	taskHTTPHandler := handlers.NewTaskHandler(serviceTasks)
	userHTTPHandler := handlers.NewUserHandler(serviceUsers)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	tasksHandler := tasks.NewStrictHandler(taskHTTPHandler, nil)
	tasks.RegisterHandlers(e, tasksHandler)

	usersHandler := users.NewStrictHandler(userHTTPHandler, nil)
	users.RegisterHandlers(e, usersHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatal("failed to start with error", err)
	}
}
