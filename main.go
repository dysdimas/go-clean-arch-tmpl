package main

import (
	"github.com/dysdimas/go-clean-arch-tmpl/config"
	"github.com/dysdimas/go-clean-arch-tmpl/db"
	"github.com/dysdimas/go-clean-arch-tmpl/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load the configuration
	_, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Set up the database connection
	db, err := db.GetDBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Set up the Gin router
	router := gin.Default()

	// Set up the API routes
	routes.SetupAPIRoutes(router)

	// Run the server
	router.Run(":8000")
}
