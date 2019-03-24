package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/quicky-dev/api/controllers"
	"github.com/quicky-dev/generator/generator"
)

func main() {

	// call the generator packages Init function to setup the script factory file
	isFileThere := generator.Init("./scripts", true)
	_ = isFileThere // always should return true (Ask cenz)
	// sets up new instance of echo
	e := echo.New()

	//currently logs the uid of file to terminal. Calls generator.GenerateGeneric()
	e.GET("/api/generic", controllers.GetGeneric).Name = "Generic-Script"

	// this route is strictly for testing our scripts
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<a href="/api/generic">Download Now</a>`)
	})

	//GET: returns object of supported downloads
	e.GET("/api/availableItems", controllers.GetItems)

	// POST Route: send arr's of software to setup, returns a custom setup script
	e.POST("/api/dynamic", controllers.GetCustom).Name = "Custom-Script"

	// returns file when user runs setup script
	// from terminal
	e.GET("/api/:uuid", controllers.GetFile).Name = "Get-File"

	// api listens on PORT: 3000
	e.Logger.Fatal(e.Start(":3000"))
}
