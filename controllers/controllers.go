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

//GetGeneric creates the Generic setup script and sends it as response
//returns the file of the setup script
func GetGeneric(c echo.Context) error {
    // generates the generic script and returns the uid
    filePath, err := generator.GenerateGeneric() 
    if err != nil {
        log.Fatalln("The following error happened generating script: ", err)
        os.Exit(1) // exits program due to error 
    }
    fmt.Println("Path to da file: ", filePath)
    return c.File(filePath)
}
