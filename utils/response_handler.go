package utils

import (
	"github.com/gin-gonic/gin"
)

// RespondWithError sends an error response with the given message
func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"error": message,
	})
}

// RespondWithSuccess sends a success response with the given data
func RespondWithSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"response": data,
	})
}
