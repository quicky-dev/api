package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/quicky-dev/generator/generator"
)

func GetAvailableOSes(c echo.Context) error {
	// generates the generic script and returns the uid
	// returns setup script as string
	return c.JSON(http.StatusOK, generator.GetSupportedOses())
}
