package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/quicky-dev/api/controllers"
	"github.com/quicky-dev/generator/generator"
)

//Handler is main function for Quicky-api
func Handler(w http.ResponseWriter, r *http.Request) {
	main()
}

func main() {

	// call the generator packages Init function to setup the script factory file
	isFileThere := generator.Init("./scripts", false)
	_ = isFileThere // always should return true (Ask cenz)
	// sets up new instance of echo
	e := echo.New()

	// Headers allowed over CORS
	var allowedHeaders = []string{
		echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept}

	// Register CORS policy for https://quicky.dev, our frontend domain.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://quicky.dev"},
		AllowHeaders: allowedHeaders,
	}))

	/* ----------------------------- MacOS Packages ------------------------- */

	//currently logs the uid of file to terminal. Calls generator.GenerateGeneric()
	e.GET(
		"/api/v1/macos/generic", controllers.GetGeneric).Name = "Generic-Script"

	//GET: returns object of supported downloads
	e.GET("/api/v1/macos/availableItems", controllers.GetItems)

	// POST Route: send arr's of software to setup, returns a custom setup script
	e.POST(
		"/api/v1/macos/dynamic", controllers.GetCustom).Name = "Custom-Script"

	// returns file when user runs setup script from terminal
	e.GET("/api/v1/macos/scripts/:uuid", controllers.GetFile).Name = "Get-File"

	/* ----------------------------- Linux Packages ----------------------------- */

	e.GET(
		"/api/v1/linux/generic", controllers.GetGeneric).Name = "Generic-Script"

	e.GET("/api/v1/linux/availableItems", controllers.GetItems)

	e.POST(
		"/api/v1/linux/dynamic", controllers.GetCustom).Name = "Custom-Script"

	e.GET("/api/v1/linux/scripts/:uuid", controllers.GetFile).Name = "Get-File"

	// api listens on PORT: 3000
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
