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

	e.GET(
		"/api/v1/os/macos/generic",
		controllers.GetMacOSGeneric).Name = "Generic-Script"

	//GET: returns object of supported downloads
	e.GET("/api/v1/os/macos/availableItems", controllers.GetMacOSItems)

	// POST Route: send arr's of software to setup, returns a custom setup script
	e.POST(
		"/api/v1/os/macos/dynamic",
		controllers.GetMacOSCustom).Name = "Custom-Script"

	// returns file when user runs setup script from terminal
	e.GET(
		"/api/v1/os/macos/scripts/:uuid",
		controllers.GetMacOSFile).Name = "Get-MacOS-File"

	/* ----------------------------- Linux Packages ----------------------------- */

	e.GET(
		"/api/v1/os/ubuntu/generic",
		controllers.GetUbuntuGeneric).Name = "Generic-Ubuntu-Script"

	e.GET("/api/v1/os/ubuntu/availableItems", controllers.GetUbuntuItems)

	e.POST(
		"/api/v1/os/ubuntu/dynamic",
		controllers.GetUbuntuCustom).Name = "Custom-Ubuntu-Script"

	e.GET(
		"/api/v1/os/ubuntu/scripts/:uuid",
		controllers.GetUbuntuFile).Name = "Get-Ubuntu-File"

	// api listens on PORT: 3000
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
