//Package controllers  will hold all of the logic for our routes that will be 
//called in our main function in main.go
package controllers

import (
    "net/http"

    "github.com/labstack/echo"
)

//HelloWorld is a fake function for testing to make sure our server 
//is setup and working properly
func HelloWorld(c echo.Context) error {
    return c.String(http.StatusOK, "Hello World!")
}
