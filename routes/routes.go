package routes

import (
	"email-service/controller" // Import the correct controller package
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupEmailRoutes(router *gin.Engine) {
	// CORS middleware configuration
	config := cors.Config{
		AllowOrigins:  []string{"*"}, // Allow all origins
		AllowMethods:  []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders:  []string{"*"}, // Allow all headers
		ExposeHeaders: []string{"Content-Length"},
	}

	// Use CORS middleware
	router.Use(cors.New(config))

	// Email routes group
	emailRoutes := router.Group("/email")
	emailRoutes.POST("/contact", controller.SendContactEmail)   
	emailRoutes.POST("/get-a-quote", controller.SendQuoteEmail) 
}
