package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
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

	//currently logs the uid of file to terminal. Calls generator.GenerateGeneric()
	e.GET("/api/generic", controllers.GetGeneric).Name = "Generic-Script"

	//GET: returns object of supported downloads
	e.GET("/api/availableItems", controllers.GetItems)

	// POST Route: send arr's of software to setup, returns a custom setup script
	e.POST("/api/dynamic", controllers.GetCustom).Name = "Custom-Script"

	// returns file when user runs setup script
	// from terminal
	e.GET("/api/scripts/:uuid", controllers.GetFile).Name = "Get-File"

	// api listens on PORT: 3000
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
