// Package main starts the Echo HTTP server.
package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/yoshiyoshifujii/go-echo-sample/internal/api"
	"github.com/yoshiyoshifujii/go-echo-sample/internal/handlers"
	"github.com/yoshiyoshifujii/go-echo-sample/internal/repository/memory"
	"github.com/yoshiyoshifujii/go-echo-sample/internal/usecase"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userRepo := memory.NewUserRepository()
	createUser := usecase.NewCreateUser(userRepo)
	findUser := usecase.NewFindUser(userRepo)

	api.RegisterHandlers(e, api.NewStrictHandler(handlers.NewHandler(createUser, findUser), nil))

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
