package routes

import (
	"github.com/dysdimas/go-clean-arch-tmpl/controllers"
	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes sets up the routes for authentication
func SetupPublicRoutes(router *gin.RouterGroup) {
	router.GET("", controllers.NewUserController().FetchAllUser)
	router.POST("/username", controllers.NewUserController().GetUserByUsername)
	router.DELETE("/delete", controllers.NewUserController().DeleteUserByUsername)
}
