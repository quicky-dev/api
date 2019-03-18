package main

import (
    "net/http"

    "github.com/labstack/echo"
)

func main() {
    
    // sets up new instance of echo 
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello world!")
    })
    // api listens on PORT: 3000
    e.Logger.Fatal(e.Start(":3000"))
}
