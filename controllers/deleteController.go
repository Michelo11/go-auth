package controllers

import (
	"backend/initializers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	user, _ := c.Get("user")
	initializers.DB.Unscoped().Delete(&models.User{}, user)
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
