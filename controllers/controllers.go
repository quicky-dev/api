//Package controllers  will hold all of the logic for our routes that will be 
//called in our main function in main.go
package controllers

import (
    "net/http"
    "fmt"
    "log"
    "os"

    "github.com/labstack/echo"
    "github.com/quicky-dev/generator/generator"
)

//HelloWorld is a fake function for testing to make sure our server 
//is setup and working properly
func HelloWorld(c echo.Context) error {
    return c.String(http.StatusOK, "Hello World!")
}


//GetGeneric creates the Generic setup script and sends it as response
func GetGeneric(c echo.Context) error {
    // generates the generic script and returns the uid
    uid, err := generator.GenerateGeneric() 
    if err != nil {
        log.Fatalln("The following error happened generating script: ", err)
        os.Exit(1)
    }
    fmt.Println(uid) 
    return c.String(http.StatusOK, "Something did indeed happen")
}
