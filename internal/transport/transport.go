package transport

import (
	"context"
	"net/http"

	"github.com/FredySosa/cleanCode/internal/models"
	"github.com/labstack/echo/v4"
)

type (
	CountriesUseCase interface {
		GetCountries(ctx context.Context, limit, offset string) ([]models.Country, string, error)
	}
	HTTPHandler struct {
		useCase CountriesUseCase
	}
)

func NewHTTPHandler(useCase CountriesUseCase) *echo.Echo {
	h := HTTPHandler{
		useCase: useCase,
	}

	e := echo.New()
	h.initRoutes(e)

	return e
}

func (h HTTPHandler) initRoutes(e *echo.Echo) {
	e.GET("/ping", h.Ping)

	countriesGroup := e.Group("/countries")
	countriesGroup.GET("", h.GetCountries)
}

func (h HTTPHandler) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}

func (h HTTPHandler) GetCountries(c echo.Context) error {
	ctx := c.Request().Context()
	limit, offset := "100", "0"

	countries, newOffset, err := h.useCase.GetCountries(ctx, limit, offset)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, "some error happened")
	}
	c.Response().Header().Add("Offset", newOffset)

	return c.JSON(http.StatusOK, countries)
}
