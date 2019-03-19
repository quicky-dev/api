package main

import (
    "github.com/labstack/echo"
    "github.com/quicky-dev/api/controllers"
)

func main() {
    
    // sets up new instance of echo 
    e := echo.New()
    
    // Hello World route to test that server is working 
    e.GET("/", controllers.HelloWorld)

    // api listens on PORT: 3000
    e.Logger.Fatal(e.Start(":3000"))
}
