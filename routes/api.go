package routes

import (
	"github.com/dysdimas/go-clean-arch-tmpl/middleware"
	"github.com/gin-gonic/gin"
)

// SetupAPIRoutes sets up the API routes
func SetupAPIRoutes(router *gin.Engine) {
	// API version 1 group
	v1 := router.Group("/api/v1")

	// Set up the auth routes under the API version 1 group
	authRoutes := v1.Group("/auth")
	{
		SetupAuthRoutes(authRoutes)
	}
	v1.Use(middleware.AuthMiddleware())
	publicRoutes := v1.Group("/user")
	{
		SetupPublicRoutes(publicRoutes)
	}

}
