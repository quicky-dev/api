//Package controllers  will hold all of the logic for our routes that will be 
//called in our main function in main.go
package controllers

import (
    "log"
    "os"
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"

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

//GetCustom takes in the list of software the user wants to download
//from the the request body, the lists are then used to generate the correct 
//setup script
func GetCustom(c echo.Context) error {

    // struct for setup script factory
    var installRequest generator.InstallRequest
    // read request body
    reqBody, err := ioutil.ReadAll(c.Request().Body)
    if err != nil {
        log.Fatalln(err)
    }

    // unmarshal json from request body into install request struct
    err = json.Unmarshal(reqBody, &installRequest)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(installRequest)
    return c.String(http.StatusOK, "All Good")
}
