package routes

import (
	"github.com/dysdimas/go-clean-arch-tmpl/controllers"
	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes sets up the routes for authentication
func SetupAuthRoutes(router *gin.RouterGroup) {
	router.POST("/register", controllers.NewAuthController().RegisterUser)
	router.POST("/login", controllers.NewAuthController().AuthenticateUser)
	// Add more auth routes as needed
}
