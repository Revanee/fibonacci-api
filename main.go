package main

import (
	"github.com/Revanee/fibonacci-api/pkg/routes"
	"github.com/gin-gonic/contrib/static"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./web/fibonacci/build", true)))

	// Setup route group for the API
	api := router.Group("/")
	{
		api.GET("/check_fibonacci/:number", routes.CheckFibonacci)
		api.POST("/check_fibonacci/:number", routes.CheckFibonacci)
		api.GET("/check_number/:number", routes.CheckPrime)
		api.POST("/check_number/:number", routes.CheckPrime)
	}

	// Start and run the server
	router.Run(":8080")
}
