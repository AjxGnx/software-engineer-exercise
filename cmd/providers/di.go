package providers

import (
	"github.com/AjxGnx/software-engineer-exercise/internal/app"
	"github.com/AjxGnx/software-engineer-exercise/internal/infra/adapters/pg"
	"github.com/AjxGnx/software-engineer-exercise/internal/infra/adapters/pg/repository"
	"github.com/AjxGnx/software-engineer-exercise/internal/infra/api/handler"
	"github.com/AjxGnx/software-engineer-exercise/internal/infra/api/router"
	"github.com/AjxGnx/software-engineer-exercise/internal/infra/api/router/group"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(pg.ConnInstance)
	_ = Container.Provide(router.New)

	_ = Container.Provide(group.NewNumberCategorizationGroup)
	_ = Container.Provide(handler.NewNumberCategorizationHandler)
	_ = Container.Provide(app.NewNumberCategorizationApp)
	_ = Container.Provide(repository.NewNumberCategorizationRepo)

	return Container
}
