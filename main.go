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

    //WIP: currently logs the uid of file to terminal. Calls generator.GenerateGeneric()
    e.GET("/generic", controllers.GetGeneric)

    // this route is strictly for testing our scripts
    e.GET("/", func(c echo.Context) error {
        return c.HTML(http.StatusOK, `<a href="/generic">Download Now</a>`)
    })
    // api listens on PORT: 3000
    e.Logger.Fatal(e.Start(":3000"))
}
