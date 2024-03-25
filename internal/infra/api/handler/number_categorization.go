package handler

import (
	"fmt"
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

// @Tags        Numbers
// @Summary     Categorize a number
// @Description Categorize a number
// @Accept      json
// @Produce     json
// @Param       request body     dto.Number true "Request Body"
// @Success     200       {object} entity.Categorization
// @Failure     400  {object} dto.MessageError
// @Failure     500  {object} dto.MessageError
// @Router      /numbers/ [post]
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

// @Tags        Numbers
// @Summary     Get Categorization by number
// @Description Get Categorization by number
// @Accept      json
// @Produce     json
// @Param       number  path  int true "value of number to find"
// @Success     200       {object} entity.Categorization
// @Failure     404  {object} dto.MessageError
// @Router      /numbers/{number} [get]
func (handler numberCategorizationHandler) GetByNumber(context echo.Context) error {
	number, _ := strconv.Atoi(context.Param("number"))

	categorization, err := handler.numberCategorizationApp.GetByNumber(int64(number))

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("the number: %v does not exist", number))
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())

	}

	return context.JSON(http.StatusOK, dto.Message{
		Message: "categorization number successfully loaded",
		Data:    categorization,
	})
}

// @Tags        Numbers
// @Summary     Get Categorization numbers
// @Description Get Categorization numbers
// @Accept      json
// @Produce     json
// @Param       page      query    string false "page to find"
// @Param       limit     query    string false "limit of page"
// @Success     200       {object} entity.Paginator
// @Failure     500  {object} dto.MessageError
// @Router      /numbers/ [get]
func (handler numberCategorizationHandler) Get(context echo.Context) error {
	page, _ := strconv.Atoi(context.QueryParam("page"))
	limit, _ := strconv.Atoi(context.QueryParam("limit"))
	paginate := dto.Paginate{
		Page:  page,
		Limit: limit,
	}

	categorizations, err := handler.numberCategorizationApp.Get(paginate)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, dto.Message{
		Message: fmt.Sprintf("categorizations successfully loaded"),
		Data:    categorizations,
	})

}
