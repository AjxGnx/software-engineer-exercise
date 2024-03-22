package main

import (
	"fmt"

	"github.com/AjxGnx/software-engineer-exercise/internal/infra/api/router"

	"github.com/AjxGnx/software-engineer-exercise/cmd/providers"
	"github.com/AjxGnx/software-engineer-exercise/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// @title        Number Categorization Manager
// @version      1.0.0
// @description  Number Categorization Manager
// @license.name Alirio Gutierrez
// @BasePath     /api/exercise
// @schemes      http
func main() {
	container := providers.BuildContainer()
	err := container.Invoke(func(router *router.Router, server *echo.Echo) {
		router.Init()

		server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%v", config.Environments().ServerHost,
			config.Environments().ServerPort)))
	})

	if err != nil {
		log.Panic(err)
	}
}
