package controllers

import (
	"backend/initializers"
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/resendlabs/resend-go"
	"golang.org/x/crypto/bcrypt"
)

func Forgot(c *gin.Context) {
	var body struct {
		Email string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email",
		})
		return
	}
	id := uuid.New()
	user.Code = id.String()
	initializers.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Check your email",
	})

	params := &resend.SendEmailRequest{
		From:    "Michele <hello@michelemanna.me>",
		To:      []string{user.Email},
		Html:    "<strong>Your reset code is:</strong> " + user.Code,
		Subject: "Reset code",
	}

	sent, err := initializers.Resend.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(sent.Id)
}

func Reset(c *gin.Context) {
	var body struct {
		Code    string
		NewPass string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "code = ?", body.Code)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid code",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.NewPass), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password",
		})
		return
	}

	user.Password = string(hash)
	user.Code = ""
	initializers.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Password changed",
	})
}
