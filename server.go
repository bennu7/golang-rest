package main

import (
	"github.com/bennu7/golang-rest/config"
	"github.com/bennu7/golang-rest/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetUpDBConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.GET("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run(":8001")
}
