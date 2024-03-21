package router

import (
	"github.com/AjxGnx/software-engineer-exercise/internal/infra/api/handler"
	"github.com/AjxGnx/software-engineer-exercise/internal/infra/api/router/group"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	server      *echo.Echo
	numberGroup group.NumberCategorization
}

func New(
	server *echo.Echo,
	numberGroup group.NumberCategorization,
) *Router {
	return &Router{
		server,
		numberGroup,
	}
}

func (r *Router) Init() {
	r.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))

	r.server.Use(middleware.Recover())

	basePath := r.server.Group("/api/exercise")
	basePath.GET("/health", handler.HealthCheck)

	r.numberGroup.Resource(basePath)
}
