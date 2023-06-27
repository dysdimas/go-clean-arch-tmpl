package controllers

import (
	"github.com/dysdimas/go-clean-arch-tmpl/db"
	"net/http"

	"github.com/dysdimas/go-clean-arch-tmpl/models"
	"github.com/dysdimas/go-clean-arch-tmpl/services"
	"github.com/dysdimas/go-clean-arch-tmpl/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController() *AuthController {
	db, err := db.GetDBConnection()
	if err != nil {
		panic(err)
	}

	userModel := models.NewUserModel(db)
	authService := services.NewAuthService(userModel)

	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) RegisterUser(ctx *gin.Context) {
	var registerRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := c.authService.RegisterUser(registerRequest.Username, registerRequest.Password, registerRequest.Role)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, "User registered successfully")
}

func (c *AuthController) AuthenticateUser(ctx *gin.Context) {
	var loginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	token, err := c.authService.AuthenticateUser(loginRequest.Username, loginRequest.Password)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{
		"token":   token,
		"message": "Token succesfully created",
	})
}
