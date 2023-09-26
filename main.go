package main

import (
	"backend/controllers"
	"backend/initializers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
	initializers.ResendMail()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.SignIn)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/signout", middleware.RequireAuth, controllers.SignOut)
	r.DELETE("/delete", middleware.RequireAuth, controllers.Delete)

	r.Run()
}
