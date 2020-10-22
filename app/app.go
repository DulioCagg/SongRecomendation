package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	// Creates a router with the default configuration.
	// Includes logger and recovery middleware
	router = gin.Default()
)

// Init loades all the html templates and runs the gin server
func Init() {
	// Loads all the templates at once
	router.LoadHTMLGlob("templates/*")

	// Initializes the routes
	initRoutes()

	// Runs the server
	log.Fatal(router.Run(":8080"))
}
