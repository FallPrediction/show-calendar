package middleware

import (
	"net/http"
	"souflair/utils"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("authorization")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token not found",
			})
			c.Abort()
			return
		}

		claims, err := (&utils.Jwt{}).ParseUserToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token invalid",
			})
			c.Abort()
			return
		}
		c.Set("userData", claims.UserData)
		c.Next()
	}
}
