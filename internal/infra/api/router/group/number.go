package group

import (
	"github.com/AjxGnx/software-engineer-exercise/internal/infra/api/handler"
	"github.com/labstack/echo/v4"
)

const NumberPath = "/numbers/"

type NumberCategorization interface {
	Resource(c *echo.Group)
}

type numberCategorizationGroup struct {
	numberCategorizationHandler handler.NumberCategorization
}

func NewNumberCategorizationGroup(numberCategorizationHandler handler.NumberCategorization) NumberCategorization {
	return numberCategorizationGroup{
		numberCategorizationHandler,
	}
}

func (routes numberCategorizationGroup) Resource(c *echo.Group) {
	groupPath := c.Group(NumberPath)
	groupPath.POST("", routes.numberCategorizationHandler.Categorize)
	groupPath.GET("", routes.numberCategorizationHandler.Get)
	groupPath.GET(":number", routes.numberCategorizationHandler.GetByNumber)
}
