package controllers

import (
	"backend/initializers"
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resendlabs/resend-go"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password",
		})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})

	params := &resend.SendEmailRequest{
		From:    "Michele <hello@michelemanna.me>",
		To:      []string{user.Email},
		Html:    "<strong>Welcome</strong> " + user.Email + " to auth backend, thanks for signing up!",
		Subject: "Your first login",
	}

	sent, err := initializers.Resend.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(sent.Id)
}
