package middleware

import (
	"net/http"
	"show-calendar/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if len(token) <= 7 && !strings.HasPrefix(token, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token invalid",
			})
			c.Abort()
			return
		}

		claims, err := (&utils.Jwt{}).ParseUserToken(token[7:])
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
