package handler

import (
	"net/http"
	"strconv"

	"github.com/AjxGnx/software-engineer-exercise/internal/app"
	"github.com/AjxGnx/software-engineer-exercise/internal/domain/dto"
	"github.com/labstack/echo/v4"
)

type NumberCategorization interface {
	Categorize(context echo.Context) error
	GetByNumber(context echo.Context) error
	Get(context echo.Context) error
}

type numberCategorizationHandler struct {
	numberCategorizationApp app.NumberCategorization
}

func NewNumberCategorizationHandler(numberCategorizationApp app.NumberCategorization) NumberCategorization {
	return numberCategorizationHandler{
		numberCategorizationApp,
	}
}

func (handler numberCategorizationHandler) Categorize(context echo.Context) error {
	var number dto.Number

	if err := context.Bind(&number); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := number.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	categorization, err := handler.numberCategorizationApp.Categorize(number.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, dto.Message{
		Message: "number categorized successfully",
		Data:    categorization,
	})
}

func (handler numberCategorizationHandler) GetByNumber(context echo.Context) error {
	number, _ := strconv.Atoi(context.Param("number"))

	categorization, err := handler.numberCategorizationApp.GetByNumber(int64(number))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return context.JSON(http.StatusOK, dto.Message{
		Message: "categorization number successfully loaded",
		Data:    categorization,
	})
}

func (handler numberCategorizationHandler) Get(context echo.Context) error {
	//TODO implement me
	panic("implement me")
}
