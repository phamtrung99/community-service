package http

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/community-service/delivery/http/v1/comment"
	"github.com/phamtrung99/community-service/repository"
	"github.com/phamtrung99/community-service/usecase"
	"github.com/phamtrung99/gopkg/middleware"
	"github.com/phamtrung99/movie-service/config"
)

// NewHTTPHandler .
func NewHTTPHandler(repo *repository.Repository, ucase *usecase.UseCase) *echo.Echo {
	e := echo.New()
	cfg := config.GetConfig()

	skipper := func(c echo.Context) bool {
		p := c.Request().URL.Path

		return strings.Contains(p, "/health_check")
	}

	e.Use(middleware.Auth(cfg.Jwt.Key, skipper, false))

	apiV1 := e.Group("/v1")

	comment.Init(apiV1.Group("/comments"), ucase)

	return e
}
