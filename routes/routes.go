package routes

import (
	"email-service/controller"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupEmailRoutes(router *gin.Engine, db *gorm.DB) {
	// CORS middleware configuration
	config := cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length"},
	}

	router.Use(cors.New(config))

	emailRoutes := router.Group("/email")
	emailRoutes.POST("/contact", func(c *gin.Context) {
		controller.SendContactEmail(c, db)
	})
	emailRoutes.POST("/get-a-quote", func(c *gin.Context) {
		controller.SendQuoteEmail(c, db)
	})
}
