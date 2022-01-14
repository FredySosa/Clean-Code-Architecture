package main

import (
	"log"

	"github.com/FredySosa/cleanCode/internal/models"
	"github.com/FredySosa/cleanCode/internal/store"
	"github.com/FredySosa/cleanCode/internal/transport"
	"github.com/FredySosa/cleanCode/internal/usecase"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	countryStore := store.NewCountryStore(map[string]models.Country{
		"AUS": {
			ID:   "AUS",
			Name: "Austria",
		},
		"ARG": {
			ID:   "ARG",
			Name: "Argentina",
		},
	})
	countryUseCase := usecase.NewCountriesUseCase(countryStore)
	echoHandler := transport.NewHTTPHandler(countryUseCase)
	echoHandler.Pre(middleware.RemoveTrailingSlash())
	echoHandler.Use(middleware.Logger())

	log.Fatal(echoHandler.Start(":8080"))
}
