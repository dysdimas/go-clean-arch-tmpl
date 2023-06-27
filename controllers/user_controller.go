package controllers

import (
	"github.com/dysdimas/go-clean-arch-tmpl/db"
	"github.com/dysdimas/go-clean-arch-tmpl/models"
	"github.com/dysdimas/go-clean-arch-tmpl/services"
	"github.com/dysdimas/go-clean-arch-tmpl/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	db, err := db.GetDBConnection()
	if err != nil {
		panic(err)
	}

	userModel := models.NewUserModel(db)
	userService := services.NewUserService(userModel)

	return &UserController{
		userService: userService,
	}
}

func (c *UserController) GetUserByUsername(ctx *gin.Context) {
	var usernameRequest struct {
		Username string `json:"username" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&usernameRequest); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}
	userData, err := c.userService.GetUserByUsername(usernameRequest.Username)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Username not found")
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{
		"user_info": userData,
		"message":   "Success",
	})
}

func (c *UserController) FetchAllUser(ctx *gin.Context) {

	listUser, err := c.userService.FetchAllUser()
	if err != nil {
		utils.RespondWithError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{
		"users":   listUser,
		"message": "Success",
	})
}

func (c *UserController) DeleteUserByUsername(ctx *gin.Context) {
	var usernameRequest struct {
		Username string `json:"username" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&usernameRequest); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}
	_, err := c.userService.DeleteUserByUsername(usernameRequest.Username)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Delete fail")
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{
		"message": "Delete successfully",
	})
}
