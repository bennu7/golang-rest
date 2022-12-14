package middleware

import (
	"log"
	"net/http"

	"github.com/bennu7/golang-rest/helper"
	"github.com/bennu7/golang-rest/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// validate token jwt
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		// jika dapat data tokennya proses di ValidateToken
		token, err := jwtService.ValidateToken(authHeader)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("claims ->", claims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer]: ", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
