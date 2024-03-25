package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/AjxGnx/software-engineer-exercise/internal/domain/dto"
	"github.com/AjxGnx/software-engineer-exercise/internal/domain/entity"
	"github.com/AjxGnx/software-engineer-exercise/internal/infra/api/handler"
	mocks "github.com/AjxGnx/software-engineer-exercise/mocks/app"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

var (
	expectedError     = errors.New("some error")
	categorizeMethod  = "Categorize"
	getByNumberMethod = "GetByNumber"
	getMethod         = "Get"
)

type numberCategorizationHandlerTestSuite struct {
	suite.Suite
	NumberCategorizationApp *mocks.NumberCategorization
	underTest               handler.NumberCategorization
}

func TestNumberCategorizationSuite(t *testing.T) {
	suite.Run(t, new(numberCategorizationHandlerTestSuite))
}

func (suite *numberCategorizationHandlerTestSuite) SetupTest() {
	suite.NumberCategorizationApp = &mocks.NumberCategorization{}
	suite.underTest = handler.NewNumberCategorizationHandler(suite.NumberCategorizationApp)
}

func (suite *numberCategorizationHandlerTestSuite) TestCategorize_WhenBindFail() {
	var httpError *echo.HTTPError

	body, _ := json.Marshal("")

	setupCase := SetupControllerCase(http.MethodPost, "/api/exercise/numbers/", bytes.NewBuffer(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.ErrorAs(suite.underTest.Categorize(setupCase.context), &httpError)
	suite.Equal(http.StatusBadRequest, httpError.Code)
}

func (suite *numberCategorizationHandlerTestSuite) TestCategorize_WhenValidateFail() {
	var httpError *echo.HTTPError

	number := `{
		"value": "some string"
	}`

	setupCase := SetupControllerCase(http.MethodPost, "/api/exercise/numbers/", strings.NewReader(number))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.ErrorAs(suite.underTest.Categorize(setupCase.context), &httpError)
	suite.Equal(http.StatusBadRequest, httpError.Code)
}

func (suite *numberCategorizationHandlerTestSuite) TestCategorize_WhenAppFail() {
	var httpError *echo.HTTPError

	Number := dto.Number{Value: 15}
	body := `{
		"value": 15
	}`

	suite.NumberCategorizationApp.Mock.On(categorizeMethod, Number.Value).
		Return(entity.Categorization{}, expectedError)

	setupCase := SetupControllerCase(http.MethodPost, "/api/exercise/numbers/", strings.NewReader(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.ErrorAs(suite.underTest.Categorize(setupCase.context), &httpError)
	suite.Equal(http.StatusInternalServerError, httpError.Code)
}

func (suite *numberCategorizationHandlerTestSuite) TestCategorize_WhenAppSuccess() {
	Number := dto.Number{Value: 15}
	body := `{
		"value": 15
	}`

	suite.NumberCategorizationApp.Mock.On(categorizeMethod, Number.Value).
		Return(entity.Categorization{}, nil)

	setupCase := SetupControllerCase(http.MethodPost, "/api/exercise/numbers/", strings.NewReader(body))
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.NoError(suite.underTest.Categorize(setupCase.context))
	suite.Equal(http.StatusOK, setupCase.Res.Code)
}

func (suite *numberCategorizationHandlerTestSuite) TestGetByNumber_WhenAppFail() {
	var httpError *echo.HTTPError

	numberToCategorize := int64(5)
	numberParam := "number"

	suite.NumberCategorizationApp.Mock.On(getByNumberMethod, numberToCategorize).
		Return(entity.Categorization{}, expectedError)

	setupCase := SetupControllerCase(http.MethodGet, "/api/exercise/numbers/", nil)
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	setupCase.context.SetParamNames(numberParam)
	setupCase.context.SetParamValues(strconv.Itoa(int(numberToCategorize)))

	suite.ErrorAs(suite.underTest.GetByNumber(setupCase.context), &httpError)
	suite.Equal(http.StatusInternalServerError, httpError.Code)
}

func (suite *numberCategorizationHandlerTestSuite) TestGetByNumber_WhenNumberNotFound() {
	var httpError *echo.HTTPError

	numberToCategorize := int64(5)
	numberParam := "number"

	NumberNotFoundError := errors.New("sql: no rows in result set")

	suite.NumberCategorizationApp.Mock.On(getByNumberMethod, numberToCategorize).
		Return(entity.Categorization{}, NumberNotFoundError)

	setupCase := SetupControllerCase(http.MethodGet, "/api/exercise/numbers/", nil)
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	setupCase.context.SetParamNames(numberParam)
	setupCase.context.SetParamValues(strconv.Itoa(int(numberToCategorize)))

	suite.ErrorAs(suite.underTest.GetByNumber(setupCase.context), &httpError)
	suite.Equal(http.StatusNotFound, httpError.Code)
}

func (suite *numberCategorizationHandlerTestSuite) TestGetByNumber_WhenAppSuccess() {
	numberToCategorize := int64(10)
	numberParam := "number"

	suite.NumberCategorizationApp.Mock.On(getByNumberMethod, numberToCategorize).
		Return(entity.Categorization{}, nil)

	setupCase := SetupControllerCase(http.MethodPost, "/api/exercise/numbers/10", nil)
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	setupCase.context.SetParamNames(numberParam)
	setupCase.context.SetParamValues(strconv.Itoa(int(numberToCategorize)))

	suite.NoError(suite.underTest.GetByNumber(setupCase.context))
	suite.Equal(http.StatusOK, setupCase.Res.Code)
}

func (suite *numberCategorizationHandlerTestSuite) TestGet_WhenAppFail() {
	var httpError *echo.HTTPError

	paginateValues := dto.Paginate{
		Page:  1,
		Limit: 10,
	}

	suite.NumberCategorizationApp.Mock.On(getMethod, paginateValues).
		Return(&entity.Paginator{}, expectedError)

	setupCase := SetupControllerCase(http.MethodGet, "/api/exercise/numbers/?page=1&limit=10", nil)
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.ErrorAs(suite.underTest.Get(setupCase.context), &httpError)
	suite.Equal(http.StatusInternalServerError, httpError.Code)
}

func (suite *numberCategorizationHandlerTestSuite) TestGet_WhenAppSuccess() {
	paginateValues := dto.Paginate{
		Page:  1,
		Limit: 10,
	}

	suite.NumberCategorizationApp.Mock.On(getMethod, paginateValues).
		Return(&entity.Paginator{}, nil)

	setupCase := SetupControllerCase(http.MethodPost, "/api/exercise/numbers/?page=1&limit=10", nil)
	setupCase.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.NoError(suite.underTest.Get(setupCase.context))
	suite.Equal(http.StatusOK, setupCase.Res.Code)
}

type ControllerCase struct {
	Req     *http.Request
	Res     *httptest.ResponseRecorder
	context echo.Context
}

func SetupControllerCase(method string, url string, body io.Reader) ControllerCase {
	engine := echo.New()
	req := httptest.NewRequest(method, url, body)
	res := httptest.NewRecorder()
	ctxEngine := engine.NewContext(req, res)

	return ControllerCase{req, res, ctxEngine}
}
