package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignOut(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "User signed out successfully",
	})
}