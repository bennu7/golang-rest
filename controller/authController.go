package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// interface contract controller
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

// service
type authController struct {
}

// *: NewAuthController creates a new instance of AuthController
func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": " hii login"})

}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": " hii register"})
}
