package main

/*
import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	// Defining the API endpoints
	api.GET("/cities", CityHandler)
	api.POST("/cities/car/:carID", WhichCar)
	api.POST("/cities/sem/:semID", WhichSem)

	// Start and run the server
	router.Run(":3000")
}

// CityHandler : algo
func CityHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Cities Handler not implemented yet",
	})
}

// WhichCar : algo
func WhichCar(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "WhichCar not implemented yet",
	})
}

// WhichSem : algo
func WhichSem(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "WhichSem not implemented yet",
	})
}
*/
