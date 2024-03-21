package router_test

import (
	"net/http"
	"testing"

	"github.com/AjxGnx/software-engineer-exercise/internal/infra/api/router"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

var (
	paths = []string{
		"/api/exercise",
		"/api/exercise/",
		"/api/exercise/*",
		"/api/exercise/health",
	}

	methods = []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodPatch,
		http.MethodOptions,
		http.MethodConnect,
		http.MethodHead,
		http.MethodTrace,
		"PROPFIND",
	}
)

func TestRouterSuite(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}

type RouterTestSuite struct {
	suite.Suite
	server    *echo.Echo
	underTest *router.Router
}

func (suite *RouterTestSuite) SetupTest() {
	suite.server = echo.New()

	suite.underTest = router.New(
		suite.server,
	)
}

func (suite *RouterTestSuite) TestInit() {
	suite.underTest.Init()

	for _, value := range suite.server.Routes() {
		suite.Contains(paths, value.Path)
		suite.Contains(methods, value.Method)
	}
}
