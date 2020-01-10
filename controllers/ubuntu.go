// Package controllers will hold all of the logic for our routes
// that will be called in our main function in main.go
package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	filestore "github.com/quicky-dev/api/fileStore"
	"github.com/quicky-dev/generator/generator"
)

var ubuntu generator.UBUNTU_GENERATOR = generator.GetUbuntuGenerator()

/* ------------------------------- GetGeneric ------------------------------- */

// GetGeneric creates the Generic setup script and sends it
// as response returns the file of the setup script
func GetUbuntuGeneric(c echo.Context) error {
	// generates the generic script and returns the uid
	script, err := macos.GenerateGenericScript()
	if err != nil {
		log.Fatalln("Caught the following error while generating setup script: ", err)
		os.Exit(1) // exits program due to error
	}
	// returns setup script as string
	return c.String(http.StatusOK, script.Payload)
}

/* -------------------------------- GetCustom ------------------------------- */

//GetCustom takes in the list of software the user wants to download
//from the the request body, the lists are then used to generate the correct
//setup script
func GetUbuntuCustom(c echo.Context) error {

	// struct for setup script factory
	var install generator.InstallRequest
	// read request body
	reqBody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	// unmarshal json from request body into install request struct
	err = json.Unmarshal(reqBody, &install)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// generate the bash script
	script, err := macos.GenerateDynamicScript(install)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Create S3 Session and get handler
	handler, err := filestore.GetHandler()
	if err != nil {
		log.Fatal("There was an error getting S3 Session: ", err)
	}

	// Uploads file to S3 Bucket
	err = handler.UploadFile(script.UUID, script.Payload)
	if err != nil {
		log.Fatal("There was an error uploading file to S3 : ", err)
	}
	// send path to file as a download
	return c.String(http.StatusOK, script.UUID)
}

/* --------------------------------- GetFile -------------------------------- */

//GetFile takes in uuid and sends user the file to the install via CL
func GetUbuntuFile(c echo.Context) error {
	uuid := c.Param("uuid")
	// get S3 Handler
	handler, err := filestore.GetHandler()
	if err != nil {
		log.Fatal("There was an error getting S3 Session: ", err)
	}
	script, err := handler.ReadFile(uuid)
	if err != nil {
		fmt.Println("KEY: ", uuid)
		fmt.Println("There was an error getting setup script: ", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	// sends the setup script as string
	return c.String(http.StatusOK, script)
}

/* -------------------------------- GetItems -------------------------------- */

//GetItems sends struct of supported items for download
func GetUbuntuItems(c echo.Context) error {
	// sends all supported mac packages as big JSON obj
	return c.JSON(http.StatusOK, ubuntu.AvailablePackages)
}
