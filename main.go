package main

import (
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
    
    // Hello World route to test that server is working 
    e.GET("/", controllers.HelloWorld)

    // api listens on PORT: 3000
    e.Logger.Fatal(e.Start(":3000"))
}
