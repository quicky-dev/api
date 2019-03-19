//Package controllers  will hold all of the logic for our routes that will be 
//called in our main function in main.go
package controllers

import (
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
        log.Fatalln("Caught the following error while generating setup script: ", err)
        os.Exit(1) // exits program due to error 
    }
    return c.Attachment(filePath, "Setup Script") // sends the script file
}

